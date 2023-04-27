package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"gin-icqqg/utils/snow_id"
	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
	"time"
)

type Config struct {
	Server  *Server
	Redis   *Redis
	Mysql   *Mysql
	Jwt     *Jwt
	Logger  *Logger
	DB      *gorm.DB
	Upload  *Upload
	CaptCha *CaptCha
	AlySms  *AlySms
}

type ImManage struct {
	Conn       *websocket.Conn            //连接实例
	Connection map[string]*websocket.Conn //维护的用户链接
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
	Page      string `json:"url"`
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

var ImConn = make(map[string]ImManage)

var AppConfig Config
var Snow *snow_id.Snow
var (
	port    string
	runMode string
	cfgpath string
)

func init() {
	err := setupFlag()
	if err != nil {
		fmt.Println(err)
	}
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	time.Local = loc
	err = initConfig()
	if err != nil {
		panic(err)
	}
	DB, err := NewDB()
	if err != nil {
		panic(err)
	}
	AppConfig.DB = DB
	Snow = snow_id.NewSnow(1)

}

func initConfig() error {
	err := NewConfig("config.yaml")
	if err != nil {
		return err
	}
	return nil
}
func NewConfig(config string) error {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("./")
	if config != "" {
		vp.AddConfigPath(config)
	}
	vp.SetConfigType("yaml") // 设置配置文件类型格式为YAML
	err := vp.ReadInConfig()
	err = vp.Unmarshal(&AppConfig)
	if err != nil {
		return err
	}
	vp.WatchConfig()
	vp.OnConfigChange(func(e fsnotify.Event) {
		if err = vp.Unmarshal(&AppConfig); err != nil {
			log.Printf("Config file changed filed: %s", err)
		}
	})
	return nil
}

func setupFlag() error {
	flag.StringVar(&port, "port", "", "启动端口")
	flag.StringVar(&runMode, "mode", "", "启动模式")
	flag.StringVar(&cfgpath, "config", "./", "配置文件的路径")
	flag.Parse()
	return nil
}

// Send 用户向管理员发送消息
func (m ImManage) Send(msgType int, ipAddr string, msg Message) (string, string) {
	data := SendMessage{
		Username:  ipAddr,
		Avatar:    msg.Avatar,
		To:        msg.From,
		Type:      msg.Type,
		MsgType:   msg.MsgType,
		Content:   msg.Content,
		GroupName: msg.GroupName,
		GroupId:   msg.GroupId,
	}
	sendMsg, _ := json.Marshal(&data)
	ImConn[msg.To].Conn.WriteMessage(msgType, sendMsg)
	return msg.From, string(sendMsg)
}

//CSend 管理员给用户发送消息
func (m ImManage) CSend(msgType int, ipAddr string, msg Message) (string, string) {
	if _, ok := ImConn[msg.From].Connection[msg.To]; ok {
		data := SendMessage{
			Username:  ipAddr,
			Avatar:    msg.Avatar,
			To:        msg.To,
			Type:      msg.Type,
			Content:   msg.Content,
			MsgType:   msg.MsgType,
			GroupName: msg.GroupName,
			GroupId:   msg.GroupId,
		}
		sendMsg, _ := json.Marshal(&data)
		ImConn[msg.From].Connection[msg.To].WriteMessage(msgType, sendMsg)
		return msg.To, string(sendMsg)
	} else {
		data := map[string]string{"type": "offline", "user_id": msg.To}
		sendMsg, _ := json.Marshal(&data)
		ImConn[msg.From].Conn.WriteMessage(msgType, sendMsg)
		return "", ""
	}
}

//Offline 用户离线
func (m ImManage) Offline(manageId, connectionID string) {
	data := map[string]string{"type": "offline", "user_id": connectionID}
	sendMsg, _ := json.Marshal(&data)
	err := ImConn[manageId].Conn.WriteMessage(websocket.TextMessage, sendMsg)
	if err != nil {
		fmt.Println("WebSocket 发送消息错误 :", err)
	}
}
