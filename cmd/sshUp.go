package cmd

import (
	"github.com/dejavuzhou/felix/flx"
	"github.com/dejavuzhou/felix/models"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "sshup",
	Short: "scp upload",
	Long: `usage: felix sshup 1 --remote=/data/temp --local=/c/Users/Felix/go/src/github.com/dejavuzhou/felix

`,
	Run: func(cmd *cobra.Command, args []string) {
		dbId, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			log.Fatal("ID must be an integer", err)
		}
		h, err := models.MachineFind(uint(dbId))
		if err != nil {
			log.Fatal("ssh info is not found:", err)
		}
		err = flx.ScpLR(h, localPath, remotePath)
		if err != nil {
			log.Fatal("scp up failed:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)
	uploadCmd.Flags().StringVarP(&remotePath, "remote", "r", "", "remote path")
	uploadCmd.Flags().StringVarP(&localPath, "local", "l", "", "download local path")
	uploadCmd.MarkFlagRequired("remote")
	uploadCmd.MarkFlagRequired("local")
}
