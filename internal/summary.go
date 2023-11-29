package internal

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"os"
	"time"
)

type SummaryCmd struct {
	Categories []string `short:"c" optional:"" help:"Show only these categories (comma separated)"`
}

func (cmd *SummaryCmd) Run(ctx *Context) error {
	expenses := LoadFile(ctx.File)
	expenses = FilterCategories(expenses, cmd.Categories)
	now := time.Now()
	lastMonthNow := now.AddDate(0, -1, 0)
	lastYearNow := now.AddDate(-1, 0, 0)

	thisMonth := FilterTime(expenses, now.Year(), int(now.Month()))
	lastMonth := FilterTime(expenses, lastMonthNow.Year(), int(lastMonthNow.Month()))
	lastYear := FilterTime(expenses, lastYearNow.Year(), int(lastYearNow.Month()))
	thisMonthTotal := Total(thisMonth)
	lastMonthTotal := Total(lastMonth)
	lastYearTotal := Total(lastYear)

	thisMonthAverage := AverageSpentPerDay(thisMonth)
	lastMonthAverage := AverageSpentPerDay(lastMonth)
	lastYearAverage := AverageSpentPerDay(lastYear)

	println()
	t := table.NewWriter()
	t.Style().Options.DrawBorder = false
	t.Style().Options.SeparateHeader = false
	t.Style().Options.SeparateColumns = false
	t.SetOutputMirror(os.Stdout)
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, AutoMerge: false, Align: text.AlignRight, WidthMin: 30},
		{Number: 2, AutoMerge: false, Align: text.AlignRight, WidthMin: 12},
	})
	t.AppendHeader(table.Row{"Total Spent"})
	t.AppendRow(table.Row{"This Month", FormatDec(thisMonthTotal)})
	t.AppendRow(table.Row{"Last Month", FormatDec(lastMonthTotal)})
	t.AppendRow(table.Row{"This Month Last Year", FormatDec(lastYearTotal)})
	t.Render()

	println()
	t = table.NewWriter()
	t.Style().Options.DrawBorder = false
	t.Style().Options.SeparateHeader = false
	t.Style().Options.SeparateColumns = false
	t.SetOutputMirror(os.Stdout)
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, AutoMerge: false, Align: text.AlignRight, WidthMin: 30},
		{Number: 2, AutoMerge: false, Align: text.AlignRight, WidthMin: 12},
	})
	t.AppendHeader(table.Row{"Average Spent per Day"})
	t.AppendRow(table.Row{"This Month", FormatDec(thisMonthAverage)})
	t.AppendRow(table.Row{"Last Month", FormatDec(lastMonthAverage)})
	t.AppendRow(table.Row{"This Month Last Year", FormatDec(lastYearAverage)})
	t.Render()
	return nil
}
