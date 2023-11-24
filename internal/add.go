package internal

import (
	"fmt"
)

type AddCmd struct {
	Amount      string `arg:"" help:"Amount"`
	Category    string `arg:"" help:"Category"`
	Description string `arg:"" help:"Description"`
	Date        string `arg:"" optional:"" help:"Date in MM/DD/YYYY format (default: today)"`
}

func (a *AddCmd) Run(ctx *Context) error {

	fmt.Println(a.Amount, a.Category, a.Description, a.Date)
	fmt.Println(ctx.File)
	// Parse fields
	// Date now if empty
	// If file doesn't exist, add header
	// append to file
	return nil
}
