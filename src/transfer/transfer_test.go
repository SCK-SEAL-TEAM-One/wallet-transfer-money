package transfer

import "testing"

func Test_TransferService_Input_AccountNumberFrom_981751424_And_AccountNumberTo_9817571425_And_AmountTransfer_500_Should_Be_TransferResponse(t *testing.T) {
	accountNumberFrom := "981751424"
	accountNumberTo := "981751425"
	amountTransfer := 500.00
	expected := TransferResponse{
		BalanceOld: 20000.00,
		BalanceNew: 19500.00,
		Withdrawal: 500.00,
	}

	actual := TransferService(accountNumberFrom, accountNumberTo, amountTransfer)

	if expected != actual {
		t.Errorf("Should Be\n%v\nbut it got\n%v", expected, actual)
	}
}

func Test_checkBalance_Input_20000_Dot_00_And_500_Dot_00_And_0_Dot_00_Should_Be_False(t *testing.T) {
	amount := 20000.00
	fee := 0.00
	tarnsferMoney := 500.00
	expected := false

	actual := checkBalance(amount, tarnsferMoney, fee)

	if expected != actual {
		t.Errorf("Should be %v but got %v", expected, actual)
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

func Test_checkBalance_Input_20000_Dot_00_And_30000_Dot_00_And_0_Dot_00_Should_Be_True(t *testing.T) {
	amount := 20000.00
	fee := 0.00
	tarnsferMoney := 30000.00
	expected := true

	actual := checkBalance(amount, tarnsferMoney, fee)

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

func Test_TransferProcess_Input_From_981757424_To_981757425_Should_Be_TransferResult(t *testing.T) {
	accountFrom := Account{Name: "Panumars Seanto", Id: "981751424", Balance: 20000.00, TransferPerDay: 0.00}
	accountTo := Account{Name: "Piyanuch Ekpiyakool", Id: "981751425", Balance: 100000.00, TransferPerDay: 0.00}
	amount := 500.00
	expected := TransferResponse{
		BalanceOld: 20000.00,
		BalanceNew: 19500.00,
		Withdrawal: 500.00,
	}

	actual := TransferProcess(accountFrom, accountTo, amount)

	if expected != actual {
		t.Errorf("Should %v but got %v", expected, actual)
	}
}
func Test_checkTransferPerTransaction_Input_500_Dot_00_Should_Be_True(t *testing.T) {
	transferPerTransaction := 500.00
	expected := false

	actual := checkTransferPerTransaction(transferPerTransaction)

	if expected != actual {
		t.Errorf("Should be %t but got %t", expected, actual)
	}

}

func Test_Transfer_Input_20000_Dot_00_And_500_Dot_00_Should_Be_19500_Dot_00(t *testing.T) {
	balance := 20000.00
	transfer := 500.00
	expected := 19500.00

	actual := Transfer(balance, transfer)

	if expected != actual {
		t.Errorf("Should be %f but got %f", expected, actual)
	}
}
