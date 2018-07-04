package transfer

func checkBalance(amount, transterMoney, fee float64) bool {
	if transterMoney+fee >= amount {
		return false
	}

	return true
}
