package controllers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/dejavuzhou/felix/flx"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

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
	sshConn := sshShellSession{
		SshClient: client,
	}
	// set io for ssh session std out
	var wsBuff wsBufferWriter
	defer wsBuff.Flush(websocket.TextMessage, ws)

	sshConn.WriterPipe = &wsBuff

	defer sshConn.close()
	// config ssh
	cols, err := strconv.Atoi(c.DefaultQuery("cols", "120"))
	if handleError(c, err) {
		return
	}
	rows, err := strconv.Atoi(c.DefaultQuery("rows", "32"))
	if handleError(c, err) {
		return
	}
	err = sshConn.config(cols, rows)
	if handleError(c, err) {
		return
	}

	// an egg:
	//if err := sshEntity.Session.Setenv("SSH_EGG", SSH_EGG); err != nil {
	//	log.Println(err)
	//}
	// after configure, the WebSocket is ok.

	done := make(chan bool, 3)
	setDone := func() { done <- true }

	// most messages are ssh output, not webSocket input
	writeMessageToSSHServer := func(wc io.WriteCloser) { // read messages from webSocket
		defer setDone()
		for {
			_, p, err := ws.ReadMessage()
			// if WebSocket is closed by some reason, then this func will return,
			// and 'done' channel will be set, the outer func will reach to the end.
			// then ssh session will be closed in defer.
			if err != nil {
				log.Println("Error: error reading webSocket message:", err)
				return
			}
			//ssh 和 ws 交互
			if err = dispatchMsg(sshConn.Session, p, wc); err != nil {
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

// copy data from WebSocket to ssh server
// and copy data from ssh server to WebSocket

// write data to WebSocket
// the data comes from ssh server.
type wsBufferWriter struct {
	buffer bytes.Buffer
	mu     sync.Mutex
}

// implement Write interface to write bytes from ssh server into bytes.Buffer.
func (w *wsBufferWriter) Write(p []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.buffer.Write(p)
}

// flush all data in this buff into WebSocket.
func (w *wsBufferWriter) Flush(messageType int, ws *websocket.Conn) error {
	if w.buffer.Len() != 0 {
		err := ws.WriteMessage(messageType, []byte(w.buffer.Bytes()))
		if err != nil {
			return err
		}
		w.buffer.Reset()
	}
	return nil
}

const (
	wsMsgTypeTerminal  = "terminal"
	wsMsgTypeHeartbeat = "heartbeat"
	wsMsgTypeResize    = "resize"
)

type wsMsg struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"` // json.RawMessage
}

// normal terminal message
type shellMsg struct {
	DataBase64 string `json:"base64"`
}

// terminal window resize
type windowResize struct {
	Cols int `json:"cols"`
	Rows int `json:"rows"`
}

func dispatchMsg(sshSession *ssh.Session, wsData []byte, wc io.WriteCloser) error {
	var socketData json.RawMessage
	socketStream := wsMsg{
		Data: &socketData,
	}

	if err := json.Unmarshal(wsData, &socketStream); err != nil {
		return nil // skip error
	}

	switch socketStream.Type {
	case wsMsgTypeHeartbeat:
		return nil
	case wsMsgTypeResize:
		var resize windowResize
		if err := json.Unmarshal(socketData, &resize); err != nil {
			return nil // skip error
		}
		sshSession.WindowChange(resize.Rows, resize.Cols)
	case wsMsgTypeTerminal:
		var message shellMsg
		if err := json.Unmarshal(socketData, &message); err != nil {
			return nil
		}
		if decodeBytes, err := base64.StdEncoding.DecodeString(message.DataBase64); err != nil { // todo ignore error
			return nil // skip error
		} else {
			if _, err := wc.Write(decodeBytes); err != nil {
				return err
			}
		}
	}
	return nil
}

// connect to ssh server using ssh session.
type sshShellSession struct {
	SshClient *ssh.Client
	// calling Write() to write data to ssh server
	StdinPipe io.WriteCloser
	// Write() be called to receive data from ssh server
	WriterPipe io.Writer
	Session    *ssh.Session
}

// setup ssh shell session
// set Session and StdinPipe here,
// and the Session.Stdout and Session.Sdterr are also set.
func (s *sshShellSession) config(cols, rows int) error {
	session, err := s.SshClient.NewSession()
	if err != nil {
		return err
	}
	s.Session = session

	// we set stdin, then we can write data to ssh server via this stdin.
	// but, as for reading data from ssh server, we can set Session.Stdout and Session.Stderr
	// to receive data from ssh server, and write back to somewhere.
	if stdin, err := s.Session.StdinPipe(); err != nil {
		log.Fatal("failed to set IO stdin: ", err)
		return err
	} else {
		// in fact, stdin it is channel.
		s.StdinPipe = stdin
	}

	// set writer, such the we can receive ssh server's data and write the data to somewhere specified by WriterPipe.
	if s.WriterPipe == nil {
		return errors.New("WriterPipe is nil")
	}
	session.Stdout = s.WriterPipe
	session.Stderr = s.WriterPipe

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // disable echo
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}
	// Request pseudo terminal
	if err := session.RequestPty("xterm", rows, cols, modes); err != nil {
		log.Fatal("request for pseudo terminal failed: ", err)
		return err
	}
	// Start remote shell
	if err := session.Shell(); err != nil {
		log.Fatal("failed to start shell: ", err)
		return err
	}
	return nil
}

func (s *sshShellSession) close() {
	if s.Session != nil {
		s.Session.Close()
	}

	if s.SshClient != nil {
		s.SshClient.Close()
	}
}
