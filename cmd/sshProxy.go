package cmd

import (
	"fmt"
	"github.com/dejavuzhou/felix/flx"
	"github.com/dejavuzhou/felix/models"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

// proxyCmd represents the proxy command
var proxyCmd = &cobra.Command{
	Use:   "sshProxy",
	Short: "SSH隧道代理: felix proxy ID -l -r",
	Long:  `代理ssh服务器上127.0.0.1:3306 到本地 127.0.0.1:5555,这样本地就可以通过5555端口访问ssh服务器上3306端口的数据库: felix proxy 2 -l 127.0.0.1:5555 -r 127.0.0.1:3306`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			flx.AllMachines("")
			return
		}
		dbId, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			log.Fatal("服务器ID必须为正整数:", err)
		}
		h, err := models.MachineFind(uint(dbId))
		if err != nil {
			log.Fatal("错误的SSH服务器ID ", err)
		}
		fmt.Printf("代理SSH服务器远程(%s)到本地(%s)...", remoteAddr, localAddr)
		if err := flx.RunProxy(h, localAddr, remoteAddr); err != nil {
			log.Fatal(err)
		}
	},
}
var localAddr, remoteAddr string

func init() {
	rootCmd.AddCommand(proxyCmd)
	proxyCmd.Flags().StringVarP(&localAddr, "local", "l", "127.0.0.1:3306", "ssh 代理到本地addr")
	proxyCmd.Flags().StringVarP(&remoteAddr, "remote", "r", "127.0.0.1:3306", "ssh 代理到本地addr")
}
