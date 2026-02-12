//go:build example

package main

import (
	"log"

	uv "github.com/charmbracelet/ultraviolet"
	"github.com/charmbracelet/ultraviolet/screen"
	"github.com/metafates/uvcasso"
)

func main() {
	console, err := uv.ControllingConsole()
	if err != nil {
		log.Fatal(err)
	}

	t := uv.NewTerminal(console, uv.DefaultOptions())

	if err := run(t); err != nil {
		log.Fatal(err)
	}
}

func display(s *uv.TerminalScreen) {
	screen.Clear(s)

	var top, bottom uv.Rectangle

	uvcasso.Vertical(
		uvcasso.Fill(1),
		uvcasso.Len(1),
		uvcasso.Len(3),
	).
		Split(s.Bounds()).
		Assign(&top, nil, &bottom)

	// ^^^^^^^^^^^^^^^^^^^^^^^^^^^
	// The interesting part

	screen.FillArea(s, &uv.Cell{
		Content: "T",
		Width:   1,
	}, top)

	screen.FillArea(s, &uv.Cell{
		Content: "B",
		Width:   1,
	}, bottom)

	// Will fill the screen like this:
	// 1. Top part takes all available space
	// 2. Middle part takes exactly 1 line which we leave empty
	// 3. Bottom part takes exactly 3 lines

	/*
	   TTTTTTTTTTTTTTTT

	   BBBBBBBBBBBBBBBB
	   BBBBBBBBBBBBBBBB
	   BBBBBBBBBBBBBBBB
	*/

	s.Render()
	s.Flush()
}

func run(t *uv.Terminal) error {
	scr := t.Screen()

	scr.EnterAltScreen()

	if err := t.Start(); err != nil {
		return err
	}

	defer t.Stop()

	display(scr)

	defer display(scr)

	var physicalWidth, physicalHeight int
	for ev := range t.Events() {
		switch ev := ev.(type) {
		case uv.WindowSizeEvent:
			physicalWidth = ev.Width
			physicalHeight = ev.Height

			scr.Resize(physicalWidth, physicalHeight)

			display(scr)

		case uv.KeyPressEvent:
			return nil
		}
	}

	return nil
}
