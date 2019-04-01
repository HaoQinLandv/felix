package cmd

import (
	"github.com/dejavuzhou/felix/models"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"strconv"
)

// taskdnCmd represents the taskdn command
var taskdnCmd = &cobra.Command{
	Use:   "taskok",
	Short: "设置reminder中一条任务完成",
	Long:  `usage:felix taskok 1;设置ID为1的任务完成`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			color.Red("第一个参数必须为正整数:[%s]", err)
		}
		if err := models.TaskUpdate(uint(id), "DONE"); err != nil {
			color.Red("%s", err)
		}
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
