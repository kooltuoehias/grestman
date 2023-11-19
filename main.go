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
)


func main() {
    myApp := app.New()
    myWindow := myApp.NewWindow("gRestman")
    myWindow.Resize(fyne.NewSize(320, 640))
    vbox := container.NewGridWithRows(10)
    https := widget.NewCheck("https", func(value bool) {
      log.Println("Check set to", value)
    })
    https.SetChecked(true)
    inputUrl := widget.NewEntry()
    inputUrl.SetPlaceHolder("Enter url")

    apiKeyUrl := widget.NewEntry()
    apiKeyUrl.SetPlaceHolder("Enter system api key")

    label := widget.NewLabel("latency")


    button := widget.NewButton("Send", func() {
       response, latency := call(renderUrl(inputUrl.Text, https.Checked), apiKeyUrl.Text);
       data := [][]string{}
       for k, v := range response.Header {
	       value := strings.Join(v," ")
	       data = append(data, []string{k, value})
       }
       fmt.Println(len(data))
       table := widget.NewTable(
	    		func() (int, int) { return len(data), 2 },
			func() fyne.CanvasObject { return widget.NewLabel("Template")},
			func(i widget.TableCellID, o fyne.CanvasObject) {
				o.(*widget.Label).SetText(data[i.Row][i.Col])
			})
       label.SetText(fmt.Sprintf("latency: %dms", latency))
       vbox.Add(label)
       vbox.Add(table)
    })
    
    vbox.Add(https)
    vbox.Add(inputUrl)
    vbox.Add(apiKeyUrl)
    vbox.Add(button)
    myWindow.SetContent(vbox)
    myWindow.ShowAndRun()
}

func call(url string, key string) (*http.Response, int64) {
    client := &http.Client{}
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Println(err)
        return nil, 0
    }
    
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("x-api-key", key)
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
	if(isHttps) {
		return "https://" + url
	} else {
		return url
	}
}

