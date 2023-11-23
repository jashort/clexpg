package internal

import (
	"fmt"
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
	for _, c := range categories {
		fmt.Printf("  %s\n", c)
	}
	return nil
}
