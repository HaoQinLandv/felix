package cmd

import (
	"fmt"
	"github.com/dejavuzhou/felix/models"
	"github.com/spf13/cobra"
)

// hostResetCmd represents the hostReset command
var hostResetCmd = &cobra.Command{
	Use:   "sshflush",
	Short: "清空SSH服务器全部记录",
	Long:  `删除~/.felix.db SQLite数据库,下次运行命令重建数据库`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := models.FlushSqliteDb(); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(hostResetCmd)
}
