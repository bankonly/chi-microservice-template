package repositories

import "gorm.io/gorm"

type TransactionRepo interface {
	Begin() *TransactionRepoOpts
	Rollback() *gorm.DB
	Commit() *gorm.DB
}

type TransactionRepoOpts struct {
	db *gorm.DB
}

func NewTransactionRepo(db *gorm.DB) TransactionRepo {
	return &TransactionRepoOpts{db: db}
}

func (opts *TransactionRepoOpts) Rollback() *gorm.DB {
	return opts.db.Rollback()
}
func (opts *TransactionRepoOpts) Commit() *gorm.DB {
	return opts.db.Commit()
}

func (opts *TransactionRepoOpts) Begin() *TransactionRepoOpts {
	return &TransactionRepoOpts{db: opts.db.Begin()}
}
