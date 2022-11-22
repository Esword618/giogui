package navdrawer

import (
	"fmt"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
	"gioui.org/x/notify"

	alo "github.com/Esword618/giogui/mylayout/component/applayout"
	"github.com/Esword618/giogui/mylayout/icon"
	page "github.com/Esword618/giogui/mylayout/pages"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

type Input struct {
	inputAlignment                                               layout.Alignment
	inputAlignmentEnum                                           widget.Enum
	username, password                                           component.TextField
	nameInput, addressInput, priceInput, tweetInput, numberInput component.TextField
	widget.List
}

// Page holds the state for a page demonstrating the features of
// the NavDrawer component.
type Page struct {
	heartBtn, plusBtn, contextBtn          widget.Clickable
	exampleOverflowState, red, green, blue widget.Clickable
	nonModalDrawer                         widget.Bool
	bottomBar, customNavIcon               widget.Bool
	favorited                              bool
	widget.List
	Input
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
		Name: "Nav Drawer Features",
		Icon: icon.SettingsIcon,
	}
}

var (
	// th       = material.NewTheme(fonts.InitFont())
	notifier = func() notify.Notifier {
		n, err := notify.NewNotifier()
		if err != nil {
			panic(fmt.Errorf("init notification manager: %w", err))
		}
		return n
	}()
	editor    component.TextField
	notifyBtn widget.Clickable
)

// -------------------------

func (p *Page) Layout(gtx C, th *material.Theme) D {
	// var pp = Login{}
	p.List.Axis = layout.Vertical
	return material.List(th, &p.Input.List).Layout(gtx, 1, func(gtx C, _ int) D {
		if notifyBtn.Clicked() {
			msg := "This is a notification send from gio."
			if txt := editor.Text(); txt != "" {
				msg = txt
			}
			fmt.Println(p.Input.username.Text())
			fmt.Println(p.Input.password.Text())
			go notifier.CreateNotification("Esword", msg)
		}

		return layout.Flex{
			Axis: layout.Vertical,
		}.Layout(gtx,
			layout.Rigid(func(gtx C) D {
				return alo.DefaultInset.Layout(gtx, material.Body2(th, "please input username").Layout)
			}),
			layout.Rigid(func(gtx C) D {
				gtx.Constraints.Max.X = gtx.Dp(unit.Dp(300))
				p.Input.username.Alignment = p.Input.inputAlignment
				return p.Input.username.Layout(gtx, th, "UserName")
			}),
			layout.Rigid(func(gtx C) D {
				return alo.DefaultInset.Layout(gtx, material.Body2(th, "please input password").Layout)
			}),
			layout.Rigid(func(gtx C) D {
				gtx.Constraints.Max.X = gtx.Dp(unit.Dp(300))

				p.Input.password.Alignment = p.Input.inputAlignment
				return p.Input.password.Layout(gtx, th, "Password")
			}),
			// -----------------------------------

			layout.Rigid(func(gtx C) D {
				// info := fmt.Sprintf("um:%S,pwd:%S",p.username,p.password)
				gtx.Constraints.Max.X = gtx.Dp(unit.Dp(300))
				return material.Button(th, &notifyBtn, "notify").Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return material.H4(th, "I can eat glass, it doesn't hurt me.").Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return material.H4(th, "我能吞下玻璃而不伤身体.").Layout(gtx)
			}),

			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					gtx.Constraints.Max.X = gtx.Dp(unit.Dp(300))
					return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
						layout.Rigid(func(gtx C) D {
							return editor.Layout(gtx, th, "enter a notification message")
						}),
						layout.Rigid(func(gtx C) D {
							return layout.Spacer{Height: unit.Dp(10)}.Layout(gtx)
						}),
						layout.Rigid(func(gtx C) D {
							return material.Button(th, &notifyBtn, "notify").Layout(gtx)
						}),
					)
				})
			}),
		)
		// 		return layout.Flex{
		// 			Alignment: layout.Middle,
		// 			Axis:      layout.Vertical,
		// 		}.Layout(gtx,
		// 			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		// 				return alo.DefaultInset.Layout(gtx, material.Body1(th, `The nav drawer widget provides a consistent interface element for navigation.
		//
		// The controls below allow you to see the various features available in our Navigation Drawer implementation.`).Layout)
		// 			}),
		// 			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		// 				return alo.DetailRow{}.Layout(gtx,
		// 					material.Body1(th, "Use non-modal drawer").Layout,
		// 					func(gtx C) D {
		// 						if p.nonModalDrawer.Changed() {
		// 							p.Router.NonModalDrawer = p.nonModalDrawer.Value
		// 							if p.nonModalDrawer.Value {
		// 								p.Router.NavAnim.Appear(gtx.Now)
		// 							} else {
		// 								p.Router.NavAnim.Disappear(gtx.Now)
		// 							}
		// 						}
		// 						return material.Switch(th, &p.nonModalDrawer, "Use Non-Modal Navigation Drawer").Layout(gtx)
		// 					})
		// 			}),
		// 			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		// 				return alo.DetailRow{}.Layout(gtx,
		// 					material.Body1(th, "Drag to Close").Layout,
		// 					material.Body2(th, "You can close the modal nav drawer by dragging it to the left.").Layout)
		// 			}),
		// 			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		// 				return alo.DetailRow{}.Layout(gtx,
		// 					material.Body1(th, "Touch Scrim to Close").Layout,
		// 					material.Body2(th, "You can close the modal nav drawer touching anywhere in the translucent scrim to the right.").Layout)
		// 			}),
		// 			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		// 				return alo.DetailRow{}.Layout(gtx,
		// 					material.Body1(th, "Bottom content anchoring").Layout,
		// 					material.Body2(th, "If you toggle support for the bottom app bar in the App Bar settings, nav drawer content will anchor to the bottom of the drawer area instead of the top.").Layout)
		// 			}),
		// 		)
	})
}
