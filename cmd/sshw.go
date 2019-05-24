package cmd

import (
	"fmt"
	"github.com/dejavuzhou/felix/ssh2ws"
	"github.com/dejavuzhou/felix/utils"
	"github.com/spf13/cobra"
	"log"
	"runtime"
	"strings"
	"time"
)

// sshwCmd represents the sshw command
var sshwCmd = &cobra.Command{
	Use:   "sshw",
	Short: "open a web UI for Felix http://localhost:2222",
	Long:  `the demo website is http://home.mojotv.cn:2222`,
	Run: func(cmd *cobra.Command, args []string) {
		if secret == "" {
			secret = utils.RandomString(32)
			fmt.Printf("use random string as jwt secret: %s\n", secret)
		}
		ex := time.Minute * time.Duration(expire)

		fmt.Println("login user:", user)
		fmt.Println("login password:", password)
		fmt.Printf("login expire in %d minutes\n", expire)

		sl := strings.Split(addr, ":")
		if len(sl) == 2 && (runtime.GOOS == "windows" || runtime.GOOS == "darwin") {
			time.AfterFunc(time.Second*3, func() {
				if err := utils.BrowserOpen(fmt.Sprintf("http://localhost:%s", sl[1])); err != nil {
					log.Println(err)
				}
			})
		}

		if err := ssh2ws.RunSsh2ws(addr, user, password, secret, ex, verbose); err != nil {
			log.Fatal(err)
		}
	},
}
var expire uint
var secret string

func init() {
	rootCmd.AddCommand(sshwCmd)
	sshwCmd.Flags().StringVarP(&secret, "secret", "s", "", "jwt secret string")
	sshwCmd.Flags().StringVarP(&addr, "addr", "a", ":2222", "listening addr")
	sshwCmd.Flags().StringVarP(&user, "user", "u", "admin", "auth user")
	sshwCmd.Flags().StringVarP(&password, "password", "p", "admin", "auth password")
	sshwCmd.Flags().UintVarP(&expire, "expire", "x", 60*24, "token expire in * minute")
}
