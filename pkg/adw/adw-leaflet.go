// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// #include <stdlib.h>
// #include <adwaita.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeLeafletTransitionType = coreglib.Type(C.adw_leaflet_transition_type_get_type())
	GTypeLeaflet               = coreglib.Type(C.adw_leaflet_get_type())
	GTypeLeafletPage           = coreglib.Type(C.adw_leaflet_page_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeLeafletTransitionType, F: marshalLeafletTransitionType},
		coreglib.TypeMarshaler{T: GTypeLeaflet, F: marshalLeaflet},
		coreglib.TypeMarshaler{T: GTypeLeafletPage, F: marshalLeafletPage},
	})
}

// LeafletTransitionType describes the possible transitions in a leaflet widget.
//
// New values may be added to this enumeration over time.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwleaflet).
type LeafletTransitionType C.gint

const (
	// LeafletTransitionTypeOver: cover the old page or uncover the new page,
	// sliding from or towards the end according to orientation, text direction
	// and children order.
	LeafletTransitionTypeOver LeafletTransitionType = iota
	// LeafletTransitionTypeUnder: uncover the new page or cover the old page,
	// sliding from or towards the start according to orientation, text
	// direction and children order.
	LeafletTransitionTypeUnder
	// LeafletTransitionTypeSlide: slide from left, right, up or down according
	// to the orientation, text direction and the children order.
	LeafletTransitionTypeSlide
)

func marshalLeafletTransitionType(p uintptr) (interface{}, error) {
	return LeafletTransitionType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for LeafletTransitionType.
func (l LeafletTransitionType) String() string {
	switch l {
	case LeafletTransitionTypeOver:
		return "Over"
	case LeafletTransitionTypeUnder:
		return "Under"
	case LeafletTransitionTypeSlide:
		return "Slide"
	default:
		return fmt.Sprintf("LeafletTransitionType(%d)", l)
	}
}

// LeafletOverrides contains methods that are overridable.
type LeafletOverrides struct {
}

func defaultLeafletOverrides(v *Leaflet) LeafletOverrides {
	return LeafletOverrides{}
}

// Leaflet: adaptive container acting like a box or a stack.
//
// <picture> <source srcset="leaflet-wide-dark.png"
// media="(prefers-color-scheme: dark)"> <img src="leaflet-wide.png"
// alt="leaflet-wide"> </picture> <picture> <source
// srcset="leaflet-narrow-dark.png" media="(prefers-color-scheme: dark)"> <img
// src="leaflet-narrow.png" alt="leaflet-narrow"> </picture>
//
// The AdwLeaflet widget can display its children like a gtk.Box does or like a
// gtk.Stack does, adapting to size changes by switching between the two modes.
//
// When there is enough space the children are displayed side by side,
// otherwise only one is displayed and the leaflet is said to be “folded”.
// The threshold is dictated by the preferred minimum sizes of the children.
// When a leaflet is folded, the children can be navigated using swipe gestures.
//
// The “over” and “under” transition types stack the children one on top of
// the other, while the “slide” transition puts the children side by side.
// While navigating to a child on the side or below can be performed by swiping
// the current child away, navigating to an upper child requires dragging it
// from the edge where it resides. This doesn't affect non-dragging swipes.
//
// # CSS nodes
//
// AdwLeaflet has a single CSS node with name leaflet. The node will get the
// style classes .folded when it is folded, .unfolded when it's not, or none if
// it hasn't computed its fold yet.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwleaflet).
type Leaflet struct {
	_ [0]func() // equal guard
	gtk.Widget

	*coreglib.Object
	Swipeable
	gtk.Orientable
}

var (
	_ gtk.Widgetter     = (*Leaflet)(nil)
	_ coreglib.Objector = (*Leaflet)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*Leaflet, *LeafletClass, LeafletOverrides](
		GTypeLeaflet,
		initLeafletClass,
		wrapLeaflet,
		defaultLeafletOverrides,
	)
}

