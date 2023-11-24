package internal

import (
	"fmt"
	"sort"
	"strings"
)

type SearchCmd struct {
	Search string `arg:"" help:"Search string (case insensitive)"`
}

func (s *SearchCmd) Run(ctx *Context) error {
	var expenses = LoadFile(ctx.File)
	sort.Sort(byDate(expenses))
	search := strings.ToLower(s.Search)
	for _, e := range expenses {
		if strings.Contains(e.Item, search) {
			fmt.Println(e)
		}
	}
	return nil
}
