package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type AccountInfo struct {
	IBAN string `json:"IBAN"`
	BIC  string `json:"BIC"`
	Note string `json:"Note"`
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	ZipCode string `json:"zip_code"`
	State   string `json:"state"`
}

type ClientInfo struct {
	Name        string        `json:"name"`
	Address     Address       `json:"address"`
	AccountInfo []AccountInfo `json:"account_info"`
}

type BalanceSummaryData struct {
	Products       string `json:"Products"`
	OpeningBalance string `json:"Opening Balance"`
	MoneyOut       string `json:"Money out"`
	MoneyIn        string `json:"Money in"`
	ClosingBalance string `json:"Closing balance"`
}

type BalanceSummary struct {
	Headers []string             `json:"headers"`
	Data    []BalanceSummaryData `json:"data"`
}

type TransactionData struct {
	Date        string `json:"Date"`
	Description string `json:"Description"`
	MoneyOut    string `json:"Money out"`
	MoneyIn     string `json:"Money in"`
	Balance     string `json:"Balance"`
}

type Transactions struct {
	Headers []string          `json:"headers"`
	Data    []TransactionData `json:"data"`
}

type Data struct {
	ClientInfo     ClientInfo     `json:"client_info"`
	BalanceSummary BalanceSummary `json:"balance_summary"`
	Transactions   Transactions   `json:"transactions"`
}

// Function to read JSON file and return the data struct
func readJSONFile(filename string) (Data, error) {
	// Open the JSON file
	jsonFile, err := os.Open(filename)
	if err != nil {
		return Data{}, fmt.Errorf("error opening JSON file: %v", err)
	}
	defer jsonFile.Close()

	// Read the file contents
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return Data{}, fmt.Errorf("error reading JSON file: %v", err)
	}

	var data Data

	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		return Data{}, fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	return data, nil
}
