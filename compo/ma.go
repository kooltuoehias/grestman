package compo

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)


type HttpMethodAndAddress struct {
	HttpMethodOptions *widget.Select
	HttpAddress       *widget.Entry
}

func NewHttpMethodAndAddress() HttpMethodAndAddress {
	inputUrl := widget.NewEntry()
	inputUrl.SetPlaceHolder("Enter url")
	combo := widget.NewSelect([]string{"GET", "POST", "PUT", "PATCH", "DELETE"}, func(value string) {
		log.Println("Select set to", value)
	})
	return HttpMethodAndAddress{
		HttpMethodOptions: combo,
		HttpAddress:       inputUrl,
	}
}

func (ma HttpMethodAndAddress) MethodIndex() int {
	return ma.HttpMethodOptions.SelectedIndex() 
}

func (ma HttpMethodAndAddress) Address() string {
	return ma.HttpAddress.Text
}

func (ma HttpMethodAndAddress) Offer() *fyne.Container {
	return container.NewBorder(nil, nil, ma.HttpMethodOptions, nil, ma.HttpAddress)
}
