package main

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/container/popup"
	"fyne.io/fyne/v2/container/storage"
	"fyne.io/fyne/v2/container/tabcontainer"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/atotto/clipboard"
	"os"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Clipboard Manager")

	historyTab := container.NewVBox()
	historyList := widget.NewList(
		func() int {
			return len(clipboardHistory)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(index int, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(clipboardHistory[index])
		},
	)

	historyList.OnSelected = func(index int) {
		if index >= 0 && index < len(clipboardHistory) {
			copyToClipboard(clipboardHistory[index])
			showToast("Text pasted to clipboard.")
		}
	}

	historyTab.Add(historyList)

	aboutTab := container.NewVBox(
		widget.NewLabel("Simple Clipboard Manager in Go"),
		widget.NewLabel("Powered by Fyne UI"),
	)

	tabs := tabcontainer.New(
		tabcontainer.NewTabItemWithIcon("History", theme.ViewRefreshIcon(), historyTab),
		tabcontainer.NewTabItemWithIcon("About", theme.InfoIcon(), aboutTab),
	)

	myWindow.SetContent(container.NewBorder(nil, nil, nil, nil, tabs))

	myWindow.ShowAndRun()
}

func copyToClipboard(text string) {
	clipboard.WriteAll(text)
}

func showToast(message string) {
	dialog.ShowInformation("Clipboard Manager", message, myWindow)
}
