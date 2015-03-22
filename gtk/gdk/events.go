package gdk

// #cgo pkg-config: gdk-3.0
// #include <gdk/gdk.h>
import "C"

import (
	"unsafe"
)

/*
// Event is a representation of GDK's GdkEvent.
type Event struct {
	GdkEvent *C.GdkEvent
}

// native returns a pointer to the underlying GdkEvent.
func (v *Event) native() *C.GdkEvent {
	if v == nil {
		return nil
	}
	return v.GdkEvent
}

// Native returns a pointer to the underlying GdkEvent.
func (v *Event) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalEvent(p uintptr) (interface{}, error) {
	c := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return &Event{(*C.GdkEvent)(unsafe.Pointer(c))}, nil
}

func (v *Event) free() {
	C.gdk_event_free(v.native())
}

//----------------
// GdkEventButton
//----------------
type EventButton struct {
	*Event
}

func (e *EventButton) Native() uintptr {
	return uintptr(unsafe.Pointer(e.native()))
}

func (e *EventButton) native() *C.GdkEventButton {
	return (*C.GdkEventButton)(unsafe.Pointer(e.Event.native()))
}

func (e *EventButton) Pos() (x, y int) {
	x = int(e.native().x)
	y = int(e.native().y)
	return
}

*/

// https://developer.gnome.org/gdk3/stable/gdk3-Event-Structures.html#GdkEventButton
//----------------
// GdkEventButton
//----------------
type EventButton struct {
	event *C.GdkEventButton
}

func (e *EventButton) native() *C.GdkEventButton {
	return e.event
}

func (e *EventButton) FromNative(ptr uintptr) {
	e.event = (*C.GdkEventButton)(unsafe.Pointer(ptr))
}

func (e *EventButton) Native() uintptr {
	return uintptr(unsafe.Pointer(e.event))
}

func (e *EventButton) Pos() (x, y int) {
	x = int(e.native().x)
	y = int(e.native().y)
	return
}