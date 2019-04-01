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
	"bufio"
	"github.com/dejavuzhou/felix/models"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// testImportCmd represents the testImport command
var testImportCmd = &cobra.Command{
	Use:   "sshimport",
	Short: "批量导入SSH服务器",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		importHost()
	},
}
var imPassword, imFile, imUser, imKey, imAuth string

func init() {
	rootCmd.AddCommand(testImportCmd)
	testImportCmd.Flags().StringVarP(&imFile, "file", "f", ``, "SSH服务器文本文件一行就是一个服务器")
	testImportCmd.Flags().StringVarP(&imPassword, "password", "p", "", "导入密码")
	testImportCmd.Flags().StringVarP(&imUser, "user", "u", "", "导入用户名")
	testImportCmd.Flags().StringVarP(&imKey, "key", "k", "~/.ssh/id_rsa", "SSH Private Key")
	testImportCmd.Flags().StringVarP(&imAuth, "auth", "", "password", "SSH验证类型 passwor key")
}

func importHost() {
	file, err := os.Open(imFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(line, "#") {
			cells := strings.Fields(line)
			if len(cells) >= 2 {
				addrT := cells[0]
				ip := cells[1]
				models.MachineAdd(addrT, addrT, ip, imUser, imPassword, imKey, imAuth, 22)
			}

		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
