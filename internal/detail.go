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

func (l *DetailCmd) Run(ctx *Context) error {
	var expenses = LoadFile(ctx.File)
	filtered := FilterTime(expenses, l.Year, l.Month)
	filtered = FilterCategories(filtered, l.Categories)
	totals := TotalByCategory(filtered)
	total := Total(filtered)
	println()

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, Name: "Category", AutoMerge: false, Align: text.AlignRight, WidthMin: 20},
		{Number: 2, Name: "Total", AutoMerge: false, Align: text.AlignRight, WidthMin: 12, AlignFooter: text.AlignRight},
	})
	if l.Year < 1 {
		t.SetTitle("Detail", table.TitleOptions{})
	} else {
		if l.Month == 0 {
			t.SetTitle(fmt.Sprintf("Detail (%d)", l.Year))
		} else {
			t.SetTitle(fmt.Sprintf("Detail (%d/%d)", l.Month, l.Year))
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
