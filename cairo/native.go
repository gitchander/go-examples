package cairo

// #cgo pkg-config: cairo cairo-gobject
// #include <stdlib.h>
// #include <cairo.h>
// #include <cairo-gobject.h>
import "C"

import (
	"unsafe"
)

func SurfaceToNative(s *Surface) uintptr {
	return uintptr(unsafe.Pointer(s.native()))
}

func CanvasToNative(c *Canvas) uintptr {
	return uintptr(unsafe.Pointer(c.native()))
}
