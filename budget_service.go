package YileTDD

import (
	"strconv"
	"time"
)

type BudgetService struct {
	br IBudgetRepo
}

func (bs BudgetService) Query(start, end time.Time) float64 {
	bList := bs.br.GetAll()

	dailyBudgetInMonthList := make(map[string]int, 0)
	for _, b := range bList {
		d, _ := time.Parse("200601", b.YearMonth())
		dailyBudgetInMonthList[b.YearMonth()] = b.Amount()/daysIn(d.Month(),d.Year())
	}

	validBList := make([]Budget, 0)
	startYearMonth := start.Year()*100+int(start.Month())
	endYearMonth := start.Year()*100+int(start.Month())
	for _, b := range bList {
		d, _ := strconv.ParseInt(b.YearMonth(), 10, 0)
		if int(d) >= startYearMonth && int(d) <= endYearMonth {
			 validBList = append(validBList, b)
		}
	}

	budgetAmount := 0
	for _, b := range validBList {
		ym := toYearMonthInt(b.yearMonth)
		ymString := strconv.Itoa(ym)
		// 完整月份
		if b.YearMonth() == ymString {
			datetime, _ := time.Parse("200601", b.YearMonth())
			monthBudgetDaily := dailyBudgetInMonthList[b.YearMonth()]
			budgetInterval := daysIn(datetime.Month(), datetime.Year()) * monthBudgetDaily
			budgetAmount += budgetInterval
		}
		// 起始月份
		if startYearMonth == ym {
			monthDays := daysIn(start.Month(), start.Year())
			days := monthDays - start.Day()

			monthBudgetDaily := dailyBudgetInMonthList[b.YearMonth()]
			budgetStart := days * monthBudgetDaily
			budgetAmount += budgetStart
		}
		// 結束月份
		if endYearMonth == ym {
			monthBudgetDaily := dailyBudgetInMonthList[b.YearMonth()]
			budgetEnd := end.Day() * monthBudgetDaily
			budgetAmount += budgetEnd
		}
	}
	return float64(budgetAmount)
}

func toYearMonthInt(yearmonth string) int {
	datetime, _ := time.Parse("200601", yearmonth)
	return datetime.Year()*100+int(datetime.Month())
}

func daysIn(m time.Month, year int) int {
	return time.Date(year, m+1, 0, 0, 0, 0, 0, time.UTC).Day()
}