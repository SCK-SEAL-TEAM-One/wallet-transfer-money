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
		t.Errorf("Should Be\n%v\nbut it got\n%v", expected, actual)
	}
}

func Test_checkTransferPerDay_Input_500_Should_Be__False(t *testing.T) {
	amount := 500.00
	expected := false

	actual := checkTransferPerDay(amount)
	if expected != actual {
		t.Errorf("should be %t but got %t", expected, actual)
	}
}
func Test_getAccount_Input_981751424_Should_Be_Panumars_Seanto(t *testing.T) {

	accountNumber := "981751424"
	expected := Account{Name: "Panumars Seanto", Id: "981751424", Balance: 20000.00, TransferPerDay: 0.00}

	actual := getAccount(accountNumber)

	if expected != actual {
		t.Errorf("Should be %v but got %v", expected, actual)
	}

}
func Test_getAccount_Input_981751425_Should_Be_Piyanuch_Ekpiyakool(t *testing.T) {

	accountNumber := "981751425"
	expected := Account{Name: "Piyanuch Ekpiyakool", Id: "981751425", Balance: 100000.00, TransferPerDay: 0.00}

	actual := getAccount(accountNumber)

	if expected != actual {
		t.Errorf("Should be %v but got %v", expected, actual)
	}

}
