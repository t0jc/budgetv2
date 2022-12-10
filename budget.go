package YileTDD

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