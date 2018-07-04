package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func Test_TransferHandler_Input_From_9817571424_And_To_9817571425_Should_Be_Balanceold_20000_Dot_00_And_Balancenew_19500_Dot_00_And_withdrawal_500_Dot_00(t *testing.T) {
	url := "/transfer"
	requestBody := map[string]interface{}{
		"accountNumberFrom": "981751424",
		"accountNumberTo":   "981751425",
		"amountTransfer":    500.00,
	}
	requestBodyString, _ := json.Marshal(requestBody)
	request := httptest.NewRequest("POST", url, bytes.NewBuffer(requestBodyString))
	responseRecorder := httptest.NewRecorder()
	expected := `{"balanceold":20000,"balancenew":19500,"withdrawal":500}`

	TransferHandler(responseRecorder, request)
	response := responseRecorder.Result()
	actual, _ := ioutil.ReadAll(response.Body)

	if string(actual) != expected {
		t.Errorf("Should Be\n%s\nbut it got\n%s", expected, actual)
	}

}
