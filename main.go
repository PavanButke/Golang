package main

import (
	// "strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	// "github.com/Knetic/govaluate"
)

func main() {

	calc := app.New()
	w := calc.NewWindow("Calcy")

	output := ""
	input := widget.NewLabel("Choko Choko Rokdo")

	hisBtn := widget.NewButton("MRC", func() {

	})

	dltBtn := widget.NewButton("DEL", func() {

	})

	opnBtn := widget.NewButton("(", func() {
		output = output + "("

	})

	divBtn := widget.NewButton("/", func() {

	})
	/*  */
	svnBtn := widget.NewButton("7", func() {

	})
	eghtBtn := widget.NewButton("8", func() {

	})
	nineBtn := widget.NewButton("9", func() {

	})
	fourBtn := widget.NewButton("4", func() {

	})

	fiveBtn := widget.NewButton("5", func() {

	})

	sixBtn := widget.NewButton("6", func() {

	})

	threeBtn := widget.NewButton("3", func() {

	})

	twoBtn := widget.NewButton("2", func() {

	})
	oneBtn := widget.NewButton("1", func() {

	})

	strBtn := widget.NewButton("*", func() {

	})
	plsBtn := widget.NewButton("+", func() {

	})

	minsBtn := widget.NewButton("-", func() {

	})
	eqlBtn := widget.NewButton("=", func() {

	})

	zeroBtn := widget.NewButton("0", func() {

	})

	dotBtn := widget.NewButton(".", func() {

	})

	clrBtn := widget.NewButton("CLR", func() {
		output = ""
		input.SetText(output)
	})

	w.SetContent(container.NewVBox(
		input,
		container.NewGridWithColumns(1,
			container.NewGridWithColumns(2,
				hisBtn,
				dltBtn,
			),
			container.NewGridWithColumns(4,
				clrBtn,
				opnBtn,
				clrBtn,
				divBtn,
			),
			container.NewGridWithColumns(4,
				svnBtn,
				eghtBtn,
				nineBtn,
				plsBtn,
			),

			container.NewGridWithColumns(4,
				fourBtn,
				fiveBtn,
				sixBtn,
				strBtn,
			),

			container.NewGridWithColumns(4,
				threeBtn,
				twoBtn,
				oneBtn,
				minsBtn,
			),

			container.NewGridWithColumns(2,
				container.NewGridWithColumns(2,
					zeroBtn,
					dotBtn,
				),
				eqlBtn,
			),
		),
	))

	// myWindo.SetContent(container.NewVBox(
	// 	history,
	// 	ip,

	// 	container.NewGridWithColumns(1,
	// 		container.NewGridWithColumns(2,
	// 			historyBtn,
	// 			backBtn,
	// 		),
	// 		container.NewGridWithColumns(4,
	// 			clearBtn,
	// 			oBtn,
	// 			cBtn,
	// 			divBtn,
	// 		),
	// 		container.NewGridWithColumns(4,
	// 			svnBtn,
	// 			ethBtn,
	// 			nineBtn,
	// 			mulBtn,
	// 		),
	// 		container.NewGridWithColumns(4,
	// 			frthBtn,
	// 			fivBtn,
	// 			sixBtn,
	// 			minsBtn,
	// 		),
	// 		container.NewGridWithColumns(4,
	// 			oneBtn,
	// 			twoBtn,
	// 			treBtn,
	// 			plsBtn,
	// 		),
	// 		container.NewGridWithColumns(2,
	// 			container.NewGridWithColumns(2,
	// 				zroBtn,
	// 				dotBtn,
	// 			),
	// 			eqlBtn,
	// 		),
	// 	),
	// ))
	// //
	// myWindo.ShowAndRun()

	w.ShowAndRun()
}
