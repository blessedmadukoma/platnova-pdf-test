package main

import (
	"fmt"
	"log"
	"os"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
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

	// output the file
	err = m.OutputFileAndClose("account_statement.pdf")
	if err != nil {
		fmt.Println("could not save pdf:", err)
		os.Exit(1)
	}

	fmt.Println("PDF generated successfully!")
}

// build the pdf heading using the logo
func buildHeading(m pdf.Maroto) {
	m.RegisterHeader(func() {
		m.Row(30, func() {
			m.Col(2, func() {
				err := m.FileImage("images/Revolut logo.png", props.Rect{
					Center:  false,
					Percent: 75,
					Top:     2,
				})

				if err != nil {
					log.Fatal("Image file was not loaded:", err)
				}
			})

			m.ColSpace(7)

			m.Col(3, func() {
				m.Text("EUR Statement", props.Text{
					Top:   0,
					Style: consts.Bold,
					Align: consts.Right,
					Size:  14,
				})

				m.Text("Generated on the 20 May 2023", props.Text{
					Top:   6,
					Style: consts.Normal,
					Size:  8,

					Align: consts.Right,
					Color: getGreyColor(),
				})

				m.Text("Revolut Bank UAB", props.Text{
					Top:   10,
					Style: consts.Normal,
					Size:  8,

					Align: consts.Right,
					Color: getGreyColor(),
				})
			})
		})
	})
}

func buildClientDetails(m pdf.Maroto, clientInfo ClientInfo) {
	m.Row(25, func() {
		m.Col(12, func() {
			m.Text(clientInfo.Name, props.Text{
				Top:   0,
				Style: consts.Bold,
				Align: consts.Left,
				Size:  12,
			})

			m.Text(clientInfo.Address.Street, props.Text{
				Top:   7,
				Style: consts.Bold,
				Align: consts.Left,
				Size:  10,
			})

			m.Text(clientInfo.Address.City, props.Text{
				Top:   12,
				Style: consts.Bold,
				Align: consts.Left,
				Size:  10,
			})

			m.Text(clientInfo.Address.State, props.Text{
				Top:   17,
				Style: consts.Bold,
				Align: consts.Left,
				Size:  10,
			})

			m.Text(clientInfo.Address.ZipCode, props.Text{
				Top:   22,
				Style: consts.Bold,
				Align: consts.Left,
				Size:  10,
			})
		})
	})
}

func buildClientAccountInfo(m pdf.Maroto, accountInfo []AccountInfo) {
	for _, info := range accountInfo {
		m.Row(4, func() {
			m.ColSpace(6)
			m.Col(6, func() {
				m.Text("IBAN: ", props.Text{
					Top:   0,
					Style: consts.Bold,
					Align: consts.Left,
					Size:  9,
				})
				m.Text(info.IBAN, props.Text{
					Top:   0,
					Left:  15,
					Style: consts.Normal,
					Align: consts.Left,
					Size:  9,
					Color: getGreyColor(),
				})
			})
		})

		m.Row(4, func() {
			m.ColSpace(6)
			m.Col(6, func() {
				m.Text("BIC: ", props.Text{
					Top:   0,
					Style: consts.Bold,
					Align: consts.Left,
					Size:  9,
				})
				m.Text(info.BIC, props.Text{
					Top:   0,
					Left:  15,
					Style: consts.Normal,
					Align: consts.Left,
					Size:  9,
					Color: getGreyColor(),
				})
			})
		})

		if info.Note != "" {
			m.Row(4, func() {
				m.ColSpace(6)
				m.Col(6, func() {
					m.Text(info.Note, props.Text{
						Top:   0,
						Style: consts.Normal,
						Left:  15,
						Align: consts.Left,
						Size:  8,
						Color: getGreyColor(),
					})
				})
			})
		}

		// a blank space between account information
		m.Row(5, func() {
			m.ColSpace(12)
		})
	}
}
