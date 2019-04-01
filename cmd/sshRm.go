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
	"strconv"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "sshrm",
	Short: "删除SSH服务器",
	Long:  `删除ID为2的SSH服务器 felix del 2`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			cmd.Help()
			fmt.Println("ID参数必须为正整数")
		}
		if err := models.MachineDelete(uint(id)); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
}
