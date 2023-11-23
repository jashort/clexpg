package internal

import (
	"bufio"
	"fmt"
	"github.com/shopspring/decimal"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"log"
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

func (e Expense) String() string {
	return fmt.Sprintf("%s\t%s\t%s\t$%s", e.Date.Format("1/2/2006"), e.Category, e.Item, e.Cost.StringFixed(2))
}

type byDate []Expense

func (a byDate) Len() int           { return len(a) }
func (a byDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byDate) Less(i, j int) bool { return a[i].Date.Before(a[j].Date) }

func LoadFile(s string) []Expense {
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
		item := ParseExpense(scanner.Text())
		expenses = append(expenses, item)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return expenses
}

// Parses a tab-separated string in to an Expense struct
// Example: "5/21/2023\t fun\tMy stuff\t$10.00"
func ParseExpense(data string) Expense {
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
