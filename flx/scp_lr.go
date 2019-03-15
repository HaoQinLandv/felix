package flx

import (
	"github.com/dejavuzhou/felix/models"
	"github.com/pkg/sftp"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func ScpLR(h *models.Machine, localPath, remotePath string) error {
	c := newSftpClient(h)
	defer c.Close()
	return lrCopy(localPath, remotePath, c)
}
func lrCopy(local, remote string, c *sftp.Client) error {
	info, err := os.Lstat(local)
	if err != nil {
		return err
	}
	//if !info.Mode().IsRegular() {
	//	return fmt.Errorf("not support irregular file")
	//}
	if info.Mode()&os.ModeSymlink != 0 {
		return lrCopyL(local, remote, c)
	}
	if info.IsDir() {
		return lrCopyD(local, remote, c)
	}
	return lrCopyF(local, remote, info, c)
}
func lrCopyL(local, remote string, c *sftp.Client) error {
	realLocal, err := os.Readlink(local)
	if err != nil {
		return err
	}
	return lrCopy(realLocal, remote, c)
}
func lrCopyD(local, remote string, c *sftp.Client) error {
	//if err := c.MkdirAll(remote); err != nil {
	//	logrus.WithError(err).Error("scp mkdir failed")
	//	return err
	//}
	//if err := c.Chmod(remote,info.Mode()); err != nil {
	//	logrus.WithError(err).Error("scp chmod dir failed")
	//	return err
	//}
	contents, err := ioutil.ReadDir(local)
	if err != nil {
		logrus.WithError(err).Error("ioutil read local dir failed")
		return err
	}
	for _, content := range contents {
		cs, cd := filepath.Join(local, content.Name()), filepath.Join(remote, content.Name())
		if err := lrCopy(cs, cd, c); err != nil {
			logrus.WithError(err).Error(cs, cd)
			return err
		}
	}
	return nil
}
func lrCopyF(local, remote string, info os.FileInfo, c *sftp.Client) error {
	localFile, err := os.Open(local)
	if err != nil {
		logrus.WithError(err).Error("BrowserOpen local file failed")
		return err
	}
	defer localFile.Close()
	err = c.MkdirAll(toUnixPath(filepath.Dir(remote)))
	if err != nil {
		logrus.WithError(err).Error("scp mkdir all failed")
		return err
	}
	remoteFile, err := c.Create(toUnixPath(remote))
	if err != nil {
		logrus.WithError(err).WithField("path", info).Error("create remote file failed:", remote)
		return err
	}
	defer remoteFile.Close()
	err = c.Chmod(remoteFile.Name(), info.Mode())
	if err != nil {
		logrus.WithError(err).Error("scp chmod failed")
		return err
	}
	_, err = io.Copy(remoteFile, localFile)
	if err != nil {
		logrus.WithError(err).Error("io copy failed")
		return err
	}
	return nil
}
