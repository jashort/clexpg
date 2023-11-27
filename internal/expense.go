package internal

import (
	"bufio"
	"errors"
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

// Expense as a printable string.
func (e Expense) String() string {
	return fmt.Sprintf("%s\t%s\t%s\t%s", e.Date.Format("1/2/2006"), e.Category, e.Item, FormatDec(e.Cost))
}

// Serialize Expense as a serialized string. Notably, does not include commas in the cost
func (e Expense) Serialize() string {
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

// SaveExpense writes the given Expense to the TSV file filename, creating it with
// a header if it doesn't exist
func SaveExpense(expense Expense, filename string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			file, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatalf("Could not create file %s", filename)
			}
			_, err := file.WriteString("Date\tCategory\tItem\tCost\n")
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}
	}
	_, err = file.WriteString(expense.Serialize() + "\n")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)
}

func Total(expenses []Expense) decimal.Decimal {
	total := decimal.Zero

	for _, e := range expenses {
		total = total.Add(e.Cost)
	}
	return total
}

func TotalByCategory(expenses []Expense) map[string]decimal.Decimal {
	output := make(map[string]decimal.Decimal)

	for _, e := range expenses {
		output[e.Category] = e.Cost.Add(output[e.Category])
	}
	return output
}

func TotalByMonth(expenses []Expense, year int) map[string]decimal.Decimal {
	output := make(map[string]decimal.Decimal)

	for _, e := range expenses {
		if year == 0 || e.Date.Year() == year {
			month := e.Date.Format("2006-01")
			output[month] = e.Cost.Add(output[month])
		}
	}

	return output
}

func FilterTime(expenses []Expense, year int, month int) []Expense {
	var output []Expense
	for _, e := range expenses {
		if year == 0 || year == e.Date.Year() {
			if month == 0 || month == int(e.Date.Month()) {
				output = append(output, e)
			}
		}
	}
	return output
}

func FilterItemSearch(expenses []Expense, search string) []Expense {
	var output []Expense
	searchString := strings.ToLower(search)
	for _, e := range expenses {
		if strings.Contains(strings.ToLower(e.Item), searchString) {
			output = append(output, e)
		}
	}
	return output

}

func FilterCategory(expenses []Expense, category string) []Expense {
	var output []Expense
	lCase := strings.ToLower(category)
	for _, e := range expenses {
		if category == "" || strings.ToLower(category) == lCase {
			output = append(output, e)
		}
	}
	return output
}

// ParseExpense parses a tab-separated string in to an Expense struct
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

func calendarDays(t2, t1 time.Time) int {
	y, m, d := t2.Date()
	u2 := time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	y, m, d = t1.In(t2.Location()).Date()
	u1 := time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	days := u2.Sub(u1) / (24 * time.Hour)
	return int(days) + 1
}

func AverageSpentPerDay(expenses []Expense) decimal.Decimal {
	if len(expenses) == 0 {
		return decimal.Zero
	}

	var firstDay = expenses[0].Date
	var lastDay = expenses[0].Date

	for _, e := range expenses {
		if firstDay.After(e.Date) {
			firstDay = e.Date
		}
		if lastDay.Before(e.Date) {
			lastDay = e.Date
		}
	}

	firstDay = time.Date(firstDay.Year(), firstDay.Month(), 1, 0, 0, 0, 0, time.Local)
	lastDay = time.Date(lastDay.Year(), lastDay.Month(), 1, 0, 0, 0, 0, time.Local).AddDate(0, 1, 0).Add(time.Nanosecond * -1)
	return Total(expenses).DivRound(decimal.NewFromInt32(int32(calendarDays(lastDay, firstDay))), 2)
}
