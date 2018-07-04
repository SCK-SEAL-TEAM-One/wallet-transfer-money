package transfer

import "testing"

func Test_checkTransferPerDay_Input_500_Should_Be__True(t *testing.T) {
	amount := 500.00
	expected := true

	actual := checkTransferPerDay(amount)
	if expected != actual {
		t.Errorf("should be %t but got %t", expected, actual)
	}
}
