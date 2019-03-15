// +build !windows

package flx

import (
	"github.com/olekukonko/tablewriter"
)

func setListTableColor(table *tablewriter.Table) {
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.FgHiRedColor, tablewriter.Bold},
		tablewriter.Colors{tablewriter.FgHiGreenColor, tablewriter.Bold},
		tablewriter.Colors{tablewriter.FgHiGreenColor, tablewriter.Bold},
		tablewriter.Colors{tablewriter.FgHiGreenColor, tablewriter.Bold},
		tablewriter.Colors{tablewriter.FgHiGreenColor, tablewriter.Bold},
		tablewriter.Colors{tablewriter.FgHiGreenColor, tablewriter.Bold},
		tablewriter.Colors{tablewriter.FgHiGreenColor, tablewriter.Bold})
	table.SetColumnColor(
		tablewriter.Colors{tablewriter.FgRedColor},
		tablewriter.Colors{tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.FgCyanColor})
}
