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

func getBalanceSummaryTable(balanceSummary BalanceSummary) (headers []string, data [][]string) {
	if len(balanceSummary.Data) == 0 {
		return nil, nil
	}

	headers = balanceSummary.Headers

	data = make([][]string, len(balanceSummary.Data))
	for i, d := range balanceSummary.Data {
		data[i] = []string{
			d.Products,
			d.OpeningBalance,
			d.MoneyOut,
			d.MoneyIn,
			d.ClosingBalance,
		}
	}

	return headers, data
}
