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
	"fmt"
	"github.com/dejavuzhou/felix/models"
	"github.com/spf13/cobra"
)

// hostResetCmd represents the hostReset command
var hostResetCmd = &cobra.Command{
	Use:   "sshFlush",
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
