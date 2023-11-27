package internal

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
	"slices"
	"sort"
)

type CategoriesCmd struct {
}

func (c CategoriesCmd) Run(ctx *Context) error {
	expenses := LoadFile(ctx.File)
	var categories []string
	for _, e := range expenses {
		if !slices.Contains(categories, e.Category) {
			categories = append(categories, e.Category)
		}
	}
	sort.Sort(sort.StringSlice(categories))

	println()
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, AutoMerge: false, WidthMin: 20},
	})
	t.AppendHeader(table.Row{"Categories"})
	for _, c := range categories {
		t.AppendRow(table.Row{c})
	}
	t.Render()

	return nil
}
