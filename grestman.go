package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/kooltuoehias/grestman/compo"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("grestman")
	myWindow.Resize(fyne.NewSize(320, 640))
	//vbox := container.NewGridWithRows(10)
	vbox := container.New(layout.NewVBoxLayout())
	https := widget.NewCheck("https", func(value bool) {
		log.Println("Check set to", value)
	})
	https.SetChecked(true)

	ma := compo.NewHttpMethodAndAddress()
	ls := compo.NewLatencyAndStatusCode()
	headers := compo.NewHeaders()

	responseTextGrid := widget.NewTextGrid()
	senderButton := compo.Sender{
		SendMethodProvider:  ma,
		SendUrlProvider:    ma,
		SendHeadersProvider: headers,
		LatencyReceiver: ls.LatencyLabel, 
		StatusCodeReceiver: ls.StatusCodeLabel, 
		ResponseFullReceiver: responseTextGrid, 
	}

	vbox.Add(https)
	vbox.Add(ma.Offer())
	vbox.Add(headers.Offer())
	vbox.Add(senderButton.Offer())
	vbox.Add(ls.Offer())
	vbox.Add(responseTextGrid)
	myWindow.SetContent(vbox)
	myWindow.ShowAndRun()
}

