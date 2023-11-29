package internal

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"os"
)

type DetailCmd struct {
	Year       int      `arg:"" optional:"" help:"Year"`
	Month      int      `arg:"" optional:"" help:"Month"`
	Categories []string `short:"c" optional:"" help:"Show only these categories (comma separated)"`
}

func (cmd *DetailCmd) Run(ctx *Context) error {
	var expenses = LoadFile(ctx.File)
	expenses = FilterTime(expenses, cmd.Year, cmd.Month)
	expenses = FilterCategories(expenses, cmd.Categories)
	totals := TotalByCategory(expenses)
	total := Total(expenses)
	println()

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, Name: "Category", AutoMerge: false, Align: text.AlignRight, WidthMin: 20},
		{Number: 2, Name: "Total", AutoMerge: false, Align: text.AlignRight, WidthMin: 12, AlignFooter: text.AlignRight},
	})
	if cmd.Year < 1 {
		t.SetTitle("Detail")
	} else {
		if cmd.Month == 0 {
			t.SetTitle(fmt.Sprintf("Detail (%d)", cmd.Year))
		} else {
			t.SetTitle(fmt.Sprintf("Detail (%d/%d)", cmd.Month, cmd.Year))
		}
	}
	t.SortBy([]table.SortBy{
		{Number: 1, Name: "Category", Mode: table.Asc},
	})
	for category, total := range totals {
		t.AppendRow(table.Row{category, FormatDec(total)})
	}

	t.AppendFooter(table.Row{"Total", FormatDec(total)})

	t.Render()
	return nil
}
