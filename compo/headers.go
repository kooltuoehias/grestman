package compo

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Headers struct {
	headerList []string
}

func NewHeaders() Headers {
	return Headers {
		headerList: []string{}, 
	}
}

func (headers Headers) Offer() *fyne.Container {
	data := binding.BindStringList(
		&headers.headerList,
	)

	newItemEntry := widget.NewEntry()
	newItemEntry.SetPlaceHolder("New Header")
	list := widget.NewListWithData(data,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		})

	add := widget.NewButton("Add Header", func() {
		val := fmt.Sprintf(newItemEntry.Text)
		data.Append(val)
	})
	return container.NewBorder(newItemEntry, add, nil, nil, list)
}

func (headers Headers) HeaderList() []string {
	return headers.headerList
} 