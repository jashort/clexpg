package internal

import (
	"github.com/shopspring/decimal"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"log"
	"strings"
	"time"
)

type AddCmd struct {
	Amount      string `arg:"" help:"Amount"`
	Category    string `arg:"" help:"Category"`
	Description string `arg:"" help:"Description"`
	Date        string `arg:"" optional:"" help:"Date in MM/DD/YYYY format (default: today)"`
}

func (a *AddCmd) Run(ctx *Context) error {
	date := time.Now()
	if a.Date != "" {
		zone, err := time.LoadLocation("Local")
		if err != nil {
			log.Fatal(`Failed to load timezone location "Local"`)
		}
		x, err := time.ParseInLocation("01/02/2006", a.Date, zone)
		if err != nil {
			log.Fatalf("Error parsing %s: %s", a.Date, err)
		}
		date = x
	}

	amount, err := decimal.NewFromString(a.Amount)
	if err != nil {
		log.Fatalf("Unable to parse %s as decimal", a.Amount)
	}
	exp := Expense{
		Date:     date,
		Category: strings.TrimSpace(cases.Title(language.English).String(a.Category)),
		Item:     strings.TrimSpace(a.Description),
		Cost:     amount,
	}

	SaveExpense(exp, ctx.File)
	return nil
}
