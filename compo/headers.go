package compo

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Headers struct {
	data binding.ExternalStringList 
}

func NewHeaders() Headers {
	return Headers {
		data: binding.BindStringList(
			&[]string{},
		),
	}
}

func (headers Headers) Offer() *fyne.Container {
	newItemEntry := widget.NewEntry()
	newItemEntry.SetPlaceHolder("New Header")
	list := widget.NewListWithData(headers.data,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		})

	add := widget.NewButton("Add Header", func() {
		val := fmt.Sprintf(newItemEntry.Text)
		headers.data.Append(val)
	})
	return container.NewBorder(newItemEntry, add, nil, nil, list)
}

func (headers Headers) HeaderList() []string {
	result,_ := headers.data.Get()
	return result;
} 