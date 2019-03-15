package cmd

import (
	"github.com/dejavuzhou/felix/flx"
	"github.com/spf13/cobra"
)

// jsonCmd represents the json command
var jsonCmd = &cobra.Command{
	Use:   "json",
	Short: "开启浏览器json2struct工具",
	Long:  `打开浏览器(app.quicktype.io)json2struct工具`,
	Run: func(cmd *cobra.Command, args []string) {
		flx.BrowserOpen("https://app.quicktype.io/")
	},
}

func init() {
	rootCmd.AddCommand(jsonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// jsonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// jsonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
