package cmd

import (
	"fmt"
	"github.com/dejavuzhou/felix/models"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "sshadd",
	Short: "添加SSH服务器",
	Long:  `实例: felix sshadd -p my_password -k ~/.ssh/id_rsa -n mySSH -a 192.168.0.01:22 -u root --auth=key`,
	Run: func(cmd *cobra.Command, args []string) {
		if addr == "" || user == "" {
			fmt.Println("addr user 不能为空")
			cmd.Help()
			return
		}
		if err := models.MachineAdd(name, addr, "", user, password, key, authType, port); err != nil {
			fmt.Println(err)
		}
	},
}

var key, name, addr, ip, user, password, authType string
var port uint

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&password, "password", "p", "", "SSH登陆密码")
	addCmd.Flags().StringVarP(&key, "key", "k", "~/.ssh/id_rsa", "SSH私钥路径默认:~/.ssh/id_rsa")
	addCmd.Flags().StringVarP(&name, "name", "n", "sshName", "SSH服务器名称")
	addCmd.Flags().StringVarP(&addr, "addr", "a", "", "ssh服务器域名或者IP")
	addCmd.Flags().UintVar(&port, "port", 22, "端口默认22")
	addCmd.Flags().StringVarP(&user, "user", "u", "", "SSH登陆用户名")
	addCmd.Flags().StringVarP(&authType, "auth", "", "password", "认证类型:password-(默认)密码,key-私钥登陆")
}
