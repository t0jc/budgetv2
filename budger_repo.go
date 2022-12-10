package YileTDD

type IBudgetRepo interface {
	GetAll() []Budget
}

//type BudgetRepo struct {}
//
//func (br BudgetRepo) GetAll() []Budget {
//	return []Budget{}
//}