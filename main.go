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
	// create a portrait, A4 pdf
	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	// set the page margins
	m.SetPageMargins(20, 20, 20)

	// add header
	buildHeading(m)

	// dummy row to force header rendering
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("This is a test row to ensure header rendering.", props.Text{
				Top:   3,
				Style: consts.Normal,
				Align: consts.Left,
				Color: getGreyColor(),
			})
		})
	})

	// output the file
	err := m.OutputFileAndClose("account_statement.pdf")
	if err != nil {
		fmt.Println("could not save pdf:", err)
		os.Exit(1)
	}

	fmt.Println("PDF generated successfully!")
}

// build the pdf heading using the logo
func buildHeading(m pdf.Maroto) {
	m.RegisterHeader(func() {
		m.Row(20, func() {
			m.Col(2, func() {
				err := m.FileImage("images/Revolut logo.png", props.Rect{
					Center:  false,
					Percent: 75,
					Top:     0,
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
				})

				m.Text("Generated on the 20 May 2023", props.Text{
					Top:   4.5,
					Style: consts.Normal,
					Size:  8,

					Align: consts.Right,
					Color: getGreyColor(),
				})

				m.Text("Revolut Bank UAB", props.Text{
					Top:   8,
					Style: consts.Normal,
					Size:  8,

					Align: consts.Right,
					Color: getGreyColor(),
				})
			})
		})
	})

}
