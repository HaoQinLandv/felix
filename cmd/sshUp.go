package cmd

import (
	"github.com/dejavuzhou/felix/flx"
	"github.com/dejavuzhou/felix/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "sshup",
	Short: "上传本地文件(目录)到SSH服务器",
	Long: `用法: felix sshup 1 --remote=/data/temp --local=/c/Users/zhouqing1/go/src/github.com/dejavuzhou/felix
	1           : SSH服务器ID,通过 felix ls 命令查询
	--remote(-r):上传到ssh服务器的绝对路径
	--local (-l): 要上传的本地路径
`,
	Run: func(cmd *cobra.Command, args []string) {
		dbId, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			log.Fatal("第一个参数必须为正整数:", err)
		}
		h, err := models.MachineFind(uint(dbId))
		if err != nil {
			log.Fatal("错误的I参数:", err)
		}
		err = flx.ScpLR(h, localPath, remotePath)
		if err != nil {
			logrus.WithError(err).Error("scp up failed")
		}
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)
	uploadCmd.Flags().StringVarP(&remotePath, "remote", "r", "", "远程服务器要下载的绝对路径")
	uploadCmd.Flags().StringVarP(&localPath, "local", "l", "", "保存到本地绝对路径")
}
