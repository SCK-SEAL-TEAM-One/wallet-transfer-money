package transfer

import "testing"

func Test_checkBalance_Input_20000_Dot_00_And_500_Dot_00_And_0_Dot_00_Should_Be_True(t *testing.T) {
	amount := 20000.00
	fee := 0.00
	tarnsferMoney := 500.00
	expected := true

	actual := checkBalance(amount, tarnsferMoney, fee)

	if expected != actual {
		t.Errorf("Should be %v but got %v", expected, actual)
	}

}

func Test_checkBalance_Input_20000_Dot_00_And_30000_Dot_00_And_0_Dot_00_Should_Be_False(t *testing.T) {
	amount := 20000.00
	fee := 0.00
	tarnsferMoney := 30000.00
	expected := false

	actual := checkBalance(amount, tarnsferMoney, fee)

	if expected != actual {
		t.Errorf("Should be %v but got %v", expected, actual)
	}

}
