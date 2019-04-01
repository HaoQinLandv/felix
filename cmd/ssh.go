package cmd

import (
	"fmt"
	"github.com/dejavuzhou/felix/flx"
	"github.com/dejavuzhou/felix/models"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

// sshCmd represents the ssh command
var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "登陆SSH服务",
	Long: `
usage: felix ssh 1
`,

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
		if err := flx.RunSshTerminal(h, enableSudoMode); err != nil {
			fmt.Println(err)
		}
	},
}
var enableSudoMode bool

func init() {
	rootCmd.AddCommand(sshCmd)
	sshCmd.Flags().BoolVarP(&enableSudoMode, "sudo", "s", true, "sudo模式:自动帮助你输sudo的密码,默认开启")
}
