package internal

import (
	"fmt"
)

type DetailCmd struct {
	Year  int `arg:"" optional:"" help:"Year"`
	Month int `arg:"" optional:"" help:"Month"`
}

func (l *DetailCmd) Run(ctx *Context) error {
	var expenses = LoadFile(ctx.File)
	filtered := FilterTime(expenses, l.Year, l.Month)
	totals := TotalByCategory(filtered)
	total := Total(filtered)
	println()
	for category, total := range totals {
		fmt.Printf("     %20s: %10s\n", category, "$"+total.StringFixed(2))
	}
	println()
	fmt.Printf("     %20s: %10s\n", "Total", "$"+total.StringFixed(2))
	println()
	return nil
}
