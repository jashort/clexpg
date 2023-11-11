package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
	"time"
)

type Expense struct {
	Date     time.Time
	Category string
	Item     string
	Cost     decimal.Decimal
}

func parseExpense(data string) Expense {
	s := strings.Split(strings.TrimSpace(data), "\t")
	timestamp, err := time.Parse("1/2/2006", s[0])
	if err != nil {
		panic(err)
	}
	cost := decimal.RequireFromString(strings.Replace(s[3], "$", "", 1))
	p := Expense{
		Date:     timestamp,
		Category: cases.Title(language.English).String(s[1]),
		Item:     s[2],
		Cost:     cost,
	}
	return p
}

func formatExpense(expense Expense) string {
	return fmt.Sprintf("%s\t%s\t%s\t$%s", expense.Date.Format("1/2/2006"), expense.Category, expense.Item, expense.Cost.StringFixed(2))
}

func main() {
	println(formatExpense(parseExpense("5/21/2023\tfun\tMy stuff\t$10.00")))
}
