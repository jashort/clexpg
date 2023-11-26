package internal

import (
	"github.com/shopspring/decimal"
	"testing"
)

func Test_FormatDec(t *testing.T) {
	result := FormatDec(decimal.NewFromInt32(123456))
	if result != "$123,456.00" {
		t.Fatalf("FormatDec(123456) got %s", result)
	}
}

func Test_FormatDecCents(t *testing.T) {
	dec, _ := decimal.NewFromString("123456.789")
	result := FormatDec(dec)
	if result != "$123,456.79" {
		t.Fatalf("FormatDec(123456) got %s", result)
	}
}
