package internal

import (
	"fmt"
	"time"
)

type SummaryCmd struct {
}

func (s *SummaryCmd) Run(ctx *Context) error {
	expenses := LoadFile(ctx.File)
	now := time.Now()
	lastMonthNow := now.AddDate(0, -1, 0)
	lastYearNow := now.AddDate(-1, 0, 0)

	thisMonth := FilterTime(expenses, now.Year(), int(now.Month()))
	lastMonth := FilterTime(expenses, lastMonthNow.Year(), int(lastMonthNow.Month()))
	lastYear := FilterTime(expenses, lastYearNow.Year(), int(lastYearNow.Month()))
	thisMonthTotal := Total(thisMonth)
	lastMonthTotal := Total(lastMonth)
	lastYearTotal := Total(lastYear)

	thisMonthAverage := AverageSpentPerDay(thisMonth)
	lastMonthAverage := AverageSpentPerDay(lastMonth)
	lastYearAverage := AverageSpentPerDay(lastYear)
	println("Total Spent:")
	fmt.Printf("              This Month: %10s\n", FormatDec(thisMonthTotal))
	fmt.Printf("              Last Month: %10s\n", FormatDec(lastMonthTotal))
	fmt.Printf("    This Month Last Year: %10s\n", FormatDec(lastYearTotal))
	fmt.Printf("\nAverage spent per day:\n")
	fmt.Printf("              This Month: %10s\n", FormatDec(thisMonthAverage))
	fmt.Printf("              Last Month: %10s\n", FormatDec(lastMonthAverage))
	fmt.Printf("    This Month Last Year: %10s\n", FormatDec(lastYearAverage))
	println("")
	return nil
}
