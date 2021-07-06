package main
import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type tappableIcon struct {
	widget.Icon
}

func newTappableIcon(res fyne.Resource) *tappableIcon {
	icon := &tappableIcon{}
	icon.ExtendBaseWidget(icon)
	icon.SetResource(res)

	return icon
}
func main() {
a := app.New()
w := a.NewWindow("Tappable")
w.SetContent(newTappableIcon(theme.FyneLogo()))
w.ShowAndRun()
}