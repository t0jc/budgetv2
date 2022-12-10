package YileTDD

import "time"

type BudgetService struct {
	br IBudgetRepo
}

func (bs BudgetService) Query(start, end time.Time) float64 {

	bList := bs.br.GetAll()

	dailyBudgetInMonthList := make(map[string]int, 0)
	for _, b := range bList {
		dailyBudgetInMonthList[b.YearMonth()] = b.Amount()/30
	}


	//validBList := make([]Budget, 0)
	//for _, b := range bList {
	//	t, _ := time.Parse("200601", b.YearMonth())
	//	t.mon
	//	b.YearMonth()
	//}

}