package flx

import (
	"github.com/mattn/go-isatty"
	"github.com/olekukonko/tablewriter"
	"os"
)

func setListTableColor(table *tablewriter.Table) {
	if isatty.IsCygwinTerminal(os.Stdout.Fd()) {
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
}
