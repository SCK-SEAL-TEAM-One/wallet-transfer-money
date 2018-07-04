package transfer

func checkTransferPerDay(amount float64) bool {
	if amount >= 100000.00 {
		return false
	}
	return true
}
