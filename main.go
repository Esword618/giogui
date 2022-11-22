/*
 * @Author: Esword
 * @Description:
 * @FileName:  main
 * @Version: 1.0.0
 * @Date: 2022-06-19 18:17
 */

package main

import (
	"flag"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font/opentype"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget/material"

	"github.com/Esword618/giogui/fonts"
	"github.com/Esword618/giogui/mylayout/pages"
	"github.com/Esword618/giogui/mylayout/pages/appbar"
	"github.com/Esword618/giogui/mylayout/pages/login"
	"github.com/Esword618/giogui/mylayout/pages/navdrawer"
)

func main() {
	flag.Parse()
	go func() {
		w := app.NewWindow()
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func loop(w *app.Window) error {
	font, err := opentype.Parse(fonts.FzstTTF)
	if err != nil {
		panic(err)
	}
	myfonts := []text.FontFace{
		{Face: font},
	}
	th := material.NewTheme(myfonts)
	var ops op.Ops

	router := pages.NewRouter()
	router.Register(0, navdrawer.New(&router))
	router.Register(1, appbar.New(&router))
	router.Register(2, login.New(&router))
	for {
		select {
		case e := <-w.Events():
			switch e := e.(type) {
			case system.DestroyEvent:
				return e.Err
			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)
				router.Layout(gtx, th)
				e.Frame(gtx.Ops)
			}
		}
	}
}
