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

func checkTransferPerDay(amount float64) bool {
	if amount > 100000.00 {
		return true
	}
	return false
}

type Account struct {
	Name           string  `json:"name"`
	Id             string  `json:"id"`
	Balance        float64 `json:"balance"`
	TransferPerDay float64 `json:"transferperday"`
}

func getAccount(accountNumber string) Account {
	if accountNumber == "981751424" {
		return Account{Name: "Panumars Seanto", Id: "981751424", Balance: 20000.00, TransferPerDay: 0.00}
	}
	if accountNumber == "981751425" {
		return Account{Name: "Piyanuch Ekpiyakool", Id: "981751425", Balance: 100000.00, TransferPerDay: 0.00}
	}
	return Account{}
}
