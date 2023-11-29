package internal

import (
	"fmt"
)

type TotalCmd struct {
	Year       int      `arg:"" optional:"" help:"Year"`
	Month      int      `arg:"" optional:"" help:"Month"`
	Categories []string `short:"c" optional:"" help:"Show only these categories (comma separated)"`
}

func (cmd *TotalCmd) Run(ctx *Context) error {
	var expenses = LoadFile(ctx.File)
	expenses = FilterTime(expenses, cmd.Year, cmd.Month)
	expenses = FilterCategories(expenses, cmd.Categories)
	total := Total(expenses)

	println()
	if cmd.Year != 0 {
		if cmd.Month == 0 {
			fmt.Printf("         Total (%d): ", cmd.Year)
		} else {
			fmt.Printf("      Total (%d/%d): ", cmd.Month, cmd.Year)
		}
	} else {
		fmt.Printf("     Total (all time): ")
	}
	fmt.Printf("%12s", FormatDec(total))

	println()
	return nil
}
