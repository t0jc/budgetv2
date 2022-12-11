package YileTDD

import (
	"strconv"
	"time"
)

type BudgetService struct {
	repo IBudgetRepo
}

func (bs BudgetService) Query(start, end time.Time) float64 {
	if end.Before(start) {
		return float64(0)
	}

	bList := bs.repo.GetAll()

	dailyBudgetInMonthList := make(map[string]int)
	validBList := make([]Budget, 0)
	startYearMonth := start.Year()*100 + int(start.Month())
	endYearMonth := end.Year()*100 + int(end.Month())

	for _, b := range bList {
		budgetTime, _ := time.Parse("200601", b.YearMonth())
		dailyBudgetInMonthList[b.YearMonth()] = b.Amount() / MonthOfDays(budgetTime.Month(), budgetTime.Year())

		d, _ := strconv.ParseInt(b.YearMonth(), 10, 0)
		if int(d) >= startYearMonth && int(d) <= endYearMonth {
			validBList = append(validBList, b)
		}
	}

	budgetAmount := 0
	if startYearMonth == endYearMonth {
		startDay := start.Day()
		endDay := end.Day()
		days := endDay - startDay + 1
		return float64(dailyBudgetInMonthList[start.Format("200601")] * days)
	}

	for _, b := range validBList {
		ym := toYearMonthInt(b.yearMonth)
		ymString := strconv.Itoa(ym)

		// 起始月份
		if startYearMonth == ym {
			monthDays := MonthOfDays(start.Month(), start.Year())
			days := monthDays - start.Day() + 1

			monthBudgetDaily := dailyBudgetInMonthList[b.YearMonth()]
			budgetStart := days * monthBudgetDaily
			budgetAmount += budgetStart
			continue
		}
		// 結束月份
		if endYearMonth == ym {
			monthBudgetDaily := dailyBudgetInMonthList[b.YearMonth()]
			budgetEnd := end.Day() * monthBudgetDaily
			budgetAmount += budgetEnd
			continue
		}

		// 完整月份
		if b.YearMonth() == ymString {
			datetime, _ := time.Parse("200601", b.YearMonth())
			monthBudgetDaily := dailyBudgetInMonthList[b.YearMonth()]
			budgetInterval := MonthOfDays(datetime.Month(), datetime.Year()) * monthBudgetDaily
			budgetAmount += budgetInterval
		}
	}
	return float64(budgetAmount)
}

func toYearMonthInt(yearmonth string) int {
	datetime, _ := time.Parse("200601", yearmonth)
	return datetime.Year()*100 + int(datetime.Month())
}

func MonthOfDays(m time.Month, year int) int {
	return time.Date(year, m+1, 0, 0, 0, 0, 0, time.UTC).Day()
}
