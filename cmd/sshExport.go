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
	"encoding/csv"
	"github.com/dejavuzhou/felix/models"
	"github.com/fatih/color"
	"github.com/mitchellh/go-homedir"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// sshExportCmd represents the sshexport command
var sshExportCmd = &cobra.Command{
	Use:   "sshexport",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if err := exportAllSshInfoToHomeDirCsvFile(); err != nil {
			log.Fatal(err)
		}
		color.Cyan("you can use export all ssh rows")
		color.Cyan("then edit them in excel")
		color.Cyan("use the file to felix sshimport with F flag")
		color.Yellow("update all ssh info massively")
	},
}

func init() {
	rootCmd.AddCommand(sshExportCmd)
}

func exportAllSshInfoToHomeDirCsvFile() error {
	mcs, err := models.MachineAll("")
	if err != nil {
		return err
	}
	filePath, _ := homedir.Expand("~/allSshInfo.csv")
	csvFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)
	rows := [][]string{
		{"ssh_user(optional can be blank)", "ssh_password(optional can be blank)", "ssh_name", "ssh_host", "ssh_port (optional can be blank)"},
	}
	for _, mc := range mcs {
		one := []string{mc.User, mc.Password, mc.Name, mc.Host, strconv.Itoa(int(mc.Port))}
		rows = append(rows, one)
	}

	err = csvWriter.WriteAll(rows)
	if err != nil {
		return err
	}
	color.Cyan("ssh import csv template has exported into %s", filePath)
	color.Yellow("use Excel to add ssh info into a row")
	return nil
}