func initLeafletClass(gclass unsafe.Pointer, overrides LeafletOverrides, classInitFunc func(*LeafletClass)) {
	if classInitFunc != nil {
		class := (*LeafletClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapLeaflet(obj *coreglib.Object) *Leaflet {
	return &Leaflet{
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
		Object: obj,
		Swipeable: Swipeable{
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
		},
		Orientable: gtk.Orientable{
			Object: obj,
		},
	}
}

func marshalLeaflet(p uintptr) (interface{}, error) {
	return wrapLeaflet(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewLeaflet creates a new AdwLeaflet.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwleaflet).
//
// The function returns the following values:
//
//   - leaflet: new created AdwLeaflet.
func NewLeaflet() *Leaflet {
	var _cret *C.GtkWidget // in

	_cret = C.adw_leaflet_new()

	var _leaflet *Leaflet // out

	_leaflet = wrapLeaflet(coreglib.Take(unsafe.Pointer(_cret)))

	return _leaflet
}

// Append adds a child to self.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwleaflet).
//
// The function takes the following parameters:
//
//   - child: widget to add.
//
// The function returns the following values:
//
//   - leafletPage: leafletpage for child.
func (self *Leaflet) Append(child gtk.Widgetter) *LeafletPage {
	var _arg0 *C.AdwLeaflet     // out
	var _arg1 *C.GtkWidget      // out
	var _cret *C.AdwLeafletPage // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	_cret = C.adw_leaflet_append(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)

	var _leafletPage *LeafletPage // out

	_leafletPage = wrapLeafletPage(coreglib.Take(unsafe.Pointer(_cret)))

	return _leafletPage
}

// AdjacentChild finds the previous or next navigatable child.
//
// This will be the same child leaflet.Navigate or swipe gestures will navigate
// to.
//
// If there's no child to navigate to, NULL will be returned instead.
//
// See leafletpage:navigatable.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwleaflet).
//
// The function takes the following parameters:
//
//   - direction: direction.
//
// The function returns the following values:
//
//   - widget (optional) previous or next child.
func (self *Leaflet) AdjacentChild(direction NavigationDirection) gtk.Widgetter {
	var _arg0 *C.AdwLeaflet            // out
	var _arg1 C.AdwNavigationDirection // out
	var _cret *C.GtkWidget             // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.AdwNavigationDirection(direction)

	_cret = C.adw_leaflet_get_adjacent_child(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(direction)

	var _widget gtk.Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(gtk.Widgetter)
				return ok
			})
			rv, ok := casted.(gtk.Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// CanNavigateBack gets whether gestures and shortcuts for navigating backward
// are enabled.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwleaflet).
//
// The function returns the following values:
//
//   - ok: whether gestures and shortcuts are enabled.
func (self *Leaflet) CanNavigateBack() bool {
	var _arg0 *C.AdwLeaflet // out
	var _cret C.gboolean    // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_leaflet_get_can_navigate_back(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CanNavigateForward gets whether gestures and shortcuts for navigating forward
// are enabled.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwleaflet).
//
// The function returns the following values:
//
//   - ok: whether gestures and shortcuts are enabled.
func (self *Leaflet) CanNavigateForward() bool {
	var _arg0 *C.AdwLeaflet // out
	var _cret C.gboolean    // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_leaflet_get_can_navigate_forward(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CanUnfold gets whether self can unfold.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwleaflet).
//
// The function returns the following values:
//
//   - ok: whether self can unfold.
func (self *Leaflet) CanUnfold() bool {
	var _arg0 *C.AdwLeaflet // out
	var _cret C.gboolean    // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_leaflet_get_can_unfold(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ChildByName finds the child of self with name.
//
// Returns NULL if there is no child with this name.
//
// See leafletpage:name.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwleaflet).
//
// The function takes the following parameters:
//
//   - name of the child to find.
//
// The function returns the following values:
//
//   - widget (optional): requested child of self.
func (self *Leaflet) ChildByName(name string) gtk.Widgetter {
	var _arg0 *C.AdwLeaflet // out
	var _arg1 *C.char       // out
	var _cret *C.GtkWidget  // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.adw_leaflet_get_child_by_name(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(name)

	var _widget gtk.Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(gtk.Widgetter)
				return ok
			})
			rv, ok := casted.(gtk.Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// ChildTransitionParams gets the child transition spring parameters for self.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwleaflet).
//
// The function returns the following values:
//
//   - springParams: child transition parameters.
func (self *Leaflet) ChildTransitionParams() *SpringParams {
	var _arg0 *C.AdwLeaflet      // out
	var _cret *C.AdwSpringParams // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_leaflet_get_child_transition_params(_arg0)
	runtime.KeepAlive(self)

	var _springParams *SpringParams // out

	_springParams = (*SpringParams)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_springParams)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.adw_spring_params_unref((*C.AdwSpringParams)(intern.C))
		},
	)

	return _springParams
}

// ChildTransitionRunning gets whether a child transition is currently running
// for self.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwleaflet).
//
// The function returns the following values:
//
//   - ok: whether a transition is currently running.
func (self *Leaflet) ChildTransitionRunning() bool {
	var _arg0 *C.AdwLeaflet // out
	var _cret C.gboolean    // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_leaflet_get_child_transition_running(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// FoldThresholdPolicy gets the fold threshold policy for self.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwleaflet).
//
// The function returns the following values:
//
//   - foldThresholdPolicy: fold threshold policy.
func (self *Leaflet) FoldThresholdPolicy() FoldThresholdPolicy {
	var _arg0 *C.AdwLeaflet            // out
	var _cret C.AdwFoldThresholdPolicy // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_leaflet_get_fold_threshold_policy(_arg0)
	runtime.KeepAlive(self)

	var _foldThresholdPolicy FoldThresholdPolicy // out

	_foldThresholdPolicy = FoldThresholdPolicy(_cret)

	return _foldThresholdPolicy
}

// Folded gets whether self is folded.
//
// The leaflet will be folded if the size allocated to it is smaller
// than the sum of the minimum or natural sizes of the children (see
// leaflet:fold-threshold-policy), it will be unfolded otherwise.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwleaflet).
//
// The function returns the following values:
//
//   - ok: whether self is folded.
func (self *Leaflet) Folded() bool {
	var _arg0 *C.AdwLeaflet // out
	var _cret C.gboolean    // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_leaflet_get_folded(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Homogeneous gets whether self is homogeneous.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwleaflet).
//
// The function returns the following values:
//
//   - ok: whether self is homogeneous.
func (self *Leaflet) Homogeneous() bool {
	var _arg0 *C.AdwLeaflet // out
	var _cret C.gboolean    // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_leaflet_get_homogeneous(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ModeTransitionDuration gets the mode transition animation duration for self.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwleaflet).
//
// The function returns the following values:
//
//   - guint: mode transition duration, in milliseconds.
func (self *Leaflet) ModeTransitionDuration() uint {
	var _arg0 *C.AdwLeaflet // out
	var _cret C.guint       // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_leaflet_get_mode_transition_duration(_arg0)
	runtime.KeepAlive(self)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Page returns the leafletpage object for child.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwleaflet).
//
// The function takes the following parameters:
//
//   - child of self.
//
// The function returns the following values:
//
//   - leafletPage: page object for child.
func (self *Leaflet) Page(child gtk.Widgetter) *LeafletPage {
	var _arg0 *C.AdwLeaflet     // out
	var _arg1 *C.GtkWidget      // out
	var _cret *C.AdwLeafletPage // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	_cret = C.adw_leaflet_get_page(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)

	var _leafletPage *LeafletPage // out

	_leafletPage = wrapLeafletPage(coreglib.Take(unsafe.Pointer(_cret)))

	return _leafletPage
}

// Pages returns a gio.ListModel that contains the pages of the leaflet.
//
// This can be used to keep an up-to-date view. The model also implements
// gtk.SelectionModel and can be used to track and change the visible page.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwleaflet).
//
// The function returns the following values:
//
//   - selectionModel: GtkSelectionModel for the leaflet's children.
func (self *Leaflet) Pages() *gtk.SelectionModel {
	var _arg0 *C.AdwLeaflet        // out
	var _cret *C.GtkSelectionModel // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_leaflet_get_pages(_arg0)
	runtime.KeepAlive(self)

	var _selectionModel *gtk.SelectionModel // out

	{
		obj := coreglib.AssumeOwnership(unsafe.Pointer(_cret))
		_selectionModel = &gtk.SelectionModel{
			ListModel: gio.ListModel{
				Object: obj,
			},
		}
	}

	return _selectionModel
}

// TransitionType gets the type of animation used for transitions between modes
// and children.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwleaflet).
//
// The function returns the following values:
//
//   - leafletTransitionType: current transition type of self.
func (self *Leaflet) TransitionType() LeafletTransitionType {
	var _arg0 *C.AdwLeaflet              // out
	var _cret C.AdwLeafletTransitionType // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_leaflet_get_transition_type(_arg0)
	runtime.KeepAlive(self)

	var _leafletTransitionType LeafletTransitionType // out

	_leafletTransitionType = LeafletTransitionType(_cret)

	return _leafletTransitionType
}

// VisibleChild gets the widget currently visible when the leaflet is folded.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwleaflet).
//
// The function returns the following values:
//
//   - widget (optional): visible child.
func (self *Leaflet) VisibleChild() gtk.Widgetter {
	var _arg0 *C.AdwLeaflet // out
	var _cret *C.GtkWidget  // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_leaflet_get_visible_child(_arg0)
	runtime.KeepAlive(self)

	var _widget gtk.Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(gtk.Widgetter)
				return ok
			})
			rv, ok := casted.(gtk.Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// VisibleChildName gets the name of the currently visible child widget.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwleaflet).
//
// The function returns the following values:
//
//   - utf8 (optional): name of the visible child.
func (self *Leaflet) VisibleChildName() string {
	var _arg0 *C.AdwLeaflet // out
	var _cret *C.char       // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_leaflet_get_visible_child_name(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// InsertChildAfter inserts child in the position after sibling in the list of
// children.
//
// If sibling is NULL, inserts child at the first position.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwleaflet).
//
// The function takes the following parameters:
//
//   - child: widget to insert.
//   - sibling (optional) after which to insert child.
//
// The function returns the following values:
//
//   - leafletPage: leafletpage for child.
func (self *Leaflet) InsertChildAfter(child, sibling gtk.Widgetter) *LeafletPage {
	var _arg0 *C.AdwLeaflet     // out
	var _arg1 *C.GtkWidget      // out
	var _arg2 *C.GtkWidget      // out
	var _cret *C.AdwLeafletPage // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	if sibling != nil {
		_arg2 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(sibling).Native()))
	}

	_cret = C.adw_leaflet_insert_child_after(_arg0, _arg1, _arg2)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
	runtime.KeepAlive(sibling)

	var _leafletPage *LeafletPage // out

	_leafletPage = wrapLeafletPage(coreglib.Take(unsafe.Pointer(_cret)))

	return _leafletPage
}

// Navigate navigates to the previous or next child.
//
// The child must have the leafletpage:navigatable property set to TRUE,
// otherwise it will be skipped.
//
// This will be the same child as returned by leaflet.GetAdjacentChild or
// navigated to via swipe gestures.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwleaflet).
//
// The function takes the following parameters:
//
//   - direction: direction.
//
// The function returns the following values:
//
//   - ok: whether the visible child was changed.
func (self *Leaflet) Navigate(direction NavigationDirection) bool {
	var _arg0 *C.AdwLeaflet            // out
	var _arg1 C.AdwNavigationDirection // out
	var _cret C.gboolean               // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.AdwNavigationDirection(direction)

	_cret = C.adw_leaflet_navigate(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(direction)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Prepend inserts child at the first position in self.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwleaflet).
//
// The function takes the following parameters:
//
//   - child: widget to prepend.
//
// The function returns the following values:
//
//   - leafletPage: leafletpage for child.
func (self *Leaflet) Prepend(child gtk.Widgetter) *LeafletPage {
	var _arg0 *C.AdwLeaflet     // out
	var _arg1 *C.GtkWidget      // out
	var _cret *C.AdwLeafletPage // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	_cret = C.adw_leaflet_prepend(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)

	var _leafletPage *LeafletPage // out

	_leafletPage = wrapLeafletPage(coreglib.Take(unsafe.Pointer(_cret)))

	return _leafletPage
}

// Remove removes a child widget from self.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwleaflet).
//
// The function takes the following parameters:
//
//   - child to remove.
func (self *Leaflet) Remove(child gtk.Widgetter) {
	var _arg0 *C.AdwLeaflet // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	C.adw_leaflet_remove(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
}

// ReorderChildAfter moves child to the position after sibling in the list of
// children.
//
// If sibling is NULL, moves child to the first position.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwleaflet).
//
// The function takes the following parameters:
//
//   - child: widget to move, must be a child of self.
//   - sibling (optional) to move child after.
func (self *Leaflet) ReorderChildAfter(child, sibling gtk.Widgetter) {
	var _arg0 *C.AdwLeaflet // out
	var _arg1 *C.GtkWidget  // out
	var _arg2 *C.GtkWidget  // out

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	if sibling != nil {
		_arg2 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(sibling).Native()))
	}

	C.adw_leaflet_reorder_child_after(_arg0, _arg1, _arg2)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
	runtime.KeepAlive(sibling)
}

// SetCanNavigateBack sets whether gestures and shortcuts for navigating
// backward are enabled.
//
// The supported gestures are:
//
// - One-finger swipe on touchscreens
//
// - Horizontal scrolling on touchpads (usually two-finger swipe)
//
// - Back/forward mouse buttons
//
// The keyboard back/forward keys are also supported, as well as the
// <kbd>Alt</kbd>+<kbd>←</kbd> shortcut for horizontal orientation, or
// <kbd>Alt</kbd>+<kbd>↑</kbd> for vertical orientation.
//
// If the orientation is horizontal, for right-to-left locales, gestures and
// shortcuts are reversed.
//
// Only children that have leafletpage:navigatable set to TRUE can be navigated
// to.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwleaflet).
//
// The function takes the following parameters:
//
//   - canNavigateBack: new value.
func (self *Leaflet) SetCanNavigateBack(canNavigateBack bool) {
	var _arg0 *C.AdwLeaflet // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if canNavigateBack {
		_arg1 = C.TRUE
	}

	C.adw_leaflet_set_can_navigate_back(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(canNavigateBack)
}

// SetCanNavigateForward sets whether gestures and shortcuts for navigating
// forward are enabled.
//
// The supported gestures are:
//
// - One-finger swipe on touchscreens
//
// - Horizontal scrolling on touchpads (usually two-finger swipe)
//
// - Back/forward mouse buttons
//
// The keyboard back/forward keys are also supported, as well as the
// <kbd>Alt</kbd>+<kbd>→</kbd> shortcut for horizontal orientation, or
// <kbd>Alt</kbd>+<kbd>↓</kbd> for vertical orientation.
//
// If the orientation is horizontal, for right-to-left locales, gestures and
// shortcuts are reversed.
//
// Only children that have leafletpage:navigatable set to TRUE can be navigated
// to.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwleaflet).
//
// The function takes the following parameters:
//
//   - canNavigateForward: new value.
func (self *Leaflet) SetCanNavigateForward(canNavigateForward bool) {
	var _arg0 *C.AdwLeaflet // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if canNavigateForward {
		_arg1 = C.TRUE
	}

	C.adw_leaflet_set_can_navigate_forward(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(canNavigateForward)
}

// SetCanUnfold sets whether self can unfold.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwleaflet).
//
// The function takes the following parameters:
//
//   - canUnfold: whether self can unfold.
func (self *Leaflet) SetCanUnfold(canUnfold bool) {
	var _arg0 *C.AdwLeaflet // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if canUnfold {
		_arg1 = C.TRUE
	}

	C.adw_leaflet_set_can_unfold(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(canUnfold)
}

// SetChildTransitionParams sets the child transition spring parameters for
// self.
//
// The default value is equivalent to:
//
//	adw_spring_params_new (1, 0.5, 500)
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwleaflet).
//
// The function takes the following parameters:
//
//   - params: new parameters.
func (self *Leaflet) SetChildTransitionParams(params *SpringParams) {
	var _arg0 *C.AdwLeaflet      // out
	var _arg1 *C.AdwSpringParams // out

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.AdwSpringParams)(gextras.StructNative(unsafe.Pointer(params)))

	C.adw_leaflet_set_child_transition_params(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(params)
}

// SetFoldThresholdPolicy sets the fold threshold policy for self.
//
// If set to ADW_FOLD_THRESHOLD_POLICY_MINIMUM, it will only fold when the
// children cannot fit anymore. With ADW_FOLD_THRESHOLD_POLICY_NATURAL, it will
// fold as soon as children don't get their natural size.
//
// This can be useful if you have a long ellipsizing label and want to let it
// ellipsize instead of immediately folding.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwleaflet).
//
// The function takes the following parameters:
//
//   - policy to use.
func (self *Leaflet) SetFoldThresholdPolicy(policy FoldThresholdPolicy) {
	var _arg0 *C.AdwLeaflet            // out
	var _arg1 C.AdwFoldThresholdPolicy // out

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.AdwFoldThresholdPolicy(policy)

	C.adw_leaflet_set_fold_threshold_policy(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(policy)
}

// SetHomogeneous sets self to be homogeneous or not.
//
// If set to FALSE, different children can have different size along the
// opposite orientation.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwleaflet).
//
// The function takes the following parameters:
//
//   - homogeneous: whether to make self homogeneous.
func (self *Leaflet) SetHomogeneous(homogeneous bool) {
	var _arg0 *C.AdwLeaflet // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if homogeneous {
		_arg1 = C.TRUE
	}

	C.adw_leaflet_set_homogeneous(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(homogeneous)
}

// SetModeTransitionDuration sets the mode transition animation duration for
// self.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwleaflet).
//
// The function takes the following parameters:
//
//   - duration: new duration, in milliseconds.
func (self *Leaflet) SetModeTransitionDuration(duration uint) {
	var _arg0 *C.AdwLeaflet // out
	var _arg1 C.guint       // out

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.guint(duration)

	C.adw_leaflet_set_mode_transition_duration(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(duration)
}

// SetTransitionType sets the type of animation used for transitions between
// modes and children.
//
// The transition type can be changed without problems at runtime, so it is
// possible to change the animation based on the mode or child that is about to
// become current.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwleaflet).
//
// The function takes the following parameters:
//
//   - transition: new transition type.
func (self *Leaflet) SetTransitionType(transition LeafletTransitionType) {
	var _arg0 *C.AdwLeaflet              // out
	var _arg1 C.AdwLeafletTransitionType // out

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.AdwLeafletTransitionType(transition)

	C.adw_leaflet_set_transition_type(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(transition)
}

// SetVisibleChild sets the widget currently visible when the leaflet is folded.
//
// The transition is determined by leaflet:transition-type and
// leaflet:child-transition-params. The transition can be cancelled by the user,
// in which case visible child will change back to the previously visible child.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwleaflet).
//
// The function takes the following parameters:
//
//   - visibleChild: new child.
func (self *Leaflet) SetVisibleChild(visibleChild gtk.Widgetter) {
	var _arg0 *C.AdwLeaflet // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(visibleChild).Native()))

	C.adw_leaflet_set_visible_child(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(visibleChild)
}

// SetVisibleChildName makes the child with the name name visible.
//
// See leaflet:visible-child.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwleaflet).
//
// The function takes the following parameters:
//
//   - name of a child.
func (self *Leaflet) SetVisibleChildName(name string) {
	var _arg0 *C.AdwLeaflet // out
	var _arg1 *C.char       // out

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	C.adw_leaflet_set_visible_child_name(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(name)
}

// LeafletPageOverrides contains methods that are overridable.
type LeafletPageOverrides struct {
}

func defaultLeafletPageOverrides(v *LeafletPage) LeafletPageOverrides {
	return LeafletPageOverrides{}
}

// LeafletPage: auxiliary class used by leaflet.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwleaflet).
type LeafletPage struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*LeafletPage)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*LeafletPage, *LeafletPageClass, LeafletPageOverrides](
		GTypeLeafletPage,
		initLeafletPageClass,
		wrapLeafletPage,
		defaultLeafletPageOverrides,
	)
}

func initLeafletPageClass(gclass unsafe.Pointer, overrides LeafletPageOverrides, classInitFunc func(*LeafletPageClass)) {
	if classInitFunc != nil {
		class := (*LeafletPageClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapLeafletPage(obj *coreglib.Object) *LeafletPage {
	return &LeafletPage{
		Object: obj,
	}
}

func marshalLeafletPage(p uintptr) (interface{}, error) {
	return wrapLeafletPage(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Child gets the leaflet child to which self belongs.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwleaflet).
//
// The function returns the following values:
//
//   - widget: child to which self belongs.
func (self *LeafletPage) Child() gtk.Widgetter {
	var _arg0 *C.AdwLeafletPage // out
	var _cret *C.GtkWidget      // in

	_arg0 = (*C.AdwLeafletPage)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_leaflet_page_get_child(_arg0)
	runtime.KeepAlive(self)

	var _widget gtk.Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(gtk.Widgetter)
			return ok
		})
		rv, ok := casted.(gtk.Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_widget = rv
	}

	return _widget
}

// Name gets the name of self.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwleaflet).
//
// The function returns the following values:
//
//   - utf8 (optional): name of self.
func (self *LeafletPage) Name() string {
	var _arg0 *C.AdwLeafletPage // out
	var _cret *C.char           // in

	_arg0 = (*C.AdwLeafletPage)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_leaflet_page_get_name(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Navigatable gets whether the child can be navigated to when folded.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwleaflet).
//
// The function returns the following values:
//
//   - ok: whether self can be navigated to when folded.
func (self *LeafletPage) Navigatable() bool {
	var _arg0 *C.AdwLeafletPage // out
	var _cret C.gboolean        // in

	_arg0 = (*C.AdwLeafletPage)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_leaflet_page_get_navigatable(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetName sets the name of the self.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwleaflet).
//
// The function takes the following parameters:
//
//   - name (optional): new value to set.
func (self *LeafletPage) SetName(name string) {
	var _arg0 *C.AdwLeafletPage // out
	var _arg1 *C.char           // out

	_arg0 = (*C.AdwLeafletPage)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if name != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.adw_leaflet_page_set_name(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(name)
}

// SetNavigatable sets whether self can be navigated to when folded.
//
// If FALSE, the child will be ignored by leaflet.GetAdjacentChild,
// leaflet.Navigate, and swipe gestures.
//
// This can be used used to prevent switching to widgets like separators.
//
// Deprecated: See the migration guide
// (migrating-to-breakpoints.html#replace-adwleaflet).
//
// The function takes the following parameters:
//
//   - navigatable: whether self can be navigated to when folded.
func (self *LeafletPage) SetNavigatable(navigatable bool) {
	var _arg0 *C.AdwLeafletPage // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.AdwLeafletPage)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if navigatable {
		_arg1 = C.TRUE
	}

	C.adw_leaflet_page_set_navigatable(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(navigatable)
}

// LeafletClass: instance of this type is always passed by reference.
type LeafletClass struct {
	*leafletClass
}

// leafletClass is the struct that's finalized.
type leafletClass struct {
	native *C.AdwLeafletClass
}

func (l *LeafletClass) ParentClass() *gtk.WidgetClass {
	valptr := &l.native.parent_class
	var _v *gtk.WidgetClass // out
	_v = (*gtk.WidgetClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}

// LeafletPageClass: instance of this type is always passed by reference.
type LeafletPageClass struct {
	*leafletPageClass
}

// leafletPageClass is the struct that's finalized.
type leafletPageClass struct {
	native *C.AdwLeafletPageClass
}
