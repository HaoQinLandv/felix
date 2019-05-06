package utils

import (
	"errors"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
)

const (
	SSH_IO_MODE_CHANNEL = 0
	SSH_IO_MODE_SESSION = 1
)

type SSHConnInterface interface {
	// close ssh connection
	Close()
	// connect using username and password
	Connect(username, password string) error
	// config connection after connected
	Config(cols, rows int) error
}

// connect to ssh server using ssh session.
type SSHShellSession struct {
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
func (s *SSHShellSession) Config(cols, rows int) error {
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

func (s *SSHShellSession) Close() {
	if s.Session != nil {
		s.Session.Close()
	}

	if s.SshClient != nil {
		s.SshClient.Close()
	}
}

type ptyRequestMsg struct {
	Term     string
	Columns  uint32
	Rows     uint32
	Width    uint32
	Height   uint32
	Modelist string
}
