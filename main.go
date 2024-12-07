package main

import "github.com/signintech/gopdf"

func main() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4}) //Page format

	pdf.AddPage() //add page

	err := pdf.AddTTFFont("Poppins", "./Poppins-Regular.ttf") //add font
	if err != nil {
		panic(err)
	}

	err = pdf.SetFont("Poppins", "", 24) //set font config
	if err != nil {
		panic(err)
	}

	pdf.Cell(nil, "Hello from Brazil") //create a cell of text
	err = pdf.WritePdf("brazil.pdf")
	if err != nil {
		panic(err)
	}
}
