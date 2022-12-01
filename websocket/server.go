package websocket

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	// 用户的连接
	uc = make(map[string]*websocket.Conn)

	// 收到客户端的消息
	umsg chan map[string]string
)

// 连接升级
func Ping(c *gin.Context) {
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	c.Request.ParseForm()
	// roomid := c.Request.Form["rid"][0]
	userid := c.Request.Form["uid"][0]
	// 用户id + 用户的连接 存入map
	uc[userid] = ws
	// 最后关闭
	defer close(ws, userid)
	for {
		// 收消息动作
		_, message, err := ws.ReadMessage()
		msg := make(map[string]string)
		json.Unmarshal(message, &msg)
		umsg <- msg
		if err != nil {
			return
		}

	}
}

func close(ws *websocket.Conn, uid string) {
	ws.Close()
	// 从map 中去掉下线的连接
	delete(uc, uid)
}

func SendMsg(msg map[string]string) {
	/* for k, v := range msg {

	} */
}
