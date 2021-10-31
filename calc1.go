package main

import (
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
)

func showCalcApp() {

	calc := app.New()
	w := calc.NewWindow("Calcy")
	output := ""
	input := widget.NewLabel("Choko Choko Rokdo")
	isHistory := false
	historyStr := ""
	history := widget.NewLabel(historyStr)
	var historyArr []string

	hisBtn := widget.NewButton("MRC", func() {

		if isHistory {
			historyStr = ""
		} else {
			for i := len(historyArr) - 1; i >= 0; i-- {
				historyStr += historyStr + historyArr[i]
				historyStr += "\n"

			}
		}
		isHistory = !isHistory
		history.SetText(historyStr)

	})

	dltBtn := widget.NewButton("DEL", func() {

		if len(output) > 0 {

			output = output[:len(output)-1]
			input.SetText(output)
		}

	})

	opnBtn := widget.NewButton("(", func() {
		output = output + "("
		input.SetText(output)

	})

	clsBtn := widget.NewButton("(", func() {
		output = output + "("
		input.SetText(output)

	})

	divBtn := widget.NewButton("/", func() {
		output = output + ")"
		input.SetText(output)
	})

	svnBtn := widget.NewButton("7", func() {
		output = output + "7"
		input.SetText(output)
	})
	eghtBtn := widget.NewButton("8", func() {

		output = output + "8"
		input.SetText(output)

	})
	nineBtn := widget.NewButton("9", func() {

		output = output + "9"
		input.SetText(output)
	})
	fourBtn := widget.NewButton("4", func() {
		output = output + "4"
		input.SetText(output)
	})

	fiveBtn := widget.NewButton("5", func() {
		output = output + "5"
		input.SetText(output)
	})

	sixBtn := widget.NewButton("6", func() {

		output = output + "6"
		input.SetText(output)

	})

	threeBtn := widget.NewButton("3", func() {

		output = output + "3"
		input.SetText(output)

	})

	twoBtn := widget.NewButton("2", func() {
		output = output + "2"
		input.SetText(output)
	})
	oneBtn := widget.NewButton("1", func() {
		output = output + "1"
		input.SetText(output)
	})

	strBtn := widget.NewButton("*", func() {

		output = output + "*"
		input.SetText(output)

	})
	plsBtn := widget.NewButton("+", func() {
		output = output + "+"
		input.SetText(output)
	})

	minsBtn := widget.NewButton("-", func() {
		output = output + "-"
		input.SetText(output)

	})
	eqlBtn := widget.NewButton("=", func() {
		expression, err := govaluate.NewEvaluableExpression(output)
		if err == nil {
			result, err := expression.Evaluate(nil)

			if err == nil {
				ans := strconv.FormatFloat(result.(float64), 'f', -1, 64)
				strToAppend := output + " = " + ans
				historyArr = append(historyArr, strToAppend)
				output = ans
			} else {
				output = "error"
			}
		} else {
			output = "error"
		}
		input.SetText(output)

	})

	zeroBtn := widget.NewButton("0", func() {
		output = output + "0"
		input.SetText(output)
	})

	dotBtn := widget.NewButton(".", func() {
		output = output + "."
		input.SetText(output)
	})

	clrBtn := widget.NewButton("CLR", func() {
		output = ""
		input.SetText(output)
	})

	w.SetContent(container.NewVBox(
		input,
		history,
		container.NewGridWithColumns(1,
			container.NewGridWithColumns(2,
				hisBtn,
				dltBtn,
			),
			container.NewGridWithColumns(4,
				clrBtn,
				opnBtn,
				clsBtn,
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
