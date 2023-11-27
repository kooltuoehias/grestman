package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/kooltuoehias/grestman/compo"
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

	button := widget.NewButton("Send", func() {

		response, latency := call(methodI2S(ma.MethodIndex()), renderUrl(ma.Address(), https.Checked), headers.HeaderList())
		data := [][]string{}
		for k, v := range response.Header {
			value := strings.Join(v, " ")
			data = append(data, []string{k, value})
		}
		table := widget.NewTable(
			func() (int, int) { return len(data), 2 },
			func() fyne.CanvasObject { return widget.NewLabel("Template") },
			func(i widget.TableCellID, o fyne.CanvasObject) {
				o.(*widget.Label).SetText(data[i.Row][i.Col])
			})
		ls.Latency(fmt.Sprintf("latency: %dms", latency))
		ls.Code(fmt.Sprintf("statusCode: %d", response.StatusCode))
		vbox.Add(ls.Offer())
		vbox.Add(table)
	})

	vbox.Add(https)
	vbox.Add(ma.Offer())
	vbox.Add(headers.Offer())
	vbox.Add(button)
	myWindow.SetContent(vbox)
	myWindow.ShowAndRun()
}

func call(method string, url string, headerList []string) (*http.Response, int64) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return nil, 0
	}

	req.Header.Set("Content-Type", "application/json")
	for _, v := range headerList {
		headerAndValue := strings.Split(v, ":")
		req.Header.Set(headerAndValue[0], headerAndValue[1])

	}
	start := time.Now()
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, 0
	}
	defer resp.Body.Close()
	elapsed := time.Since(start).Milliseconds()
	fmt.Printf("Request took %d milliseconds\n", elapsed)

	return resp, elapsed
}

func renderUrl(url string, isHttps bool) string {
	if isHttps {
		return "https://" + url
	} else {
		return url
	}
}

func methodI2S(v int) string {
	switch v {
		case 0:
			return "GET"
		case 1:
			return "POST"
		case 2:
			return "PUT"
		case 3:
			return "PATCH"
		case 5:
			return "DELETE"
		default:
			return "GET"
	}
}
