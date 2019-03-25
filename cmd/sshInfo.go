package cmd

import (
	"github.com/dejavuzhou/felix/models"
	"github.com/fatih/color"
	"github.com/mattn/go-isatty"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

// sshInfoCmd represents the sshInfo command
var sshInfoCmd = &cobra.Command{
	Use:   "sshInfo",
	Short: "",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			color.Red("ID参数必须为正整数")
			return
		}
		mac, err := models.MachineFind(uint(id))
		if err != nil {
			color.Red("can't fid machine by id of %d, [%s]", id, err)
			return
		}
		renderInfoTable(mac)
	},
}

func init() {
	rootCmd.AddCommand(sshInfoCmd)
}

func renderInfoTable(m *models.Machine) {
	data := [][]string{
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
	if isatty.IsCygwinTerminal(os.Stdout.Fd()) {
		table.SetHeaderColor(
			tablewriter.Colors{tablewriter.FgHiRedColor, tablewriter.Bold},
			tablewriter.Colors{tablewriter.FgHiGreenColor, tablewriter.Bold},
		)

		table.SetColumnColor(
			tablewriter.Colors{tablewriter.FgRedColor},
			tablewriter.Colors{tablewriter.FgCyanColor})
	}
}
