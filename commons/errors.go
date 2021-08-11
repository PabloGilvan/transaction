package commons

import (
	"errors"
)

var (
	ErrAccountNotFound = errors.New("account not found")
	ErrAccountInactive = errors.New("account inactive")

	ErrOperationNotFound = errors.New("operation not found")

	ErrLimitExceeded = errors.New("the account doesnt have enough limit for this transaction")
)
