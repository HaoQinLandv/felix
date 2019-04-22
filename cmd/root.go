package cmd

import (
	"fmt"
	"github.com/dejavuzhou/felix/models"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "felix",
	Short: "",
	Long:  ``,
	//Run: func(cmd *cobra.Command, args []string) {
	//	if len(args) == 0 {
	//		flx.AllMachines("")
	//		cmd.Help()
	//	}
	//},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var verbose bool

func init() {
	cobra.OnInitialize(initFunc)
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "verbose")
}

func initFunc() {
	models.CreateSqliteDB(verbose)
	//intializeSlackLogrus()
}
