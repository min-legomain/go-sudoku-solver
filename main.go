package main

import (
	"fmt"
	"go-sudoku-solver/front"
	"go-sudoku-solver/solver"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("SudokuSolver")

	inputForm, selectors := front.InputForm()
	outputForm := container.NewVBox()
	var executeButton, resetFormButton *widget.Button

	executeButton = front.ExecuteButton(func() {
		result, err := solver.Execute(front.GetGridValues(selectors))
		if err != nil {
			dialog.ShowError(err, w)
			fmt.Println("cant solve")
			return
		}
		outputForm = front.OutputForm(result)

		newContent := container.NewVBox(
			container.NewHBox(
				inputForm,
				outputForm,
			),
			resetFormButton,
		)
		w.SetContent(newContent)
	})

	resetFormButton = front.ResetFormButton(func() {
		for _, row := range selectors {
			for _, selector := range row {
				selector.SetSelected("")
				selector.Refresh()
			}
		}

		outputForm.Objects = nil
		outputForm.Refresh()

		inputForm, selectors = front.InputForm()

		newContent := container.NewVBox(
			container.NewHBox(
				inputForm,
				outputForm,
			),
			executeButton,
		)
		w.SetContent(newContent)
	})

	gridWidth := inputForm.MinSize().Width + 20
	gridHeight := inputForm.MinSize().Height + executeButton.MinSize().Height + 20
	w.Resize(fyne.NewSize(float32(gridWidth), float32(gridHeight)))

	content := container.NewVBox(
		container.NewGridWrap(fyne.NewSize(gridWidth, gridHeight), inputForm),
		executeButton,
	)
	w.SetContent(content)
	w.ShowAndRun()
}
