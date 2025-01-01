package front

import (
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func OutputForm(result [][]int) *fyne.Container {
	grid := container.NewVBox()

	for row := 0; row < 9; row++ {
		rowContainer := container.NewHBox()
		for col := 0; col < 9; col++ {
			cell := createLabeledCell(result[row][col])

			// 縦の区切り線
			if col > 0 && col%3 == 0 {
				rowContainer.Add(widget.NewSeparator())
			}

			rowContainer.Add(cell)
		}

		// 横の区切り線
		if row > 0 && row%3 == 0 {
			grid.Add(widget.NewSeparator())
		}

		grid.Add(rowContainer)
	}

	return grid
}

func createLabeledCell(value int) *fyne.Container {
	background := canvas.NewRectangle(color.White)
	background.SetMinSize(fyne.NewSize(50, 50))

	text := " "
	if value != 0 {
		text = strconv.Itoa(value)
	}
	label := widget.NewLabel(text)
	label.Alignment = fyne.TextAlignCenter

	return container.NewMax(background, label)
}
