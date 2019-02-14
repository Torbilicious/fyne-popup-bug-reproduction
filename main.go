package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)

func main() {
	application := app.New()
	w := application.NewWindow("Test")
	dialog.ShowInformation("Information", text, w)
	w.SetContent(widget.NewVBox(
		widget.NewLabel(""),
	))
	w.ShowAndRun()

}
