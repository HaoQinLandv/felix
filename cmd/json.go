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
}
