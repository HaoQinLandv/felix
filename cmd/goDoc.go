package cmd

import (
	"github.com/dejavuzhou/felix/flx"
	"github.com/fatih/color"
	"time"

	"github.com/spf13/cobra"
)

// godocCmd represents the godoc command
var godocCmd = &cobra.Command{
	Use:   "goDoc",
	Short: "golang.google.cn/pkg",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		color.Blue("访问Go中国官方网站 https://golang.google.cn/pkg/")
		time.Sleep(time.Second * 1)
		flx.BrowserOpen("https://golang.google.cn/pkg/")
	},
}

func init() {
	rootCmd.AddCommand(godocCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// godocCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// godocCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
