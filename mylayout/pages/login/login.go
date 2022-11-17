/*
 * @Author: Esword
 * @Description:
 * @FileName:  qr
 * @Version: 1.0.0
 * @Date: 2022-06-23 12:46
 */

package login

import (
	"context"
	"os"
	"os/signal"
	"sync"

	"gioui.org/app"
	"gioui.org/font/opentype"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	"guigio/fonts"
	"guigio/mylayout/icon"
	page "guigio/mylayout/pages"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

// Page holds the state for a page demonstrating the features of
// the Menu component.
type Page struct {
	Image    widget.Image
	LoginBtn widget.Clickable
	// redButton, greenButton, blueButton       widget.Clickable
	// balanceButton, accountButton, cartButton widget.Clickable
	// leftFillColor                            color.NRGBA
	// leftContextArea                          component.ContextArea
	// leftMenu, rightMenu                      component.MenuState
	// menuInit                                 bool
	menuDemoList       layout.List
	menuDemoListStates []component.ContextArea
	ShowImg            bool
	widget.List

	*page.Router
}

// New constructs a Page with the provided router.
func New(router *page.Router) *Page {
	return &Page{
		Router: router,
	}
}

var _ page.Page = &Page{}

func (p *Page) Actions() []component.AppBarAction {
	return []component.AppBarAction{}
}

func (p *Page) Overflow() []component.OverflowAction {
	return []component.OverflowAction{}
}

func (p *Page) NavItem() component.NavItem {
	return component.NavItem{
		Name: "Login",
		Icon: icon.RestaurantMenuIcon,
	}
}

func (p *Page) Layout(gtx C, th *material.Theme) D {
	p.List.Axis = layout.Vertical
	return material.List(th, &p.List).Layout(gtx, 1, func(gtx C, _ int) D {

		if p.LoginBtn.Clicked() {
			p.ShowImg = !p.ShowImg
			ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
			defer stop()

			go func() {
				a := NewApplication(ctx)

				log := NewLog()
				log.Printf("[Application Started]")
				letters := NewLetters(log)

				a.NewWindow("Log", log)
				a.NewWindow("Letters", letters)

				a.Wait()

				// os.Exit(0)
			}()

			app.Main()
		}
		return layout.Flex{
			Axis: layout.Vertical,
		}.Layout(gtx,
			layout.Rigid(func(gtx C) D {
				gtx.Constraints.Max.X = gtx.Dp(unit.Dp(300))
				return material.Button(th, &p.LoginBtn, "Login").Layout(gtx)
			}),
		)
	})
}

// Application keeps track of all the windows and global state.
type Application struct {
	// Context is used to broadcast application shutdown.
	Context context.Context
	// Shutdown shuts down all windows.
	Shutdown func()
	// active keeps track the open windows, such that application
	// can shut down, when all of them are closed.
	active sync.WaitGroup
}

func NewApplication(ctx context.Context) *Application {
	ctx, cancel := context.WithCancel(ctx)
	return &Application{
		Context:  ctx,
		Shutdown: cancel,
	}
}

// Wait waits for all windows to close.
func (a *Application) Wait() {
	a.active.Wait()
}

// NewWindow creates a new tracked window.
func (a *Application) NewWindow(title string, view View, opts ...app.Option) {
	opts = append(opts, app.Title(title))
	w := &Window{
		App:    a,
		Window: app.NewWindow(opts...),
	}
	a.active.Add(1)
	go func() {
		defer a.active.Done()
		view.Run(w)
	}()
}

// Window holds window state.
type Window struct {
	App *Application
	*app.Window
}

// View describes .
type View interface {
	// Run handles the window event loop.
	Run(w *Window) error
}

// WidgetView allows to use layout.Widget as a view.
type WidgetView func(gtx layout.Context, th *material.Theme) layout.Dimensions

// Run displays the widget with default handling.
func (view WidgetView) Run(w *Window) error {
	font, err := opentype.Parse(fonts.TTF)
	if err != nil {
		panic(err)
	}
	myfonts := []text.FontFace{
		{Face: font},
	}
	var ops op.Ops
	th := material.NewTheme(myfonts)

	//
	applicationClose := w.App.Context.Done()

	for {
		select {
		case <-applicationClose:
			return nil
		case e := <-w.Events():
			switch e := e.(type) {
			case system.DestroyEvent:
				return e.Err
			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)
				view(gtx, th)
				e.Frame(gtx.Ops)
			}
		}
	}
}
