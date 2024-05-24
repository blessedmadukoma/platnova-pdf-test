package main

import (
	"fmt"
	"log"
	"os"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
)

func main() {
	data, err := readJSONFile("account_statement.json")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	m.SetPageMargins(20, 20, 20)

	buildHeading(m)

	buildClientDetails(m, data.ClientInfo)

	buildClientAccountInfo(m, data.ClientInfo.AccountInfo)

	buildBalanceSummary(m, data.BalanceSummary)

	buildAccountTransactions(m, data.Transactions)

	err = m.OutputFileAndClose("account_statement.pdf")
	if err != nil {
		fmt.Println("could not save pdf:", err)
		os.Exit(1)
	}

	fmt.Println("PDF generated successfully!")
}
