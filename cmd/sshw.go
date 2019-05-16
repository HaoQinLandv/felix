// Copyright © 2019 Eric Freeman Zhou <neochau@qq.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"github.com/dejavuzhou/felix/ssh2ws"
	"github.com/dejavuzhou/felix/utils"
	"github.com/spf13/cobra"
	"log"
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
		//time.AfterFunc(time.Second*3, func() {
		//	if err = utils.BrowserOpen(fmt.Sprintf("http://localhost%s", bindAddress)); err != nil {
		//		log.Println(err)
		//	}
		//})
		if err := ssh2ws.RunSsh2ws(addr, user, password, ex, []byte(secret)); err != nil {
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
