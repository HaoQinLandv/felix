package controllers

import (
	"github.com/dejavuzhou/felix/flx"
	"github.com/dejavuzhou/felix/ssh2ws/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

//const SSH_EGG = `genshen<genshenchu@gmail.com> https://github.com/genshen/sshWebConsole"`

var upGrader = websocket.Upgrader{
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
	id := c.Param("id")
	client, err := flx.NewSshClient(id)
	if handleError(c, err) {
		return
	}
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if handleError(c, err) {
		return
	}
	defer ws.Close()

	//setup ssh connection
	sshConn := utils.SSHShellSession{
		SshClient: client,
	}
	// set io for ssh session std out
	var wsBuff WebSocketBufferWriter
	sshConn.WriterPipe = &wsBuff

	defer sshConn.Close()
	// config ssh
	cols, err := strconv.Atoi(c.DefaultQuery("cols", "120"))
	if handleError(c, err) {
		return
	}
	rows, err := strconv.Atoi(c.DefaultQuery("rows", "32"))
	if handleError(c, err) {
		return
	}
	err = sshConn.Config(cols, rows)
	if handleError(c, err) {
		return
	}

	// an egg:
	//if err := sshEntity.Session.Setenv("SSH_EGG", SSH_EGG); err != nil {
	//	log.Println(err)
	//}
	// after configure, the WebSocket is ok.
	defer wsBuff.Flush(websocket.TextMessage, ws)

	done := make(chan bool, 3)
	setDone := func() { done <- true }

	// most messages are ssh output, not webSocket input
	writeMessageToSSHServer := func(wc io.WriteCloser) { // read messages from webSocket
		defer setDone()
		for {
			msgType, p, err := ws.ReadMessage()
			// if WebSocket is closed by some reason, then this func will return,
			// and 'done' channel will be set, the outer func will reach to the end.
			// then ssh session will be closed in defer.
			if err != nil {
				log.Println("Error: error reading webSocket message:", err)
				return
			}
			//ssh 和 ws 交互
			if err = DispatchMessage(sshConn.Session, msgType, p, wc); err != nil {
				log.Println("Error: error write data to ssh server:", err)
				return
			}
		}
	}

	stopper := make(chan bool) // timer stopper
	// check webSocketWriterBuffer(if not empty,then write back to webSocket) every 120 ms.
	writeBufferToWebSocket := func() {
		defer setDone()
		//TODO: buffer_checker_cycle_time
		tick := time.NewTicker(time.Millisecond * time.Duration(120))
		//for range time.Tick(120 * time.Millisecond){}
		defer tick.Stop()
		for {
			select {
			case <-tick.C:
				if err := wsBuff.Flush(websocket.TextMessage, ws); err != nil {
					log.Println("Error: error sending data via webSocket:", err)
					return
				}
			case <-stopper:
				return
			}
		}
	}

	go writeMessageToSSHServer(sshConn.StdinPipe)
	go writeBufferToWebSocket()
	go func() {
		defer setDone()
		if err := sshConn.Session.Wait(); err != nil {
			log.Println("ssh exist from server", err)
		}
		// if ssh is closed (wait returns), then 'done', web socket will be closed.
		// by the way, buffered data will be flushed before closing WebSocket.
	}()

	<-done
	stopper <- true // stop tick timer(if tick is finished by due to the bad WebSocket, this line will just only set channel(no bad effect). )
	log.Println("Info: websocket finished!")
}
