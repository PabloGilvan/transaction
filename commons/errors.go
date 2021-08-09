package commons

import (
	"errors"
)

var (
	ErrAccountNotFound = errors.New("account not found")
	ErrAccountInactive = errors.New("account inactive")

	ErrOperationNotFound = errors.New("operation not found")
)
