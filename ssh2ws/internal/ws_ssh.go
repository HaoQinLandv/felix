package internal

import (
	"github.com/dejavuzhou/felix/flx"
	"github.com/dejavuzhou/felix/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024 * 1024 * 10,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// handle webSocket connection.
// first,we establish a ssh connection to ssh server when a webSocket comes;
// then we deliver ssh data via ssh connection between browser and ssh server.
// That is, read webSocket data from browser (e.g. 'ls' command) and send data to ssh server via ssh connection;
// the other hand, read returned ssh data from ssh server and write back to browser via webSocket API.
func WsSsh(c *gin.Context) {
	cols, err := strconv.Atoi(c.DefaultQuery("cols", "120"))
	if handleError(c, err) {
		return
	}
	rows, err := strconv.Atoi(c.DefaultQuery("rows", "32"))
	if handleError(c, err) {
		return
	}
	id := c.Param("id")
	client, err := flx.NewSshClient(id)
	if handleError(c, err) {
		return
	}
	defer client.Close()
	ssConn, err := utils.NewSshConn(cols, rows, client)
	if handleError(c, err) {
		return
	}
	defer ssConn.Close()
	// after configure, the WebSocket is ok.
	wsConn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if handleError(c, err) {
		return
	}
	defer wsConn.Close()

	quitChan := make(chan bool, 3)

	// most messages are ssh output, not webSocket input
	go ssConn.ReceiveWsMsg(wsConn, quitChan)
	go ssConn.SendComboOutput(wsConn, quitChan)
	ssConn.SessionWait(quitChan)

	<-quitChan
	logrus.Info("websocket finished")
}
