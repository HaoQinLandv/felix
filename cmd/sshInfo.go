package cmd

import (
	"github.com/dejavuzhou/felix/flx"
	"github.com/dejavuzhou/felix/models"
	"github.com/fatih/color"
	"github.com/mattn/go-isatty"
	"github.com/olekukonko/tablewriter"
	"github.com/prometheus/common/log"
	"github.com/spf13/cobra"
	"os"
	"runtime"
	"strconv"
)

// sshInfoCmd represents the sshInfo command
var sshInfoCmd = &cobra.Command{
	Use:   "sshinfo",
	Short: "view a ssh connection",
	Long:  `usage:felix sshinfo 1`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			color.Red("ID must be an integer")
			return
		}
		mac, err := models.MachineFind(uint(id))
		if err != nil {
			color.Red("can't fid machine by id of %d, [%s]", id, err)
			return
		}
		renderInfoTable(mac)
		err = flx.ShowHardwareInfo(mac)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(sshInfoCmd)
}

func renderInfoTable(m *models.Machine) {
	data := [][]string{
		{"ID", strconv.Itoa(int(m.ID))},
		{"Name", m.Name},
		{"Host", m.Host},
		{"IP", m.Ip},
		{"Port", strconv.Itoa(int(m.Port))},
		{"AuthType", m.Type},
		{"User", m.User},
		{"Password", "******"},
		{"Key", m.Key},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Column", "Value"})
	table.SetBorder(true)
	setInfoTableColor(table)
	table.AppendBulk(data) // Add Bulk Data
	table.Render()
}
func setInfoTableColor(table *tablewriter.Table) {
	if isatty.IsCygwinTerminal(os.Stdout.Fd()) || (runtime.GOOS != "windows") {
		table.SetHeaderColor(
			tablewriter.Colors{tablewriter.FgHiRedColor, tablewriter.Bold},
			tablewriter.Colors{tablewriter.FgHiGreenColor, tablewriter.Bold},
		)

		table.SetColumnColor(
			tablewriter.Colors{tablewriter.FgRedColor},
			tablewriter.Colors{tablewriter.FgCyanColor})
	}
}
