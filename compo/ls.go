package compo

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type LatencyAndStatusCode struct {
	LatencyLabel    *widget.Label
	StatusCodeLabel *widget.Label
}

func NewLatencyAndStatusCode() LatencyAndStatusCode {
	return LatencyAndStatusCode{
		LatencyLabel:    widget.NewLabel("latency"),
		StatusCodeLabel: widget.NewLabel("statusCode"),
	}
}

func (ls LatencyAndStatusCode) Latency(text string) {
	ls.LatencyLabel.SetText(text)
}

func (ls LatencyAndStatusCode) Code(text string) {
	ls.StatusCodeLabel.SetText(text)
}

func (ls LatencyAndStatusCode) Offer() *fyne.Container {
	return container.New(layout.NewHBoxLayout(), ls.LatencyLabel, layout.NewSpacer(), ls.StatusCodeLabel)
}
