package cmd

import (
	"github.com/dejavuzhou/felix/models"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// taskadCmd represents the taskad command
var taskadCmd = &cobra.Command{
	Use:   "taskad",
	Short: "add a task",
	Long:  `usage: felix taskad 'subscribe Pewdiepie's Youtube channel`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		content := args[0]
		err := models.TaskAdd(content, category, deadline)
		if err != nil {
			color.Red("%s", err)
		}
	},
}
var deadline, category string

func init() {
	rootCmd.AddCommand(taskadCmd)

	taskadCmd.Flags().StringVarP(&category, "category", "c", "defaul", "task category")
	taskadCmd.Flags().StringVarP(&deadline, "deadline", "d", "", "task deadline")
}
