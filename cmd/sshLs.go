package cmd

import (
	"github.com/dejavuzhou/felix/flx"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "sshLs",
	Short: "查看全部的SSH服务器",
	Long:  `查看全部的SSH服务器: felix ls`,
	Run: func(cmd *cobra.Command, args []string) {
		flx.AllMachines(searchKey)
	},
}
var searchKey string

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	listCmd.Flags().StringVarP(&searchKey, "search", "s", "", "模糊搜索ssh服务器名称")
}
