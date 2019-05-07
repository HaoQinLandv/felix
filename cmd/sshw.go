// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"github.com/dejavuzhou/felix/ssh2ws"
	"github.com/spf13/cobra"
)

// sshwCmd represents the sshw command
var sshwCmd = &cobra.Command{
	Use:   "sshw",
	Short: "open ssh terminal in web browser",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		ssh2ws.RunSsh2ws(addr, user, password, isDev)
	},
}
var isDev bool

func init() {
	rootCmd.AddCommand(sshwCmd)
	sshwCmd.Flags().BoolVarP(&isDev, "dev", "d", false, "is development mode")
	sshwCmd.Flags().StringVarP(&addr, "addr", "a", ":2222", "listening addr")
	sshwCmd.Flags().StringVarP(&user, "user", "u", "", "auth user")
	sshwCmd.Flags().StringVarP(&password, "password", "p", "", "auth password")
}
