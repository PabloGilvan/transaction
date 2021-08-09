package transaction

type TransactionPersist struct {
	AccountId       string  `json:"account_id"`
	OperationTypeId int     `json:"operation_type_id"`
	Amount          float64 `json:"amount"`
}

type TransactionResponse struct {
	TransactionIdentifier string `json:"transaction_identifier"`
}
