package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Define structs that match the JSON structure
type AccountInfo struct {
	IBAN string `json:"IBAN"`
	BIC  string `json:"BIC"`
	Note string `json:"Note"`
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	ZipCode string `json:"zip_code"`
	State  string `json:"state"`
}

type ClientInfo struct {
	Name        string        `json:"name"`
	Address     Address        `json:"address"`
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
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return Data{}, fmt.Errorf("error reading JSON file: %v", err)
	}

	// Variable to hold the unmarshaled data
	var data Data

	// Unmarshal the JSON data
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		return Data{}, fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	return data, nil
}

// func main() {
// 	// Call the function to read the JSON file
// 	data, err := readJSONFile("accounts_statement.json")
// 	if err != nil {
// 		log.Fatalf("Error: %v", err)
// 	}

// 	// Use the returned data
// 	fmt.Printf("Client Name: %s\n", data.ClientInfo.Name)
// 	fmt.Printf("Client Address: %s\n", data.ClientInfo.Address)
// 	for i, account := range data.ClientInfo.AccountInfo {
// 		fmt.Printf("Account %d IBAN: %s\n", i+1, account.IBAN)
// 		fmt.Printf("Account %d BIC: %s\n", i+1, account.BIC)
// 		fmt.Printf("Account %d Note: %s\n", i+1, account.Note)
// 	}

// 	fmt.Println("Balance Summary Headers:", data.BalanceSummary.Headers)
// 	for _, balance := range data.BalanceSummary.Data {
// 		fmt.Printf("Product: %s, Opening Balance: %s, Money Out: %s, Money In: %s, Closing Balance: %s\n",
// 			balance.Products, balance.OpeningBalance, balance.MoneyOut, balance.MoneyIn, balance.ClosingBalance)
// 	}

// 	fmt.Println("Transaction Headers:", data.Transactions.Headers)
// 	for _, transaction := range data.Transactions.Data {
// 		fmt.Printf("Date: %s, Description: %s, Money Out: %s, Money In: %s, Balance: %s\n",
// 			transaction.Date, transaction.Description, transaction.MoneyOut, transaction.MoneyIn, transaction.Balance)
// 	}
// }
