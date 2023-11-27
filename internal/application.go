package internal

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/shopspring/decimal"
	"golang.org/x/term"
	"log"
	"os"
)

type Context struct {
	File string
}

// FormatDec formats a decimal.Decimal value as a dollar amount, with thousands
// separated by ",". Eg: 123456.789 > "$123,456.79"
func FormatDec(number decimal.Decimal) string {
	output := number.StringFixedBank(2)
	startOffset := 3
	if number.LessThan(decimal.Zero) {
		startOffset++
	}

	for outputIndex := len(output) - 3; outputIndex > startOffset; {
		outputIndex -= 3
		output = output[:outputIndex] + "," + output[outputIndex:]
		if outputIndex <= 4 {
			break
		}
	}
	return "$" + output
}

func printAsTable(expenses []Expense) {
	maxItemLength := 80

	if term.IsTerminal(0) {
		// If we're outputting to a terminal, figure out what the maximum length of
		// the item column should be.
		width, _, err := term.GetSize(0)
		if err != nil {
			log.Fatal(err)
		}
		dateLength := 0
		categoryLength := 0
		costLength := 0
		for _, e := range expenses {
			dateLength = max(dateLength, len(e.Date.Format("1/2/2006")))
			categoryLength = max(categoryLength, len(e.Category))
			costLength = max(costLength, len(FormatDec(e.Cost)))
		}
		maxItemLength = width - (dateLength + categoryLength + costLength + 13)
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Date", "Category", "Item", "Cost"})
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, AutoMerge: false, Align: text.AlignRight},
		{Number: 2, AutoMerge: false},
		{Number: 3, AutoMerge: false, WidthMax: maxItemLength},
		{Number: 4, AutoMerge: false, Align: text.AlignRight},
	})
	for _, e := range expenses {
		t.AppendRow(table.Row{
			e.Date.Format("1/2/2006"),
			e.Category, text.WrapSoft(e.Item, maxItemLength),
			FormatDec(e.Cost)})
	}
	t.Render()
}
