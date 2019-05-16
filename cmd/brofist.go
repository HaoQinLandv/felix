package cmd

import (
	"github.com/dejavuzhou/felix/utils"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"time"
)

// broFistCmd represents the brofist command
var broFistCmd = &cobra.Command{
	Use:   "brofist",
	Short: "Pewdiepie needs your help.Do your part to subscribe Pewdiepie's Youtube Channel.",
	Long:  `PewDiePie vs T-Series`,
	Run: func(cmd *cobra.Command, args []string) {
		color.Red("Pewdiepie needs your help.")
		color.Yellow("Do your part to subscribe Felix's Youtube Channel.")
		color.Cyan("Your browser will go to Pewdiepie's Channel.")
		color.Blue("Please click the subscribe button on the right.")
		time.Sleep(time.Second * 2)
		utils.BrowserOpen("https://www.youtube.com/channel/UC-lHJZR3Gqxm24_Vd_AJ5Yw")
	},
}

func init() {
	rootCmd.AddCommand(broFistCmd)
}
