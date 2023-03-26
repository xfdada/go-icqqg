package im

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// 在线客服聊天功能
type Im struct {
	upGrader websocket.Upgrader
}

type Message struct {
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	To       string `json:"id"`   //接收者的ID
	From     string `json:"from"` //发送者ID
	Type     string `json:"type"`
	Content  string `json:"content"`   //发送内容
	ISManage string `json:"is_manage"` //是否是管理员
}
type SendMessage struct {
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	To       string `json:"id"` //接收者的ID
	Type     string `json:"type"`
	Content  string `json:"content"` //发送内容
}

type Manage struct {
	Conn       *websocket.Conn            //连接实例
	Connection map[string]*websocket.Conn //维护的用户链接
}

// 管理者的ID
var manageConn = make(map[string]Manage)

func NewIm(upGrader websocket.Upgrader) *Im {
	return &Im{
		upGrader: upGrader,
	}
}

func (im *Im) Toke(c *gin.Context) {

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
	var message Message
	err = json.Unmarshal(firstMsg, &message)
	if err != nil {
		fmt.Println("消息序列化错误 :", err)
		return
	}

	if message.ISManage == "manage" {
		manageConn[message.To] = Manage{
			Conn:       conn,
			Connection: make(map[string]*websocket.Conn),
		}
	} else {
		if _, ok := manageConn[message.From]; !ok {
			conn.WriteMessage(websocket.TextMessage, []byte("管理员不存在"))
			err = conn.Close()
			if err != nil {
				fmt.Println("服务端断开失败 :", err)
				return
			}

		} else {
			fmt.Println("71--------客户端初始化或者更新 :", message)
			manageConn[message.From].Connection[message.To] = conn
			fmt.Println("73--------客户端初始化或者更新 :", manageConn)
		}
		//判断是否存在这个ID，如果存在就更新，不存在就创建

	}
	//不存在就直接退出，发送消息断开连接

	fmt.Println("85 ------------", manageConn, "---------", manageConn[message.To])
	//上述是第一次接收消息下面可以从redis中获取或者数据库中获取自动发送的欢迎消息
	defer conn.Close()

	for {

		msgTypes, Msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("WebSocket 读取消息错误 :", err)
			return
		}
		err = json.Unmarshal(Msg, &message)
		if err != nil {
			fmt.Println("消息序列化错误 :", err)
			return
		}
		fmt.Println("101--------", manageConn, "----", manageConn[message.To])
		//管理员的话直接通
		if message.ISManage == "manage" {
			fmt.Println("103------", manageConn[message.From], "------", message.To)
			manageConn[message.From].CSend(msgTypes, message)
		} else {
			fmt.Println("103------", manageConn[message.To], "------", message.From)
			manageConn[message.To].Send(msgTypes, message)
		}
	}

}

// Send 发送消息
func (m Manage) Send(msgType int, msg Message) {
	fmt.Println("22222222222")
	data := SendMessage{
		Username: msg.Username,
		Avatar:   msg.Avatar,
		To:       msg.From,
		Type:     msg.Type,
		Content:  msg.Content,
	}
	sendMsg, _ := json.Marshal(&data)
	manageConn[msg.To].Conn.WriteMessage(msgType, sendMsg)
}

func (m Manage) CSend(msgType int, msg Message) {
	fmt.Println(msg, "----------", msg)
	fmt.Println(m)
	if _, ok := manageConn[msg.From].Connection[msg.To]; ok {
		data := SendMessage{
			Username: msg.Username,
			Avatar:   msg.Avatar,
			To:       msg.To,
			Type:     msg.Type,
			Content:  msg.Content,
		}
		sendMsg, _ := json.Marshal(&data)
		manageConn[msg.From].Connection[msg.To].WriteMessage(msgType, sendMsg)
	} else {
		fmt.Println("11111111111ffffff")
		manageConn[msg.From].Conn.WriteMessage(msgType, []byte("用户不在线"))
	}

}
