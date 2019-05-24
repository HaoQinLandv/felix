package flx

import (
	"github.com/dejavuzhou/felix/models"
	"github.com/pkg/sftp"
	"path/filepath"
)

const maxPacket = 1 << 15

func NewSftpClient(h *models.Machine) (*sftp.Client, error) {
	conn, err := NewSshClient(h)
	if err != nil {
		return nil, err
	}
	return sftp.NewClient(conn, sftp.MaxPacket(maxPacket))
}
func toUnixPath(path string) string {
	return filepath.Clean(path)
}
