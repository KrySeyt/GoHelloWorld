package interfaces

type TransactionManager interface {
	Begin()
	Commit()
}
