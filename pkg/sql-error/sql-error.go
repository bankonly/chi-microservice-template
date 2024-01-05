package sqlError

import (
	"github.com/bankonly/go-pkg/v1/sanitizers"
	"gorm.io/gorm"
)

var (
	ErrDuplicatedRecord = "23505"
)

func IsDuplicatedError(err error) bool {
	errCode := sanitizers.GetSQLErrorCode(err)
	return errCode == ErrDuplicatedRecord
}

func CatchWithDuplicatedErr(err error, out error) error {
	if err != nil {
		if IsDuplicatedError(err) {
			return out
		}
		return err
	}
	return nil
}

func CatchWithNotFoundErr(err error, out error) error {
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return out
		}
		return err
	}
	return nil
}

func CatchUpdateErr(err error, duplicateErrOut error, notFoundErrOut error) error {
	if err != nil {
		if IsDuplicatedError(err) {
			return duplicateErrOut
		}

		if err == gorm.ErrRecordNotFound {
			return notFoundErrOut
		}
		return err
	}
	return nil
}
