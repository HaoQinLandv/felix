package cmd

import (
	"fmt"
	"github.com/dejavuzhou/felix/models"
	"strconv"

	"github.com/spf13/cobra"
)

// hostUpdateCmd represents the hostEdit command
var hostUpdateCmd = &cobra.Command{
	Use:   "sshedit",
	Short: "更新SSH服务器信息",
	Long:  `更新ID:1 SSH服务器的名称 felix update 1 -n=Awesome`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		argId, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			cmd.Help()
			fmt.Println("ID参数必须为正整数:", err)
		}
		if err := models.MachineUpdate(updateName, updateAddr, updateUser, updatePassword, updateKey, updateAuth, uint(argId), updatePort); err != nil {
			fmt.Println(err)
		}
	},
}
var updateKey, updateName, updateAddr, updateUser, updatePassword, updateAuth string
var updatePort uint

func init() {
	rootCmd.AddCommand(hostUpdateCmd)

	hostUpdateCmd.Flags().StringVarP(&updatePassword, "password", "p", "", "SSH登陆密码")
	hostUpdateCmd.Flags().StringVarP(&updateKey, "key", "k", "", "SSH登陆私钥路径")
	hostUpdateCmd.Flags().StringVarP(&updateName, "name", "n", "", "SSH服务器部门")
	hostUpdateCmd.Flags().StringVarP(&updateAddr, "addr", "a", "", "连接字符串")
	hostUpdateCmd.Flags().StringVarP(&updateUser, "user", "u", "", "SSH登陆用户名")
	//hostUpdateCmd.Flags().UintVarP(&id, "id", "i", 0, "SSH服务器ID(部署参数)")
	hostUpdateCmd.Flags().StringVarP(&updateAuth, "auth", "", "", "password密码登陆,key私钥登陆")
	hostUpdateCmd.Flags().UintVar(&updatePort, "port", 0, "SSH服务端口")

}
