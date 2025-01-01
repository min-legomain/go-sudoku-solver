package front

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func InputForm() (*fyne.Container, [][]*widget.Select) {
	grid := container.NewVBox()
	selectors := make([][]*widget.Select, 9) // セル情報を保持する

	for row := 0; row < 9; row++ {
		selectors[row] = make([]*widget.Select, 9) // 各行のセル情報
		rowContainer := container.NewHBox()
		for col := 0; col < 9; col++ {
			cell := widget.NewSelect([]string{" ", "1", "2", "3", "4", "5", "6", "7", "8", "9"}, nil)
			cell.PlaceHolder = " "     // 空白時はプレースホルダー
			selectors[row][col] = cell // セル情報を保持

			// 縦の区切り線
			if col > 0 && col%3 == 0 {
				rowContainer.Add(widget.NewSeparator())
			}

			rowContainer.Add(cell)
		}

		grid.Add(rowContainer)

		// 横の区切り線
		if row > 0 && row%3 == 2 {
			grid.Add(widget.NewSeparator())
		}
	}

	return grid, selectors
}

func GetGridValues(selectors [][]*widget.Select) [][]int {
	values := make([][]int, 9)
	for i := 0; i < 9; i++ {
		values[i] = make([]int, 9)
		for j := 0; j < 9; j++ {
			if selectors[i][j] != nil {
				if num, err := strconv.Atoi(selectors[i][j].Selected); err == nil {
					values[i][j] = num
				} else {
					values[i][j] = 0 //未入力の場合は0を返す
				}
			} else {
				values[i][j] = 0
			}
		}
	}
	return values
}
