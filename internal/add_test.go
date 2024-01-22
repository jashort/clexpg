package internal

import (
	"testing"
)

func Test_parseExpense(t *testing.T) {
	type args struct {
		cmd *AddCmd
	}

	var tests = []struct {
		name string
		args args
		want string
	}{
		{name: "Happy Path",
			args: args{cmd: &AddCmd{
				Amount:      "1",
				Category:    "food",
				Description: "hamburger",
				Date:        "3/19/2024",
			}},
			want: "3/19/2024\tFood\thamburger\t$1.00"},
		{name: "Zero padded date",
			args: args{cmd: &AddCmd{
				Amount:      "1",
				Category:    "food",
				Description: "hamburger",
				Date:        "03/09/2024",
			}},
			want: "3/9/2024\tFood\thamburger\t$1.00"},
		{name: "Addition in amount",
			args: args{cmd: &AddCmd{
				Amount:      "10+15",
				Category:    "food",
				Description: "hamburger",
				Date:        "03/09/2024",
			}},
			want: "3/9/2024\tFood\thamburger\t$25.00"},
		{name: "Subtraction amount",
			args: args{cmd: &AddCmd{
				Amount:      "335-35",
				Category:    "food",
				Description: "hamburger",
				Date:        "03/09/2024",
			}},
			want: "3/9/2024\tFood\thamburger\t$300.00"},
		{name: "Multiple addition",
			args: args{cmd: &AddCmd{
				Amount:      "9.99+3.01+20.55",
				Category:    "groceries",
				Description: "groceries",
				Date:        "03/09/2024",
			}},
			want: "3/9/2024\tGroceries\tgroceries\t$33.55"},
		{name: "Multiple subtraction",
			args: args{cmd: &AddCmd{
				Amount:      "(100-9.99)-10.33",
				Category:    "groceries",
				Description: "groceries",
				Date:        "03/09/2024",
			}},
			want: "3/9/2024\tGroceries\tgroceries\t$79.68"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseExpense(tt.args.cmd)
			if got.String() != tt.want {
				t.Errorf("parseExpense() = %v, want %v", got, tt.want)
			}
		})
	}
}
