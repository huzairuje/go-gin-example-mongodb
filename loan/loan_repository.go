package loan

type Repository interface {
	Store(req CreateLoanRequest) (*Loan, error)
	isSurpassLimit() bool
	List() ([]*Loan, error)
	IsKtpNumberExist(ktpNumber string) bool
}
