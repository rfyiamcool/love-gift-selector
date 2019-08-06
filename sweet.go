package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var (
	fakeRow = 6

	data = []string{
		"[1]  N手宝马",
		"[2]  口红",
		"[3]  sk2 神仙水",
		"[4]  钻戒",
		"[5]  驴包",
		"[6]  爱情之吻",
		"[7]  万元红包",
		"[8]  千元红包",
		"[9]  521红包",
		"[10] 大包sod秘",
		"[11]  一手保时泰",
	}
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize ui: %v", err)
	}
	defer ui.Close()

	p := widgets.NewParagraph()
	p.Title = "选中的礼物"
	p.Text = "..."
	p.TextStyle.Fg = ui.ColorWhite
	p.BorderStyle.Fg = ui.ColorCyan
	p.SetRect(30, 5, 80, 8)
	updateParagraph := func(count int) {
		switch count % 5 {
		case 0:
			p.TextStyle.Fg = ui.ColorRed
		case 1:
			p.TextStyle.Fg = ui.ColorWhite
		case 2:
			p.TextStyle.Fg = ui.ColorBlue
		case 3:
			p.TextStyle.Fg = ui.ColorGreen
		case 4:
			p.TextStyle.Fg = ui.ColorMagenta
		}
	}

	l := widgets.NewList()
	l.Title = "七夕 礼物选择器"
	l.Rows = data
	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = false
	l.SetRect(10, 10, 100, 25)
	ui.Render(l, p)

	uiEvents := ui.PollEvents()

	go func() {
		for index := 0; index < 100; index++ {
			if l.SelectedRow == 0 {
				l.ScrollHalfPageDown()
				ui.Render(l)
				continue
			}

			if l.SelectedRow == len(l.Rows)-1 {
				l.ScrollHalfPageUp()
				ui.Render(l, p)
				continue
			}

			n := rand.Intn(9)
			switch n {
			case 0, 1:
				l.ScrollDown()
			case 2, 3:
				l.ScrollUp()
			case 4, 5, 6:
				l.ScrollHalfPageUp()
			case 7, 8, 9:
				l.ScrollHalfPageDown()
			}

			ui.Render(l, p)
			time.Sleep(100 * time.Millisecond)
		}

		l.ScrollTop()
		for index := 1; index < fakeRow; index++ {
			l.ScrollDown()
			p.Text = data[index]
			ui.Render(l, p)
		}
	}()

	go func() {
		tickerCount := 0
		tickerCount++
		ticker := time.NewTicker(100 * time.Millisecond)
		for {
			<-ticker.C
			updateParagraph(tickerCount)
			ui.Render(p)
			tickerCount++
		}
	}()

	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			ui.Close()
			os.Exit(0)
			return
		}
	}
}
