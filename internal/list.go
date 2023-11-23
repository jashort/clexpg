package internal

import (
	"fmt"
	"sort"
)

type ListCmd struct {
	Year  int `arg:"" optional:"" help:"Year"`
	Month int `arg:"" optional:"" help:"Month"`
}

func (l *ListCmd) Run(ctx *Context) error {
	var expenses = LoadFile(ctx.File)
	sort.Sort(byDate(expenses))
	for _, e := range expenses {
		if l.Year == 0 || l.Year == e.Date.Year() {
			if l.Month == 0 || l.Month == int(e.Date.Month()) {
				fmt.Println(e)
			}
		}
	}
	return nil
}
