package cmd

import (
	"github.com/dejavuzhou/felix/models"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"strconv"
)

// taskrmCmd represents the taskrm command
var taskrmCmd = &cobra.Command{
	Use:   "taskRm",
	Short: "remove one task from task list",
	Long:  `usage: felix taskrm 1`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			color.Red("第一个参数必须为正整数:[%s]", err)
		}
		if err := models.TaskRm(uint(id)); err != nil {
			color.Red("%s", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(taskrmCmd)
}
