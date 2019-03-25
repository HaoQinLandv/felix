package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// taskdnCmd represents the taskdn command
var taskdnCmd = &cobra.Command{
	Use:   "taskOk",
	Short: "mark a task has been done",
	Long:  `usage:felix taskdn 1`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("taskdn called")
	},
}

func init() {
	rootCmd.AddCommand(taskdnCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// taskdnCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// taskdnCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
