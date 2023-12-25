package internal

import (
	"github.com/shopspring/decimal"
	"github.com/vjeantet/govaluate"
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

func (cmd *AddCmd) Run(ctx *Context) error {
	date := time.Now()
	if cmd.Date != "" {
		zone, err := time.LoadLocation("Local")
		if err != nil {
			log.Fatal(`Failed to load timezone location "Local"`)
		}
		x, err := time.ParseInLocation("01/02/2006", cmd.Date, zone)
		if err != nil {
			log.Fatalf("Error parsing %s: %s", cmd.Date, err)
		}
		date = x
	}

	expression, err := govaluate.NewEvaluableExpression(cmd.Amount)
	if err != nil {
		log.Fatalf("Unable to parse %s as expression", cmd.Amount)
	}
	result, err := expression.Evaluate(nil)
	if err != nil {
		log.Fatalf("Unable to parse %s as decimal", cmd.Amount)
	}
	amount := decimal.NewFromFloat(result.(float64))
	exp := Expense{
		Date:     date,
		Category: strings.TrimSpace(cases.Title(language.English).String(cmd.Category)),
		Item:     strings.TrimSpace(cmd.Description),
		Cost:     amount,
	}

	SaveExpense(exp, ctx.File)
	return nil
}
