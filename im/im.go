package im

import (
	"encoding/json"
	"fmt"
	"gin-icqqg/config"
	"gin-icqqg/model"
	"gin-icqqg/utils/redis"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// 在线客服聊天功能
type Im struct {
	upGrader websocket.Upgrader
}
type Message struct {
	Username  string `json:"username"`
	Avatar    string `json:"avatar"`
	To        string `json:"id"`   //接收者的ID
	From      string `json:"from"` //发送者ID
	Type      string `json:"type"`
	Content   string `json:"content"`   //发送内容
	ISManage  string `json:"is_manage"` //是否是管理员
	GroupName string `json:"groupname"`
	GroupId   string `json:"groupid"`
	MsgType   string `json:"msgtype"`
	Origin    string `json:"rerfer"`
}
type SendMessage struct {
	Username  string `json:"username"`
	Avatar    string `json:"avatar"`
	To        string `json:"id"` //接收者的ID
	Type      string `json:"type"`
	Content   string `json:"content"` //发送内容
	GroupName string `json:"groupname"`
	GroupId   string `json:"groupid"`
	MsgType   string `json:"msgtype"`
}

func NewIm(upGrader websocket.Upgrader) *Im {
	return &Im{
		upGrader: upGrader,
	}
}

var ipAddr string

//Toke 单点登录
//如果页面刷新，或者页面新开，则旧链接失效
func (im *Im) Toke(c *gin.Context) {
	var myStruct config.ImManage
	conn, err := im.upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Failed to upgrade HTTP connection to WebSocket:", err)
		return
	}
	//第一次读取消息
	_, firstMsg, err := conn.ReadMessage()
	if err != nil {
		fmt.Println("WebSocket 第一次读取消息错误 :", err)
		return
	}
	var message config.Message
	err = json.Unmarshal(firstMsg, &message)
	if err != nil {
		fmt.Println("消息序列化错误 :", err)
		return
	}
	ipAddr = GetIpAddr(c.ClientIP())
	//生成UUID
	connectionID := uuid.New().String()
	var manageID string
	//判断是否是管理员登录
	if message.ISManage == "manage" {
		manageID = message.To
		connectionID = message.To
		if _, ok := config.ImConn[message.To]; !ok {
			myStruct = config.ImManage{
				Conn:       conn,
				Connection: make(map[string]*websocket.Conn),
			}
			config.ImConn[message.To] = myStruct
		} else {
			myStruct = config.ImConn[message.To]
			myStruct.Conn = conn
			config.ImConn[message.To] = myStruct
		}
	} else {
		manageID = message.From
		//先判断管理员是否上线
		var flow model.AddFlow
		flow.Page = message.Page
		flow.PageNum = 1
		flow.City = ipAddr
		flow.GroupId = message.GroupId
		flow.IP = c.ClientIP()
		flow.Origin = message.Origin
		fl := model.NewFlow()
		if _, ok := config.ImConn[message.From]; !ok {
			//管理员不在线
			data := make(map[string]string)
			if message.To != "" {
				data = map[string]string{"msgtype": "offline", "id": message.To}
				flow.UserId = message.To
			} else {
				data = map[string]string{"msgtype": "offline", "id": connectionID}
				flow.UserId = connectionID
			}
			fl.UpOrAdd(flow)
			msg, _ := json.Marshal(data)
			conn.WriteMessage(websocket.TextMessage, msg)
			err = conn.Close()
			if err != nil {
				fmt.Println("服务端断开失败 :", err)
			}
			return
		} else {
			//判断用户是否有ID，如果没有则新建绑定
			if message.To != "" {
				config.ImConn[message.From].Connection[message.To] = conn
				connectionID = message.To
			} else {
				//新增用户并存储到数据库中
				var user model.AddImUser
				ip := c.ClientIP()
				if message.Origin == "" {
					user.Origin = "直接访问流量"
				} else {
					user.Origin = message.Origin
				} //获取流量来源
				user.UserId = connectionID
				user.IP = ip
				user.UserName = ipAddr
				user.Manage = message.From
				user.Group = message.GroupName
				user.GroupId = message.GroupId
				ims := model.NewImUser()
				go ims.Add(user)
				config.ImConn[message.From].Connection[connectionID] = conn
				data := map[string]string{"msgtype": "bind", "id": connectionID}
				msg, _ := json.Marshal(data)
				conn.WriteMessage(websocket.TextMessage, msg)
			}
			flow.UserId = connectionID
			go fl.UpOrAdd(flow)
			Online(message.From, connectionID)
		}
	}

	defer conn.Close()
	for {
		msgTypes, Msg, err := conn.ReadMessage()
		if err != nil {
			config.ImConn[manageID].Offline(manageID, connectionID)
			//断开后存储聊天记录到数据库中
			insertMysql(connectionID)
			redis.RedisHDel(manageID, connectionID)
			delete(config.ImConn[manageID].Connection, connectionID)
			return
		}
		err = json.Unmarshal(Msg, &message)
		if err != nil {
			fmt.Println("消息序列化错误 :", err)
			return
		}
		if message.ISManage == "manage" {
			uid, msg := config.ImConn[message.From].CSend(msgTypes, ipAddr, message)
			if uid != "" {
				rPush(uid, msg)
			}

		} else {

			uid, msg := config.ImConn[message.To].Send(msgTypes, ipAddr, message)
			rPush(uid, msg)
		}
	}

}

