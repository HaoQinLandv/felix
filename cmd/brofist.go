package cmd

import (
	"github.com/dejavuzhou/felix/flx"
	"github.com/fatih/color"
	"time"

	"github.com/spf13/cobra"
)

// brofistCmd represents the brofist command
var brofistCmd = &cobra.Command{
	Use:   "brofist",
	Short: "Pewdiepie needs your help.Do your part to subscribe Felix's Youtube Channel.",
	Long:  `PewDiePie vs T-Series`,
	Run: func(cmd *cobra.Command, args []string) {
		color.Red("Pewdiepie needs your help.")
		color.Yellow("Do your part to subscribe Felix's Youtube Channel.")
		color.Cyan("Your browser will go to Pewdiepie's Channel.")
		color.Blue("Please click the subscribe button on the right.")
		time.Sleep(time.Second * 2)
		flx.BrowserOpen("https://www.youtube.com/channel/UC-lHJZR3Gqxm24_Vd_AJ5Yw")
	},
}

func init() {
	rootCmd.AddCommand(brofistCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// brofistCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// brofistCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
