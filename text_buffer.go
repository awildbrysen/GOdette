package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
)

type TextBuffer struct {
	Buf      []byte
	txt      *text.Text
	readonly bool
}

func NewTextBuffer() TextBuffer {
	buf := TextBuffer{}
	buf.txt = text.New(pixel.V(50, 500), text.Atlas7x13)
	buf.readonly = false
	return buf
}

func (tb *TextBuffer) Draw(win *pixelgl.Window) {
	tb.txt.Clear()
	tb.txt.WriteString(string(tb.Buf))
	tb.txt.Draw(win, pixel.IM)
}

func (tb *TextBuffer) WriteString(s string) {
	if !tb.readonly {
		println("buffer is readonly") // TODO(brysen): replace with in editor error messaging
		return
	}
	tb.Buf = append(tb.Buf, s...)
}

func (tb *TextBuffer) WriteRune(r rune) {
	if !tb.readonly {
		println("buffer is readonly") // TODO(brysen): replace with in editor error messaging
		return
	}
	tb.Buf = append(tb.Buf, byte(r))
}
