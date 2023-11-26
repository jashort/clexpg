package internal

import "github.com/shopspring/decimal"

type Context struct {
	File string
}

// FormatDec formats a decimal.Decimal value as a dollar amount, with thousands
// separated by ",". Eg: 123456.789 > "$123,456.79"
func FormatDec(number decimal.Decimal) string {
	output := number.StringFixedBank(2)
	startOffset := 3
	if number.LessThan(decimal.Zero) {
		startOffset++
	}

	for outputIndex := len(output) - 3; outputIndex > startOffset; {
		outputIndex -= 3
		output = output[:outputIndex] + "," + output[outputIndex:]
		if outputIndex <= 4 {
			break
		}
	}
	return "$" + output
}
