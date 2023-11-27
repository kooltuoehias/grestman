package compo

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var methods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
var httpsPrefix = "https://"

type HttpMethodAndAddress struct {
	HttpMethodOptions *widget.Select
	HttpAddress       *widget.Entry
	HttpsCheck        *widget.Check
}

func NewHttpMethodAndAddress(checked *widget.Check) HttpMethodAndAddress {
	inputUrl := widget.NewEntry()
	inputUrl.SetPlaceHolder("Enter url")
	combo := widget.NewSelect(methods, func(value string) {})
	return HttpMethodAndAddress{
		HttpMethodOptions: combo,
		HttpAddress:       inputUrl,
		HttpsCheck:		   checked,
	}
}

func (ma HttpMethodAndAddress) Method() string {
	index := ma.HttpMethodOptions.SelectedIndex()
	if index == -1 {
		return methods[0]
	}
	return methods[index]
}

func (ma HttpMethodAndAddress) Address() string {
	return ma.renderUrl(ma.HttpAddress.Text, ma.HttpsCheck.Checked)
}

func (ma HttpMethodAndAddress) Offer() *fyne.Container {
	return container.NewBorder(nil, nil, ma.HttpMethodOptions, nil, ma.HttpAddress)
}

func (ma HttpMethodAndAddress) renderUrl(url string, isHttps bool) string {
	if isHttps && !strings.HasPrefix(url, httpsPrefix) {
		return httpsPrefix + url
	} else {
		return url
	}
}