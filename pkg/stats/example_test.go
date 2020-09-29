package stats

import (
	"fmt"
	"github.com/rukha89/bank/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			Amount: 10_000_00,
		},
		{
			Amount: 5_000_00,
		},
		{
			Amount: 6_000_00,
		},
	}
	fmt.Println(Avg(payments))
	// Output: 700000
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			Amount:   10_000_00,
			Category: `A`,
		},
		{
			Amount:   5_000_00,
			Category: `B`,
		},
		{
			Amount:   6_000_00,
			Category: `A`,
		},
	}
	fmt.Println(TotalInCategory(payments, `A`))
	// Output: 1600000
}
