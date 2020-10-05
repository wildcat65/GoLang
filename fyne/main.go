package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	hello := widget.NewLabel("Hello Alan")
	w.SetContent(widget.NewVBox(
		hello,
		widget.NewButton("Hi", func() {
			hello.SetText("Welcom :)")
		}),
	))

	w.ShowAndRun()
}

