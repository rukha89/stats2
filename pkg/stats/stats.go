package stats

import "github.com/rukha89/bank/pkg/types"


func Avg(payments []types.Payment) types.Money{
	avg := types.Money(0)
	sum := types.Money(0)
	cnt := types.Money(0)
	for _, payment := range payments {
		sum += payment.Amount
		cnt += types.Money(1)
	}
	avg = sum / cnt
	return avg
	
}

// TotalInCategory находит сумму покупок в определенной категории.
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := types.Money(0)
	for _, payment := range payments {
		if payment.Category == category{
			sum += payment.Amount
		}
	}
	return sum
}