package cmd

import (
	"github.com/dejavuzhou/felix/models"
	"github.com/spf13/cobra"
	"log"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "sshadd",
	Short: "add a ssh connection configuration",
	Long:  `add a ssh connection,usage: felix sshadd -p my_password -k ~/.ssh/id_rsa -n mySSH -a 192.168.0.01:22 -u root --auth=key`,
	Run: func(cmd *cobra.Command, args []string) {
		if authType != "password" || authTable != "key" {
			log.Fatalln("--auth flag's value must be password or key")
		}
		if err := models.MachineAdd(name, addr, "", user, password, key, authType, port); err != nil {
			log.Fatal(err)
		}
	},
}

var key, name, addr, ip, user, password, authType string
var port uint

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&password, "password", "p", "", "ssh login password")
	addCmd.Flags().StringVarP(&key, "key", "k", "~/.ssh/id_rsa", "ssh login private key path eg:~/.ssh/id_rsa")
	addCmd.Flags().StringVarP(&name, "name", "n", "", "ssh server name, name whatever you want")
	addCmd.Flags().StringVarP(&addr, "addr", "a", "", "ssh server's domain or ip")
	addCmd.Flags().UintVar(&port, "port", 22, "ssh port")
	addCmd.Flags().StringVarP(&user, "user", "u", "", "ssh login user name")
	addCmd.Flags().StringVarP(&authType, "auth", "", "password", "ssh auth type, only alows 'password' and 'key'")
	addCmd.MarkFlagRequired("addr")
	addCmd.MarkFlagRequired("name")
}
