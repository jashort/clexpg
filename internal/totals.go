package internal

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/shopspring/decimal"
	"os"
	"sort"
)

type TotalsCmd struct {
	Year int `arg:"" optional:"" help:"Year"`
}

func (l *TotalsCmd) Run(ctx *Context) error {
	var expenses = LoadFile(ctx.File)

	totals := TotalByMonth(expenses, l.Year)

	keys := make([]string, 0, len(totals))
	for k := range totals {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	println()
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, Name: "Month", AutoMerge: false, Align: text.AlignRight, WidthMin: 10},
		{Number: 2, Name: "Total", AutoMerge: false, Align: text.AlignRight, WidthMin: 12, AlignFooter: text.AlignRight},
	})
	if l.Year < 1 {
		t.AppendHeader(table.Row{"Total by Year"})
	} else {
		t.AppendHeader(table.Row{fmt.Sprintf("Total by Month (%d)", l.Year)})
	}

	totalAmount := decimal.Zero
	for _, key := range keys {
		t.AppendRow(table.Row{key, FormatDec(totals[key])})
		totalAmount = totalAmount.Add(totals[key])
	}
	t.AppendFooter(table.Row{"Total", FormatDec(totalAmount)})
	t.Render()

	println()
	return nil
}
