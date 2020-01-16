package main

import (
	"os"

	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
)

func display(entries Entries) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	for e := range entries {
		var m string
		m = text.WrapSoft(entries[e].Message, 50)
		t.AppendRow(table.Row{
			entries[e].getTimeObject().Format(lifeLogFormat), m,
		})
		t.AppendRow(table.Row{
			"", "",
		})
	}
	t.SetStyle(table.StyleColoredBright)
	t.Style().Color.Row = text.Colors{text.Reset, text.Reset}
	t.Style().Color.RowAlternate = text.Colors{text.Reset, text.Reset}
	t.Style().Color.Header = text.Colors{text.Reset, text.FgBlue}
	t.Style().Color.Footer = text.Colors{text.Reset, text.Reset}
	t.Style().Options.DrawBorder = false
	t.AppendFooter(table.Row{"", ""})
	t.SetAllowedRowLength(100)

	t.Render()
}
