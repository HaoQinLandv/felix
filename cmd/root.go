package cmd

import (
	"fmt"
	"github.com/dejavuzhou/felix/flx"
	"github.com/dejavuzhou/felix/models"
	"github.com/johntdyer/slackrus"
	"github.com/sirupsen/logrus"
	"os"

	"github.com/spf13/cobra"
)

var language = "zh_cn"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "felix",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			flx.AllMachines("")
			cmd.Help()
		}
	},
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
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "显示日志")
}

func initFunc() {
	models.CreateSqliteDB(verbose)
	//intializeSlackLogrus()
}
func intializeSlackLogrus() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	//file, _ := os.Create(time.Now().Format("2006_01_02.log"))
	//logrus.SetOutput(file)

	logrus.SetLevel(logrus.DebugLevel)

	logrus.AddHook(&slackrus.SlackrusHook{
		HookURL:        "https://hooks.slack.com/services/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",
		AcceptedLevels: slackrus.LevelThreshold(logrus.DebugLevel),
		Channel:        "#felix",
		IconEmoji:      ":shark:",
		Username:       "FelixZhou",
	})
	//TODO::get langu
}
