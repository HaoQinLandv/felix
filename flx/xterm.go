package flx

import (
	"github.com/dejavuzhou/felix/models"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
	"io"
	"net/http"
)

type windowSize struct {
	Rows uint16 `json:"rows"`
	Cols uint16 `json:"cols"`
	X    uint16
	Y    uint16
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func RunSshXtermJs(h *models.Machine, sudoMode bool, w http.ResponseWriter, r *http.Request) error {
	_, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.WithError(err).Error("Unable to upgrade connection")
		return err
	}

	client := newSshClient(h)
	defer client.Close()
	session, err := client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()
	if h.Password == "" {
		sudoMode = false
	}
	t := SSHTerminal{
		Session:            session,
		Password:           h.Password,
		EnableSudoPassword: sudoMode,
	}

	termWidth, termHeight := 640, 480

	err = t.Session.RequestPty("xterm-256color", termHeight, termWidth, ssh.TerminalModes{})
	if err != nil {
		return err
	}

	//t.updateTerminalSize()

	t.stdin, err = t.Session.StdinPipe()
	if err != nil {
		return err
	}
	t.stdout, err = t.Session.StdoutPipe()
	if err != nil {
		return err
	}
	t.stderr, err = t.Session.StderrPipe()

	go io.Copy(w, t.stderr)
	go io.Copy(w, t.stdout)
	go io.Copy(t.stdin, r.Body)

	err = t.Session.Shell()
	if err != nil {
		return err
	}
	return t.Session.Wait()
}
