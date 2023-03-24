package im

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

//在线客服聊天功能
type Im struct {
	upGrader websocket.Upgrader
}

type Message struct {
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	ID       string `json:"id"`
	Type     string `json:"type"`
	Content  string `json:"content"`
}

var connections = make(map[string]*websocket.Conn)

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
	defer conn.Close()
	// 此处可以处理连接事件，比如记录连接数，发送欢迎消息等
	connectionID := uuid.New().String()
	connections[connectionID] = conn
	// 此处可以处理连接事件，比如记录连接数，发送欢迎消息等
	data, err := json.Marshal(map[string]interface{}{"msg": "您好，欢迎您使用聊天平台"})
	conn.WriteMessage(websocket.TextMessage, data)
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Failed to read message from WebSocket connection:", err)
			delete(connections, connectionID) // 从映射关系中删除连接标识符
			return
		}
		var msg Message
		json.Unmarshal(p, &msg)
		fmt.Printf("Received message from connection %s: %s\n", connectionID, string(p))
		if connections[msg.ID] != nil {
			toconn := connections[msg.ID]
			ms, _ := json.Marshal(msg)
			toconn.WriteMessage(messageType, ms)
		} else {
			connections[msg.ID] = conn
			ms, _ := json.Marshal(msg)
			conn.WriteMessage(messageType, ms)
		}

		// 此处可以处理消息事件，比如广播消息给其他连接的客户端
	}

}
