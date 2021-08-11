package helpers

import (
	"errors"
	"github.com/PabloGilvan/transaction/commons"
	"net/http"
)

type CustomResponseError struct {
	ErrorMessage string
}

var (
	ErrInvalidRequest = errors.New("invalid request")
)

func ProcessIfBusinessError(err error) (*CustomResponseError, int) {
	if err == nil {
		return nil, 0
	}

	var errorMessage = &CustomResponseError{ErrorMessage: err.Error()}

	if err == commons.ErrAccountNotFound {
		return errorMessage, http.StatusNotFound
	}

	if err == commons.ErrAccountInactive {
		return errorMessage, http.StatusBadRequest
	}

	if err == commons.ErrOperationNotFound {
		return errorMessage, http.StatusBadRequest
	}

	if err == commons.ErrLimitExceeded {
		return errorMessage, http.StatusBadRequest
	}

	return errorMessage, http.StatusInternalServerError
}
