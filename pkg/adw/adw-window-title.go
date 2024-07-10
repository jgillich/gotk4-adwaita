// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// #include <stdlib.h>
// #include <adwaita.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeWindowTitle = coreglib.Type(C.adw_window_title_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeWindowTitle, F: marshalWindowTitle},
	})
}

// WindowTitleOverrides contains methods that are overridable.
type WindowTitleOverrides struct {
}

func defaultWindowTitleOverrides(v *WindowTitle) WindowTitleOverrides {
	return WindowTitleOverrides{}
}

// WindowTitle: helper widget for setting a window's title and subtitle.
//
// <picture> <source srcset="window-title-dark.png"
// media="(prefers-color-scheme: dark)"> <img src="window-title.png"
// alt="window-title"> </picture>
//
// AdwWindowTitle shows a title and subtitle. It's intended to be used as the
// title child of gtk.HeaderBar or headerbar.
//
// # CSS nodes
//
// AdwWindowTitle has a single CSS node with name windowtitle.
type WindowTitle struct {
	_ [0]func() // equal guard
	gtk.Widget
}

var (
	_ gtk.Widgetter = (*WindowTitle)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*WindowTitle, *WindowTitleClass, WindowTitleOverrides](
		GTypeWindowTitle,
		initWindowTitleClass,
		wrapWindowTitle,
		defaultWindowTitleOverrides,
	)
}

func initWindowTitleClass(gclass unsafe.Pointer, overrides WindowTitleOverrides, classInitFunc func(*WindowTitleClass)) {
	if classInitFunc != nil {
		class := (*WindowTitleClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapWindowTitle(obj *coreglib.Object) *WindowTitle {
	return &WindowTitle{
		Widget: gtk.Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			Accessible: gtk.Accessible{
				Object: obj,
			},
			Buildable: gtk.Buildable{
				Object: obj,
			},
			ConstraintTarget: gtk.ConstraintTarget{
				Object: obj,
			},
		},
	}
}

func marshalWindowTitle(p uintptr) (interface{}, error) {
	return wrapWindowTitle(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewWindowTitle creates a new AdwWindowTitle.
//
// The function takes the following parameters:
//
//   - title: title.
//   - subtitle: subtitle.
//
// The function returns the following values:
//
//   - windowTitle: newly created AdwWindowTitle.
func NewWindowTitle(title, subtitle string) *WindowTitle {
	var _arg1 *C.char      // out
	var _arg2 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(subtitle)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.adw_window_title_new(_arg1, _arg2)
	runtime.KeepAlive(title)
	runtime.KeepAlive(subtitle)

	var _windowTitle *WindowTitle // out

	_windowTitle = wrapWindowTitle(coreglib.Take(unsafe.Pointer(_cret)))

	return _windowTitle
}

// Subtitle gets the subtitle of self.
//
// The function returns the following values:
//
//   - utf8: subtitle.
func (self *WindowTitle) Subtitle() string {
	var _arg0 *C.AdwWindowTitle // out
	var _cret *C.char           // in

	_arg0 = (*C.AdwWindowTitle)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_window_title_get_subtitle(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Title gets the title of self.
//
// The function returns the following values:
//
//   - utf8: title.
func (self *WindowTitle) Title() string {
	var _arg0 *C.AdwWindowTitle // out
	var _cret *C.char           // in

	_arg0 = (*C.AdwWindowTitle)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_window_title_get_title(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// SetSubtitle sets the subtitle of self.
//
// The subtitle should give the user additional details.
//
// The function takes the following parameters:
//
//   - subtitle: subtitle.
func (self *WindowTitle) SetSubtitle(subtitle string) {
	var _arg0 *C.AdwWindowTitle // out
	var _arg1 *C.char           // out

	_arg0 = (*C.AdwWindowTitle)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(subtitle)))
	defer C.free(unsafe.Pointer(_arg1))

	C.adw_window_title_set_subtitle(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(subtitle)
}

// SetTitle sets the title of self.
//
// The title typically identifies the current view or content item, and
// generally does not use the application name.
//
// The function takes the following parameters:
//
//   - title: title.
func (self *WindowTitle) SetTitle(title string) {
	var _arg0 *C.AdwWindowTitle // out
	var _arg1 *C.char           // out

	_arg0 = (*C.AdwWindowTitle)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg1))

	C.adw_window_title_set_title(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(title)
}

// WindowTitleClass: instance of this type is always passed by reference.
type WindowTitleClass struct {
	*windowTitleClass
}

// windowTitleClass is the struct that's finalized.
type windowTitleClass struct {
	native *C.AdwWindowTitleClass
}

func (w *WindowTitleClass) ParentClass() *gtk.WidgetClass {
	valptr := &w.native.parent_class
	var _v *gtk.WidgetClass // out
	_v = (*gtk.WidgetClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
