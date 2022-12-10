package YileTDD

type IBudget interface {
	YearMonth() string
	Amount() int
}

type Budget struct {
	yearMonth string
	amount int
}

func (b Budget) YearMonth() string {
	return b.yearMonth
}

func (b Budget) Amount() int {
	return b.amount
}