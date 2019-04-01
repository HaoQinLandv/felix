package cmd

import (
	"github.com/dejavuzhou/felix/models"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"strconv"
)

// hostCpCmd represents the hostCp command
var hostCpCmd = &cobra.Command{
	Use:   "sshdu",
	Short: "复制(duplicate)一行ssh登陆信息,提供sshedit使用",
	Long:  `felix sshdu 1`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		argId, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			cmd.Help()
			color.Yellow("ID参数必须为正整数:", err)
		}
		if err := models.MachineDuplicate(uint(argId)); err != nil {
			color.Red("%s", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(hostCpCmd)
}
