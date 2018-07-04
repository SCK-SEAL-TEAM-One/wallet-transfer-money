package transfer

import "testing"

func Test_TransferService_Input_AccountNumberFrom_9817571424_And_AccountNumberTo_9817571425_And_AmountTransfer_500_Should_Be_TransferResponse(t *testing.T) {
	accountNumberFrom := "9817571424"
	accountNumberTo := "9817571425"
	amountTransfer := 500.00
	expected := TransferResponse{
		BalanceOld: 20000.00, BalanceNew: 19500.00, Withdrawal: 500.00,
	}

	actual := TransferService(accountNumberFrom, accountNumberTo, amountTransfer)

	if expected != actual {
		t.Errorf("Should Be\n%v but it got\n%v", expected, actual)
	}
}
