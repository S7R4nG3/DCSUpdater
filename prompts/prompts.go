package prompts

import (
	"fmt"
	"strings"

	"github.com/S7R4nG3/DCSUpdater/flags"
	"github.com/S7R4nG3/DCSUpdater/run"
	"github.com/S7R4nG3/DCSUpdater/types"
	"github.com/lxn/walk"
	declare "github.com/lxn/walk/declarative"
)

type UI struct {
	UpdaterPath string
	InstallPath string
	CLiFlags    []string
}

func (ui *UI) Prompt() {
	var out *walk.TextEdit

	declare.MainWindow{
		Title: "DCS Updater",
		Size: declare.Size{
			Width:  500,
			Height: 500,
		},
		Layout: declare.VBox{
			Spacing: 10,
		},
		Children: []declare.Widget{
			declare.Label{
				Font: declare.Font{
					Bold: true,
				},
				Text: "Located DCS Install at: ",
			},
			declare.Label{
				Alignment: declare.AlignHCenterVNear,
				Text:      ui.InstallPath,
			},
			declare.Label{
				Text: "Please select how you would like to update DCS:",
			},
			declare.RadioButtonGroup{
				DataMember: "vers",
				Buttons: []declare.RadioButton{
					{
						Text:       "DCS Open Beta",
						Alignment:  declare.AlignHNearVCenter,
						ColumnSpan: 10,
						RowSpan:    10,
						OnClicked: func() {
							ui.CLiFlags = flags.Get(types.OpenBeta)
						},
					},
					{
						Text:       "DCS Stable Release",
						Alignment:  declare.AlignHNearVCenter,
						ColumnSpan: 10,
						OnClicked: func() {
							ui.CLiFlags = flags.Get(types.Stable)
						},
					},
					{
						Text:       "DCS Open Alpha",
						Alignment:  declare.AlignHNearVCenter,
						ColumnSpan: 10,
						OnClicked: func() {
							ui.CLiFlags = flags.Get(types.OpenAlpha)
						},
					},
					{
						Text:      "Repair",
						Alignment: declare.AlignHNearVCenter,
						OnClicked: func() {
							ui.CLiFlags = flags.Get(types.Repair)
						},
					},
					{
						Text:      "Cleanup",
						Alignment: declare.AlignHNearVCenter,
						OnClicked: func() {
							ui.CLiFlags = flags.Get(types.Cleanup)
						},
					},
				},
			},
			declare.PushButton{
				Text: "Update",
				Font: declare.Font{
					Bold: true,
				},
				OnClicked: func() {
					out.SetReadOnly(true)
					out.AppendText("Starting Updater...")
					out.AppendText(fmt.Sprintf("\r\nRunning Command:\r\n\t%s %s", ui.UpdaterPath, strings.Join(ui.CLiFlags[:], " ")))
					out.AppendText(run.Cmd(ui.UpdaterPath, ui.CLiFlags))
					out.AppendText("\r\nSUCCESS!")
				},
			},
			declare.TextEdit{
				Name:     "Log",
				AssignTo: &out,
				ReadOnly: true,
				MaxSize: declare.Size{
					Width:  100,
					Height: 100,
				},
			},
		},
	}.Run()
}
