package internal

type SearchCmd struct {
	Search     string   `arg:"" help:"Search string (case insensitive)"`
	Categories []string `short:"c" optional:"" help:"Show only these categories (comma separated)"`
}

func (cmd *SearchCmd) Run(ctx *Context) error {
	var expenses = LoadFile(ctx.File)
	expenses = FilterItemSearch(expenses, cmd.Search)
	expenses = FilterCategories(expenses, cmd.Categories)
	printAsTable(expenses)
	return nil
}
