package internal

import (
	"sort"
)

type ListCmd struct {
	Year  int `arg:"" optional:"" help:"Year"`
	Month int `arg:"" optional:"" help:"Month"`
}

func (l *ListCmd) Run(ctx *Context) error {
	var expenses = LoadFile(ctx.File)
	expenses = FilterTime(expenses, l.Year, l.Month)
	sort.Sort(byDate(expenses))
	printAsTable(expenses)
	return nil
}
