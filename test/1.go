/*
 * @Author: Esword
 * @Description:
 * @FileName:  1
 * @Version: 1.0.0
 * @Date: 2022-06-19 20:35
 */

package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font/opentype"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"

	"eliasnaur.com/font/noto/sans/jp/regular"
)

func main() {
	go func() {
		size := app.Size(unit.Dp(600), unit.Dp(480))
		w := app.NewWindow(size)
		err := run(w)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(w *app.Window) error {
	font, err := opentype.Parse(regular.TTF)
	if err != nil {
		panic(err)
	}
	fonts := []text.FontFace{
		{Face: font},
	}

	th := material.NewTheme(fonts)
	_ = th
	var ops op.Ops

	for {
		select {
		case e := <-w.Events():
			switch e := e.(type) {
			case system.DestroyEvent:
				return e.Err
			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)
				layout.Flex{
					Axis:    layout.Vertical,
					Spacing: layout.SpaceEvenly,
				}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return material.H4(th, "I can eat glass, it doesn't hurt me.").Layout(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return material.H4(th, "我能吞下玻璃而不伤身体.").Layout(gtx)
					}),
				)
				e.Frame(gtx.Ops)
			}
		}
	}
}
