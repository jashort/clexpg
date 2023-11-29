package internal

import (
	"sort"
)

type ListCmd struct {
	Year       int      `arg:"" optional:"" help:"Year"`
	Month      int      `arg:"" optional:"" help:"Month"`
	Categories []string `short:"c" optional:"" help:"Show only these categories (comma separated)"`
}

func (cmd *ListCmd) Run(ctx *Context) error {
	var expenses = LoadFile(ctx.File)
	expenses = FilterTime(expenses, cmd.Year, cmd.Month)
	expenses = FilterCategories(expenses, cmd.Categories)
	sort.Sort(byDate(expenses))
	printAsTable(expenses)
	return nil
}
