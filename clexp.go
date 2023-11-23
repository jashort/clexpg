package main

import (
	"bufio"
	"fmt"
	"github.com/alecthomas/kong"
	"github.com/shopspring/decimal"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"log"
	"os"
	"strings"
	"time"
)

var CLI struct {
	File string `default:"expenses.csv" name:"file" help:"Data file"`

	Rm struct {
		Force     bool `help:"Force removal."`
		Recursive bool `help:"Recursively remove files."`

		Paths []string `arg:"" name:"path" help:"Paths to remove." type:"path"`
	} `cmd:"" help:"Remove files."`

	List struct {
		Year  int `arg:"" help:"Year" optional:""`
		Month int `arg:"" help:"Month" optional:""`
	} `cmd:"" help:"List expenses"`
	Test struct {
	} `cmd:"" help:"Test"`
}

// Expense A single expense
type Expense struct {
	Date     time.Time
	Category string
	Item     string
	Cost     decimal.Decimal
}

func (e Expense) String() string {
	return fmt.Sprintf("%s\t%s\t%s\t%s", e.Date.Format("1/2/2006"), e.Category, e.Item, e.Cost)
}

type byDate []Expense

func (a byDate) Len() int           { return len(a) }
func (a byDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byDate) Less(i, j int) bool { return a[i].Date.Before(a[j].Date) }

// Parses a tab-separated string in to an Expense struct
// Example: "5/21/2023\t fun\tMy stuff\t$10.00"
func parseExpense(data string) Expense {
	s := strings.Split(strings.TrimSpace(data), "\t")
	timestamp, err := time.Parse("1/2/2006", s[0])
	if err != nil {
		log.Fatal("Error parsing ", data, err)
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
	ctx := kong.Parse(&CLI)
	switch ctx.Command() {
	case "rm <path>":
		println("rm")
	case "list":
		println("list")
	case "list <year> <month>":
		println("list")
	case "test":
		println(formatExpense(parseExpense("5/21/2023\tfun\tMy stuff\t$10.00")))
	default:
		panic(ctx.Command())
	}
}

func loadFile(s string) []Expense {
	var expenses []Expense
	file, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Scan() // Skip first line
	for scanner.Scan() {
		item := parseExpense(scanner.Text())
		expenses = append(expenses, item)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return expenses
}
