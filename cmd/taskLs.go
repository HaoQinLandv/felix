package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// tasklsCmd represents the taskls command
var tasklsCmd = &cobra.Command{
	Use:   "taskLs",
	Short: "display all task in a table",
	Long:  `usage:felix taskls`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("taskls called")
	},
}

func init() {
	rootCmd.AddCommand(tasklsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tasklsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tasklsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
