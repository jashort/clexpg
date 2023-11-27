package internal

type SearchCmd struct {
	Search string `arg:"" help:"Search string (case insensitive)"`
}

func (s *SearchCmd) Run(ctx *Context) error {
	var expenses = LoadFile(ctx.File)
	expenses = FilterItemSearch(expenses, s.Search)
	printAsTable(expenses)
	return nil
}
