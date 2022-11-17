/*
 * @Author: Esword
 * @Description:
 * @FileName:  layout
 * @Version: 1.0.0
 * @Date: 2022-06-19 21:37
 */

package global

import (
	"fmt"

	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
	"gioui.org/x/notify"

	"guigio/fonts"
)

type (
	// C quick alias for Context.
	C = layout.Context
	// D quick alias for Dimensions.
	D = layout.Dimensions
)

type Page struct {
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
