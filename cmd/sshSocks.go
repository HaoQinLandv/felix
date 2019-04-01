package cmd

import (
	"github.com/dejavuzhou/felix/flx"
	"github.com/dejavuzhou/felix/models"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

// proxyCmd represents the proxy command
var proxySocksCmd = &cobra.Command{
	Use:   "sshsocks",
	Short: "SSH隧道SOCKS代理: felix proxy socks ID",
	Long:  `把目标SSH服务器ID:2 作为SOCKS代理 felix proxy socks 2 --localPort=1080`,
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
		if err := flx.RunSocksProxy(h, localPort); err != nil {
			log.Fatal(err)
		}
	},
}
var localPort int

func init() {
	proxyCmd.AddCommand(proxySocksCmd)
	proxySocksCmd.Flags().IntVarP(&localPort, "localPort", "l", 1080, "socks4/5 代理到本地[127.0.0.1]本地端口")
}
