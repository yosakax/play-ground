package main

import (
	// "fmt"
	"slices"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func buildApplication() *tview.Application {
	return tview.NewApplication()
}

func main() {
	// app := tview.NewApplication()
	app := buildApplication()
	table := tview.NewTable().
		SetBorders(true)
	lorem := strings.Split("Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.", " ")
	cols, rows := 15, 40
	word := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			color := tcell.ColorWhite
			if c < 1 || r < 1 {
				color = tcell.ColorYellow
			}
			table.SetCell(r, c,
				tview.NewTableCell(lorem[word]).
					SetTextColor(color).
					SetAlign(tview.AlignCenter))
			word = (word + 1) % len(lorem)
		}
	}
	table.SetFixed(1, 1)
	table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlQ:
			app.Stop()
		case tcell.KeyCtrlU:
			slices.Reverse(lorem)
			word := 0
			for r := 0; r < rows; r++ {
				for c := 0; c < cols; c++ {
					color := tcell.ColorWhite
					if c < 1 || r < 1 {
						color = tcell.ColorYellow
					}
					table.SetCell(r, c,
						tview.NewTableCell(lorem[word]).
							SetTextColor(color).
							SetAlign(tview.AlignCenter))
					word = (word + 1) % len(lorem)
				}
			}
		}
		return event
	})

	// table.Select(0, 0).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
	// 	fmt.Print(key)
	// 	switch key {
	// 	case tcell.KeyEnter:
	// 		table.SetSelectable(true, true)
	// 	case tcell.KeyBackspace:
	// 		app.Stop()

	// 	case tcell.KeyEscape:
	// 		// reverse table
	// 		slices.Reverse(lorem)
	// 		word := 0
	// 		for r := 0; r < rows; r++ {
	// 			for c := 0; c < cols; c++ {
	// 				color := tcell.ColorWhite
	// 				if c < 1 || r < 1 {
	// 					color = tcell.ColorYellow
	// 				}
	// 				table.SetCell(r, c,
	// 					tview.NewTableCell(lorem[word]).
	// 						SetTextColor(color).
	// 						SetAlign(tview.AlignCenter))
	// 				word = (word + 1) % len(lorem)
	// 			}
	// 		}
	// 	}
	// }).SetSelectedFunc(func(row int, column int) {
	// 	table.GetCell(row, column).SetTextColor(tcell.ColorRed)
	// 	table.SetSelectable(false, false)
	// })
	if err := app.SetRoot(table, true).SetFocus(table).Run(); err != nil {
		panic(err)
	}
}
