// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package transaction

import mock "github.com/stretchr/testify/mock"

// MockTransactionRepository is an autogenerated mock type for the TransactionRepository type
type MockTransactionRepository struct {
	mock.Mock
}

// SaveTransaction provides a mock function with given fields: transactionModel
func (_m *MockTransactionRepository) SaveTransaction(transactionModel Transaction) (*Transaction, error) {
	ret := _m.Called(transactionModel)

	var r0 *Transaction
	if rf, ok := ret.Get(0).(func(Transaction) *Transaction); ok {
		r0 = rf(transactionModel)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(Transaction) error); ok {
		r1 = rf(transactionModel)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}