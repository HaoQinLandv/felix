package flx

import (
	"bufio"
	"fmt"
	"github.com/dejavuzhou/felix/models"
	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"os"
	"strings"
)

func newSshClient(h *models.Machine) *ssh.Client {
	config := &ssh.ClientConfig{
		User:            h.User,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //这个可以， 但是不够安全
		//HostKeyCallback: hostKeyCallBackFunc(h.Host),
	}
	if h.Type == "password" {
		config.Auth = []ssh.AuthMethod{ssh.Password(h.Password)}
	} else {
		config.Auth = []ssh.AuthMethod{publicKeyAuthFunc(h.Key)}
	}
	addr := fmt.Sprintf("%s:%d", h.Host, h.Port)
	c, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		logrus.WithError(err).Fatal("ssh dial failed")
	}
	return c
}
func hostKeyCallBackFunc(host string) ssh.HostKeyCallback {
	hostPath, err := homedir.Expand("~/.ssh/known_hosts")
	if err != nil {
		logrus.WithError(err).Fatal("find known_hosts's home dir failed")
	}
	file, err := os.Open(hostPath)
	if err != nil {
		logrus.WithError(err).Fatal("can't find known_host file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var hostKey ssh.PublicKey
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		if len(fields) != 3 {
			continue
		}
		if strings.Contains(fields[0], host) {
			var err error
			hostKey, _, _, _, err = ssh.ParseAuthorizedKey(scanner.Bytes())
			if err != nil {
				logrus.WithError(err).Fatalf("error parsing %q: %v", fields[2], err)
			}
			break
		}
	}
	if hostKey == nil {
		logrus.WithError(err).Fatalf("no hostkey for %s", host)
	}
	return ssh.FixedHostKey(hostKey)
}

func publicKeyAuthFunc(kPath string) ssh.AuthMethod {
	keyPath, err := homedir.Expand(kPath)
	if err != nil {
		logrus.WithError(err).Fatal("find key's home dir failed")
	}
	key, err := ioutil.ReadFile(keyPath)
	if err != nil {
		logrus.WithError(err).Fatal("ssh key file read failed")
	}
	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		logrus.WithError(err).Fatal("ssh key signer failed")
	}
	return ssh.PublicKeys(signer)
}
