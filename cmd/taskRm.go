package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// taskrmCmd represents the taskrm command
var taskrmCmd = &cobra.Command{
	Use:   "taskRm",
	Short: "remove one task from task list",
	Long:  `usage: felix taskrm 1`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("taskrm called")
	},
}

func init() {
	rootCmd.AddCommand(taskrmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// taskrmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// taskrmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
