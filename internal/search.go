package internal

type SearchCmd struct {
	Search     string   `arg:"" help:"Search string (case insensitive)"`
	Categories []string `short:"c" optional:"" help:"Show only these categories (comma separated)"`
}

func (s *SearchCmd) Run(ctx *Context) error {
	var expenses = LoadFile(ctx.File)
	expenses = FilterItemSearch(expenses, s.Search)
	expenses = FilterCategories(expenses, s.Categories)
	printAsTable(expenses)
	return nil
}
