package internal

import (
	"fmt"
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
	if l.Year == 0 {
		fmt.Printf("Totals by Month:")
	} else {
		fmt.Printf("Totals by Month for %d:\n", l.Year)
	}
	for _, k := range keys {
		fmt.Printf("     %10s: %10s\n", k, "$"+totals[k].StringFixed(2))
	}

	println()
	return nil
}
