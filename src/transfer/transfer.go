package transfer

type TransferResponse struct {
	BalanceOld float64
	BalanceNew float64
	Withdrawal float64
}

func TransferService(accountNumberFrom, accountNumberTo string, amountTransfer float64) TransferResponse {
	return TransferResponse{
		BalanceOld: 20000.00,
		BalanceNew: 19500.00,
		Withdrawal: 500.00,
	}
}
