package front

import (
	"fyne.io/fyne/v2/widget"
)

func ResetFormButton(onClick func()) *widget.Button {
	return widget.NewButton("やり直す", onClick)
}
