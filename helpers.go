package main

import "github.com/johnfercher/maroto/pkg/color"

// generate grey color
func getGreyColor() color.Color {
	return color.Color{
		Red:   128,
		Green: 128,
		Blue:  128,
	}
}

func getAccountInfoTable(accountInfo []AccountInfo) ([][]string, []string) {
	var table [][]string
	headers := []string{"IBAN", "BIC", "Note"}

	for _, info := range accountInfo {
		row := []string{info.IBAN, info.BIC, info.Note}
		table = append(table, row)
	}

	return table, headers
}
