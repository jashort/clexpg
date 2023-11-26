package internal

import (
	"fmt"
)

type TotalCmd struct {
	Year  int `arg:"" optional:"" help:"Year"`
	Month int `arg:"" optional:"" help:"Month"`
}

func (l *TotalCmd) Run(ctx *Context) error {
	var expenses = LoadFile(ctx.File)
	filtered := FilterTime(expenses, l.Year, l.Month)

	total := Total(filtered)

	println()
	if l.Year != 0 {
		if l.Month == 0 {
			fmt.Printf("         Total (%d): ", l.Year)
		} else {
			fmt.Printf("      Total (%d/%d): ", l.Month, l.Year)
		}
	} else {
		fmt.Printf("     Total (all time): ")
	}
	fmt.Printf("%12s", "$"+FormatDec(total))

	println()
	return nil
}
