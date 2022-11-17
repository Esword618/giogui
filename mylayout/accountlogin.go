/*
 * @Author: Esword
 * @Description:
 * @FileName:  login
 * @Version: 1.0.0
 * @Date: 2022-06-19 21:35
 */

package mylayout

import (
	"fmt"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
	"gioui.org/x/notify"

	"guigio/fonts"
)

type Login struct {
	inputAlignment                                               layout.Alignment
	inputAlignmentEnum                                           widget.Enum
	username, password                                           component.TextField
	nameInput, addressInput, priceInput, tweetInput, numberInput component.TextField
	widget.List
}

var (
	th       = material.NewTheme(fonts.InitFont())
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

type (
	// C quick alias for Context.
	C = layout.Context
	// D quick alias for Dimensions.
	D = layout.Dimensions
)

func (p *Login) Frame(gtx C) D {
	if notifyBtn.Clicked() {
		msg := "This is a notification send from gio."
		if txt := editor.Text(); txt != "" {
			msg = txt
		}
		fmt.Println(p.username.Text())
		fmt.Println(p.password.Text())
		go notifier.CreateNotification("Esword", msg)
	}

	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(gtx,
		// layout.Rigid(func(gtx C) D {
		// 	return alo.DefaultInset.Layout(gtx, material.Body2(th, "input username").Layout)
		// }),
		layout.Rigid(func(gtx C) D {
			gtx.Constraints.Max.X = gtx.Dp(unit.Dp(300))
			p.username.Alignment = p.inputAlignment
			return p.username.Layout(gtx, th, "UserName")
		}),
		// layout.Rigid(func(gtx C) D {
		// 	return alo.DefaultInset.Layout(gtx, material.Body2(th, "input password").Layout)
		// }),
		layout.Rigid(func(gtx C) D {
			gtx.Constraints.Max.X = gtx.Dp(unit.Dp(300))

			p.password.Alignment = p.inputAlignment
			return p.password.Layout(gtx, th, "Password")
		}),
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
	)
}
