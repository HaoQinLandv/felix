package cmd

import (
	"github.com/dejavuzhou/felix/webshell"
	"github.com/spf13/cobra"
)

// xtermCmd represents the xterm command
var xtermCmd = &cobra.Command{
	Use:   "xterm",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		webshell.RunXterm()
	},
}

func init() {
	rootCmd.AddCommand(xtermCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// xtermCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// xtermCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
