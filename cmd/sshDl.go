package cmd

import (
	"github.com/dejavuzhou/felix/flx"
	"github.com/dejavuzhou/felix/models"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "sshDl",
	Short: "从服务器上下载文件",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dbId, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			log.Fatal("第一个参数必须为正整数:", err)
		}
		h, err := models.MachineFind(uint(dbId))
		if err != nil {
			log.Fatal("错误的I参数:", err)
		}
		err = flx.ScpRL(h, remotePath, localPath)
		if err != nil {
			log.Println(err)
		}
	},
}
var localPath, remotePath string

func init() {
	rootCmd.AddCommand(downloadCmd)
	downloadCmd.Flags().StringVarP(&remotePath, "remote", "r", "", "远程服务器要下载的绝对路径")
	downloadCmd.Flags().StringVarP(&localPath, "local", "l", "", "保存到本地绝对路径")

}
