package gui

import (
	"strings"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func Left2Right() {
	var inTE, outTE *walk.TextEdit

	MainWindow{
		Title:   "SCREAMO",
		MinSize: Size{Width: 600, Height: 400},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inTE},
					TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
			PushButton{
				Text: "SCREAM",
				OnClicked: func() {
					outTE.SetText(strings.ToUpper(inTE.Text()))
				},
			},
		},
	}.Run()
}
