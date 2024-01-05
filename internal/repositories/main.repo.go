package repositories

import "gorm.io/gorm"

type Repositories struct {
	Transaction TransactionRepo
}

func New(db *gorm.DB) *Repositories {
	return &Repositories{}
}
