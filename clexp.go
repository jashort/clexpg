package main

import (
	"flag"
	"fmt"
	"github.com/shopspring/decimal"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"strings"
	"time"
)

// Expense A single expense
type Expense struct {
	Date     time.Time
	Category string
	Item     string
	Cost     decimal.Decimal
}

// Parses a tab-separated string in to an Expense struct
// Example: "5/21/2023\t fun\tMy stuff\t$10.00"
func parseExpense(data string) Expense {
	s := strings.Split(strings.TrimSpace(data), "\t")
	timestamp, err := time.Parse("1/2/2006", s[0])
	if err != nil {
		panic(err)
	}
	cost := decimal.RequireFromString(strings.Replace(s[3], "$", "", 1))
	p := Expense{
		Date:     timestamp,
		Category: strings.TrimSpace(cases.Title(language.English).String(s[1])),
		Item:     strings.TrimSpace(s[2]),
		Cost:     cost,
	}
	return p
}

// converts an Expense in to a tab-delimited string
func formatExpense(expense Expense) string {
	return fmt.Sprintf("%s\t%s\t%s\t$%s", expense.Date.Format("1/2/2006"), expense.Category, expense.Item, expense.Cost.StringFixed(2))
}

func main() {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addFile := addCmd.String("f", "expenses.csv", "Path to data file")

	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	listFile := listCmd.String("f", "expenses.csv", "Path to data file")

	command := "help"
	if len(os.Args) > 1 {
		command = os.Args[1]
	}

	switch command {
	case "add":
		err := addCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("add")
		fmt.Printf("  file: %s\n", *addFile)
	case "list":
		fmt.Println("list")
		err := listCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("list")
		fmt.Printf("  file: %s\n", *listFile)
	case "total":
		fmt.Println("total")
	case "summary":
		fmt.Println("summary")
	case "totals":
		fmt.Println("totals")
	case "categories":
		fmt.Println("categories")
	case "detail":
		fmt.Println("detail")
	case "search":
		fmt.Println("search")
	case "test":
		println(formatExpense(parseExpense("5/21/2023\tfun\tMy stuff\t$10.00")))
	default:
		fmt.Println(`Expected subcommand:
        add, list, summary, total, totals, categories, search, detail, help`)
		os.Exit(1)
	}
}
