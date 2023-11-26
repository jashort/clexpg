package main

import (
	"clexpg/internal"
	"fmt"
	"github.com/alecthomas/kong"
)

type TestCmd struct {
}

func (l *TestCmd) Run(ctx *internal.Context) error {
	fmt.Println("test")
	println("File: ", ctx.File)
	println(internal.ParseExpense("5/21/2023\tfun\tMy stuff\t$10.00").String())
	return nil
}

var cli struct {
	File  string `default:"expenses.csv" name:"file" help:"Data file"`
	Debug bool   `help:"Enable debug mode."`

	List       internal.ListCmd       `cmd:"" help:"List expenses"`
	Summary    internal.SummaryCmd    `cmd:"" help:"Summarize this month vs last month"`
	Categories internal.CategoriesCmd `cmd:"" help:"List currently used categories"`
	Detail     internal.DetailCmd     `cmd:"" help:"Show totals by category for the given time period"`
	Add        internal.AddCmd        `cmd:"" help:"Add expense"`
	Search     internal.SearchCmd     `cmd:"" help:"Search expenses"`
	Total      internal.TotalCmd      `cmd:"" help:"Total"`
	Totals     internal.TotalsCmd     `cmd:"" help:"Total expenses by month"`
	Test       TestCmd                `cmd:"" help:"Test"`
}

func main() {
	ctx := kong.Parse(&cli)
	err := ctx.Run(&internal.Context{File: cli.File})
	ctx.FatalIfErrorf(err)
}
