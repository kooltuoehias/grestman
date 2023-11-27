package main

import (
	"fmt"
	"log"
    "net/http/httputil"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/kooltuoehias/grestman/compo"
	"github.com/kooltuoehias/grestman/rest"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("grestman")
	myWindow.Resize(fyne.NewSize(320, 640))
	vbox := container.NewGridWithRows(10)
	https := widget.NewCheck("https", func(value bool) {
		log.Println("Check set to", value)
	})
	https.SetChecked(true)

	ma := compo.NewHttpMethodAndAddress()
	ls := compo.NewLatencyAndStatusCode()
	headers := compo.NewHeaders()

	responseTextGrid := widget.NewTextGrid()
	var builder strings.Builder
	button := widget.NewButton("Send", func() {

		response, latency := rest.Call(rest.MethodI2S(ma.MethodIndex()), rest.RenderUrl(ma.Address(), https.Checked), headers.HeaderList())
		for k, v := range response.Header {
			value := strings.Join(v, " ")
			builder.WriteString(k + ": " + value + "\n")
		}
		ls.Latency(fmt.Sprintf("latency: %dms", latency))
		ls.Code(fmt.Sprintf("statusCode: %d", response.StatusCode))
		respDump, err := httputil.DumpResponse(response, true)
		if err != nil {
			responseTextGrid.SetText(fmt.Sprintf("%s\nError: %v", builder.String(), err))
		} else {
			responseTextGrid.SetText(builder.String() + "\n" + string(respDump))
		}
		builder.Reset()
	
	})

	vbox.Add(https)
	vbox.Add(ma.Offer())
	vbox.Add(headers.Offer())
	vbox.Add(button)
	vbox.Add(ls.Offer())
	vbox.Add(responseTextGrid)
	myWindow.SetContent(vbox)
	myWindow.ShowAndRun()
}

