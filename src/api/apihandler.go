package api

import (
	"encoding/json"
	"net/http"
)

type TransferRequest struct {
	AccountNumberFrom string  `json:"accountNumberFrom"`
	AccountNumberTo   string  `json:"accountNumberTo"`
	AccountTransfer   float64 `json:"accountTransfer"`
}

type TransferResponse struct {
	BalanceOld float64 `json:"balanceold"`
	BalanceNew float64 `json:"balancenew"`
	Withdrawal float64 `json:"withdrawal"`
}

func TransferHandler(responseWriter http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var transferRequest TransferRequest
	err := decoder.Decode(&transferRequest)

	if err != nil {
		http.Error(responseWriter, err.Error(), 500)
		return
	}

	transferResponse := TransferResponse{}
	transferResponseJSON, _ := json.Marshal(transferResponse)
	responseWriter.Write(transferResponseJSON)
}
