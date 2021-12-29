package main

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

type Cursor struct {
	Width    float64
	Height   float64
	Position pixel.Vec

	imd *imdraw.IMDraw
}

func NewCursor(ctx Context) *Cursor {
	w := ctx.CurrentBuffer.txt.Atlas().Glyph(' ').Advance
	h := ctx.CurrentBuffer.txt.Atlas().LineHeight()
	c := Cursor{Width: w, Height: h, Position: pixel.V(50, 500)}
	c.imd = imdraw.New(nil)
	return &c
}

func (c *Cursor) Update() {
	fmt.Printf("Cursor position %+v\n", c.Position)
	c.imd.Clear()
	c.imd.Color = pixel.RGB(1, 0, 0)
	c.imd.Push(c.Position)
	c.imd.Push(c.Position.Add(pixel.V(c.Width, 0)))
	c.imd.Push(c.Position.Add(pixel.V(0, c.Height)))
	c.imd.Push(c.Position.Add(pixel.V(c.Width, c.Height)))
	c.imd.Rectangle(0)
}

func (c *Cursor) Draw(ctx Context) {
	c.imd.Draw(ctx.win)
}
