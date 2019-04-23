package cmd

import (
	"fmt"
	"github.com/dejavuzhou/felix/models"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
	"runtime"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "felix",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if isShowVersion {
			color.HiYellow("Golang Env: %s %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH)
			color.Cyan("UTC build time:%s", buildTime)
			color.Yellow("Build from Github repo version: https://github.com/dejavuzhou/felix/commit/%s", gitHash)
		}
		//if len(args) == 0 {
		//	flx.AllMachines("")
		//	cmd.Help()
		//}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(bTime, gHash string) {
	buildTime = bTime
	gitHash = gHash
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var buildTime, gitHash string
var verbose, isShowVersion bool

func init() {
	cobra.OnInitialize(initFunc)
	rootCmd.Flags().BoolVarP(&isShowVersion, "version", "V", false, "show binary build information")
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "verbose")
}

func initFunc() {
	models.CreateSqliteDB(verbose)
	//intializeSlackLogrus()
}