//rPush 向聊天列表尾部插入数据
func rPush(uid, msg string) {
	redis.RPUSH(uid, msg)
}

//insertMysql 把聊天信息插入到数据库中
func insertMysql(uid string) {
	length, err := redis.LLEN(uid)
	if err != nil {
		config.ErrorLog(fmt.Sprintf("im 获取聊天数据错误:%s", err))
	}
	var add []model.ImMessage
	for i := 0; i < int(length); i++ {
		var one model.ImMessage
		val, _ := redis.LPOP(time.Second, uid)
		json.Unmarshal([]byte(val[1]), &one)
		add = append(add, one)
	}
	msg := model.NewImMessage()
	msg.Add(add)

}

type IpAddr struct {
	Country  string `json:"country,omitempty"`
	Province string `json:"prov,omitempty"`
	City     string `json:"city,omitempty"`
}

type IpData struct {
	Data IpAddr `json:"data,omitempty"`
}

func GetIpAddr(ip string) string {
	client := &http.Client{}
	request, err := http.NewRequest("GET", "https://gwgp-cekvddtwkob.n.bdcloudapi.com/ip/geo/v1/district?ip="+ip, nil)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.64 Safari/537.36 Edg/101.0.1210.53")
	if err != nil {
		config.ErrorLog(fmt.Sprintf("http 请求失败: %s", err.Error()))
		return "未知地区"
	}
	res, errs := client.Do(request)
	if errs != nil {
		config.ErrorLog(fmt.Sprintf("http 请求失败: %s", errs))
		return "未知地区"
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			config.ErrorLog(fmt.Sprintf("http 请求: %d %s", res.StatusCode, res.Status))
		}
	}(res.Body)
	if res.StatusCode != 200 {
		config.ErrorLog(fmt.Sprintf("http 请求: %d %s", res.StatusCode, res.Status))
		return "未知地区"
	}
	body, _ := ioutil.ReadAll(res.Body)
	var addr IpData
	json.Unmarshal(body, &addr)
	return addr.Data.Country + "-" + addr.Data.Province + "-" + addr.Data.City
}

//Online 用户在线
func Online(manageId, connectionID string) {
	data := map[string]string{"type": "online", "user_id": connectionID}
	sendMsg, _ := json.Marshal(&data)
	err := config.ImConn[manageId].Conn.WriteMessage(websocket.TextMessage, sendMsg)
	if err != nil {
		fmt.Println("WebSocket 发送消息错误 :", err)
	}
}
