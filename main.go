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

	// create a portrait, A4 pdf
	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	// set the page margins
	m.SetPageMargins(20, 20, 20)

	// add header
	buildHeading(m)

	// add client details
	buildClientDetails(m, data.ClientInfo)

	// client account info
	buildClientAccountInfo(m, data.ClientInfo.AccountInfo)

	// balance summary
	buildBalanceSummary(m, data.BalanceSummary)

	// output the file
	err = m.OutputFileAndClose("account_statement.pdf")
	if err != nil {
		fmt.Println("could not save pdf:", err)
		os.Exit(1)
	}

	fmt.Println("PDF generated successfully!")
}
