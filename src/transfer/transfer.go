package transfer

const LIMITPERTRANSACTION = 20000.00
const LIMITPERDAY = 100000.00
const FEE = 0.00

type Account struct {
	Name           string  `json:"name"`
	Id             string  `json:"id"`
	Balance        float64 `json:"balance"`
	TransferPerDay float64 `json:"transferperday"`
}

type TransferResponse struct {
	BalanceOld float64 `json:"balanceold"`
	BalanceNew float64 `json:"balancenew"`
	Withdrawal float64 `json:"withdrawal"`
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
func checkBalance(amount, transterMoney, fee float64) bool {
	return transterMoney+fee >= amount
}
func checkTransferPerDay(amount float64) bool {
	return amount > LIMITPERDAY
}

func checkTransferPerTransaction(checkTransferPerTransaction float64) bool {
	return checkTransferPerTransaction > LIMITPERTRANSACTION
}

func TransferProcess(accountFrom, accountTo Account, amount float64) TransferResponse {
	if checkBalance(accountFrom.Balance, amount, FEE) {
		return TransferResponse{
			BalanceOld: accountFrom.Balance,
			BalanceNew: accountFrom.Balance,
			Withdrawal: 0,
		}
	}
	if checkTransferPerTransaction(amount) {
		return TransferResponse{
			BalanceOld: accountFrom.Balance,
			BalanceNew: accountFrom.Balance,
			Withdrawal: 0,
		}
	}
	if checkTransferPerDay(amount + accountFrom.TransferPerDay) {
		return TransferResponse{
			BalanceOld: accountFrom.Balance,
			BalanceNew: accountFrom.Balance,
			Withdrawal: 0,
		}
	}
	return TransferResponse{
		BalanceOld: accountFrom.Balance,
		BalanceNew: (accountFrom.Balance - amount),
		Withdrawal: amount,
	}

}
