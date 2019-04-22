package cmd

import (
	"github.com/dejavuzhou/felix/models"
	"github.com/prometheus/common/log"
	"github.com/spf13/cobra"
	"strconv"
)

// taskdnCmd represents the taskdn command
var taskdnCmd = &cobra.Command{
	Use:   "taskok",
	Short: "set a task done",
	Long:  `usage:felix taskok 1`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			log.Fatal("ID must be an integer")
		}
		if err := models.TaskUpdate(uint(id), "DONE"); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(taskdnCmd)
}
