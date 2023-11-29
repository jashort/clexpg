package main

import (
	"clexpg/internal"
	"github.com/alecthomas/kong"
)

var cli struct {
	File string `default:"expenses.csv" name:"file" short:"f" help:"Data file"`

	List       internal.ListCmd       `cmd:"" help:"List expenses"`
	Summary    internal.SummaryCmd    `cmd:"" help:"Summarize this month vs last month"`
	Categories internal.CategoriesCmd `cmd:"" help:"List currently used categories"`
	Detail     internal.DetailCmd     `cmd:"" help:"Show totals by category for the given time period"`
	Add        internal.AddCmd        `cmd:"" help:"Add expense"`
	Search     internal.SearchCmd     `cmd:"" help:"Search expenses"`
	Total      internal.TotalCmd      `cmd:"" help:"Total"`
	Totals     internal.TotalsCmd     `cmd:"" help:"Total expenses by year/month"`
}

func main() {
	ctx := kong.Parse(&cli)
	err := ctx.Run(&internal.Context{File: cli.File})
	ctx.FatalIfErrorf(err)
}
