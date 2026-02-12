# 🏗️ uvcasso

Layout splitting for [charmbracelet/ultraviolet]
based on [Cassowary constraint-solving algorithm].

## TL;DR

It allows you to dynamically split terminal screen into multiple rectangles
based on the given constraints.

For example:

```go
func display(s *uv.TerminalScreen) {
	screen.Clear(s)

	var top, bottom uv.Rectangle

	uvcasso.New(
		uvcasso.Fill(1),
		uvcasso.Len(1),
		uvcasso.Len(3),
	).
		Vertical().
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
	   TTTTTTTTTTTTTTTT
	   TTTTTTTTTTTTTTTT
	   TTTTTTTTTTTTTTTT
	   TTTTTTTTTTTTTTTT

	   BBBBBBBBBBBBBBBB
	   BBBBBBBBBBBBBBBB
	   BBBBBBBBBBBBBBBB
	*/

	s.Render()
	s.Flush()
}
```

## Acknowledgements

This code is roughly 1:1 translation of how it's
implemented in the [ratatui], including the tests.

[charmbracelet/ultraviolet]: https://github.com/charmbracelet/ultraviolet
[Cassowary constraint-solving algorithm]: https://en.wikipedia.org/wiki/Cassowary_(software)
[ratatui]: https://ratatui.rs/
