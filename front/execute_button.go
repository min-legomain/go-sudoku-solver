package front

import (
	"fyne.io/fyne/v2/widget"
)

func ExecuteButton(onClick func()) *widget.Button {
	return widget.NewButton("計算する", onClick)
}
