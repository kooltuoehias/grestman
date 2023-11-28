package compo

import (
	"fmt"
    "net/http/httputil"
	"strings"

	"fyne.io/fyne/v2/widget"

	"github.com/kooltuoehias/grestman/interfaces"
	"github.com/kooltuoehias/grestman/rest"
)

type Sender struct {
	SendMethodProvider interfaces.MethodAble 
	SendUrlProvider    interfaces.AddressAble 
	SendHeadersProvider interfaces.HeaderListAble
	LatencyReceiver interfaces.SetextAble
	StatusCodeReceiver interfaces.SetextAble
	ResponseFullReceiver interfaces.SetextAble
}

var builder strings.Builder

func (sender Sender) Offer()  *widget.Button {
	return widget.NewButton("Send", func() {

		response, latency := rest.Call(sender.SendMethodProvider.Method(), sender.SendUrlProvider.Address(), sender.SendHeadersProvider.HeaderList())
		defer response.Body.Close()
		for k, v := range response.Header {
			value := strings.Join(v, " ")
			builder.WriteString(k + ": " + value + "\n")
		}
		sender.LatencyReceiver.SetText(fmt.Sprintf("latency: %dms", latency))
		sender.StatusCodeReceiver.SetText(fmt.Sprintf("statusCode: %d", response.StatusCode))
		respDump, err := httputil.DumpResponse(response, true)
		if err != nil {
			sender.ResponseFullReceiver.SetText(fmt.Sprintf("%s\nError: %v", builder.String(), err))
		} else {
			sender.ResponseFullReceiver.SetText(string(respDump))
		}
		builder.Reset()
	
	})
	
}