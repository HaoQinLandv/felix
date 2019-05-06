package flx

import (
	"github.com/dejavuzhou/felix/models"
	"github.com/pkg/sftp"
	"log"
	"path/filepath"
)

const maxPacket = 1 << 15

func NewSftpClient(h *models.Machine) *sftp.Client {
	conn, err := newSshClient(h)
	if err != nil {
		log.Fatal("create ssh client failed", err)
	}
	c, err := sftp.NewClient(conn, sftp.MaxPacket(maxPacket))
	if err != nil {
		log.Fatal("create sftp client failed", err)
	}
	return c
}
func toUnixPath(path string) string {
	return filepath.Clean(path)
}
