package internal

import (
	"sort"
)

type ListCmd struct {
	Year       int      `arg:"" optional:"" help:"Year"`
	Month      int      `arg:"" optional:"" help:"Month"`
	Categories []string `short:"c" optional:"" help:"Show only these categories (comma separated)"`
}

func (l *ListCmd) Run(ctx *Context) error {
	var expenses = LoadFile(ctx.File)
	expenses = FilterTime(expenses, l.Year, l.Month)
	expenses = FilterCategories(expenses, l.Categories)
	sort.Sort(byDate(expenses))
	printAsTable(expenses)
	return nil
}
