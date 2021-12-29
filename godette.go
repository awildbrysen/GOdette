package main

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type Context struct {
	win           *pixelgl.Window
	keyListeners  []KeyListener
	CurrentBuffer *TextBuffer
	Buffers       []*TextBuffer
	Cursor        *Cursor
}

type KeyListener struct {
	name string
	key  pixelgl.Button
	mod  pixelgl.Button // -1 for no modifier
	exec func(Context)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "GOdette",
		Bounds: pixel.R(0, 0, 1280, 720),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	buf := NewTextBuffer()

	fps := time.Tick(time.Second / 120)

	ctx := Context{}
	ctx.keyListeners = append(ctx.keyListeners, KeyListener{name: "add newline", key: pixelgl.KeyEnter, mod: -1, exec: func(ctx Context) {
		ctx.CurrentBuffer.WriteRune('\n')
	}})
	ctx.keyListeners = append(ctx.keyListeners, KeyListener{name: "remove character", key: pixelgl.KeyBackspace, mod: -1, exec: func(ctx Context) {
		if len(ctx.CurrentBuffer.Buf) > 0 {
			ctx.CurrentBuffer.Buf = ctx.CurrentBuffer.Buf[:len(ctx.CurrentBuffer.Buf)-1]
		}
	}})
	ctx.win = win
	ctx.CurrentBuffer = &buf
	ctx.Buffers = append(ctx.Buffers, &buf)

	// primitive file reading
	ctx.keyListeners = append(ctx.keyListeners, KeyListener{name: "open file", key: pixelgl.KeyO, mod: pixelgl.KeyLeftControl, exec: func(ctx Context) {
		println("Opening file")
		// TODO(brysen): show current directory in new buffer and make it the active buffer
	}})

	c := NewCursor(ctx)
	ctx.Cursor = c

	ctx.keyListeners = append(ctx.keyListeners, KeyListener{name: "move cursor up", key: pixelgl.KeyUp, mod: -1, exec: func(ctx Context) {
		ctx.Cursor.Position = ctx.Cursor.Position.Add(pixel.V(0, ctx.Cursor.Height))
	}})
	ctx.keyListeners = append(ctx.keyListeners, KeyListener{name: "move cursor down", key: pixelgl.KeyDown, mod: -1, exec: func(ctx Context) {
		ctx.Cursor.Position = ctx.Cursor.Position.Sub(pixel.V(0, ctx.Cursor.Height))
	}})
	ctx.keyListeners = append(ctx.keyListeners, KeyListener{name: "move cursor left", key: pixelgl.KeyLeft, mod: -1, exec: func(ctx Context) {
		ctx.Cursor.Position = ctx.Cursor.Position.Sub(pixel.V(ctx.Cursor.Width, 0))
	}})
	ctx.keyListeners = append(ctx.keyListeners, KeyListener{name: "move cursor right", key: pixelgl.KeyRight, mod: -1, exec: func(ctx Context) {
		ctx.Cursor.Position = ctx.Cursor.Position.Add(pixel.V(ctx.Cursor.Width, 0))
	}})

	for !win.Closed() {
		ctx.CurrentBuffer.WriteString(win.Typed())

		handleKeyListeners(ctx)

		win.Clear(colornames.Lightcoral)

		c.Update()
		c.Draw(ctx)

		ctx.CurrentBuffer.Draw(ctx)
		win.Update()

		<-fps
	}
}

func handleKeyListeners(ctx Context) {
	for _, kl := range ctx.keyListeners {
		if kl.mod == -1 {
			if ctx.win.JustPressed(kl.key) || ctx.win.Repeated(kl.key) {
				kl.exec(ctx)
			}
		} else {
			if ctx.win.Pressed(kl.mod) && ctx.win.JustPressed(kl.key) {
				kl.exec(ctx)
			}
		}
	}
}

func main() {
	pixelgl.Run(run)
}
