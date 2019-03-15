package flx

import (
	"github.com/dejavuzhou/felix/models"
	"github.com/pkg/sftp"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
)

func ScpRL(h *models.Machine, remotePath, localPath string) error {
	c := newSftpClient(h)
	defer c.Close()
	return rlCopy(remotePath, localPath, c)
}

func rlCopy(remote, local string, c *sftp.Client) error {
	info, err := c.Lstat(toUnixPath(remote))
	if err != nil {
		return err
	}
	//if !info.Mode().IsRegular() {
	//	return fmt.Errorf("not support irregular file")
	//}
	if info.Mode()&os.ModeSymlink != 0 {
		return rlCopyL(local, remote, c)
	}
	if info.IsDir() {
		return rlCopyD(remote, local, c)
	}
	return rlCopyF(remote, local, info, c)
}
func rlCopyL(remote, local string, c *sftp.Client) error {
	realRemote, err := c.ReadLink(toUnixPath(remote))
	if err != nil {
		return err
	}
	return lrCopy(realRemote, local, c)
}
func rlCopyD(remote, local string, c *sftp.Client) error {
	contents, err := c.ReadDir(toUnixPath(remote))
	if err != nil {
		logrus.WithError(err).Error("ioutil read scp remote dir failed")
		return err
	}
	logrus.Info("dir size:", len(contents))
	for _, info := range contents {
		cdL, csR := filepath.Join(local, info.Name()), filepath.Join(remote, info.Name())
		//mkdir local dir by remote
		err := os.MkdirAll(filepath.Dir(cdL), info.Mode())
		if err != nil {
			logrus.WithError(err).Error("os local sub mkdir all failed")
			return err
		}
		csR = toUnixPath(csR)
		if err := rlCopy(csR, cdL, c); err != nil {
			logrus.WithError(err).Errorf("dir walk remote:%s, local:%s", csR, cdL)
			return err
		}
	}
	return nil
}
func rlCopyF(remote, local string, info os.FileInfo, c *sftp.Client) error {
	rFile, err := c.Open(toUnixPath(remote))
	if err != nil {
		logrus.WithError(err).Error("BrowserOpen scp remote file failed")
		return err
	}
	defer rFile.Close()

	lFile, err := os.Create(local)
	if err != nil {
		logrus.WithError(err).Errorf("os create local file failed:%s", local)
		return err
	}
	defer lFile.Close()

	size, err := io.Copy(lFile, rFile)
	if err != nil {
		logrus.WithError(err).Errorf("io copy remote to local failed.size:%d", size)
		return err
	}

	err = os.Chmod(lFile.Name(), info.Mode())
	if err != nil {
		logrus.WithError(err).Error("os local chmod failed")
		return err
	}
	return nil
}
