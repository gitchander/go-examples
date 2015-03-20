package cairo

// #cgo pkg-config: cairo cairo-gobject
// #include <stdlib.h>
// #include <cairo.h>
// #include <cairo-gobject.h>
import "C"

import (
	"unsafe"
)

type Surface struct {
	surface *C.cairo_surface_t
}

func (s *Surface) native() *C.cairo_surface_t {

	if s == nil {
		return nil
	}

	return s.surface
}

func (s *Surface) Native() uintptr {
	return uintptr(unsafe.Pointer(s.native()))
}

func NewSurface(format Format, width, height int) *Surface {

	surfaceNative := C.cairo_image_surface_create(C.cairo_format_t(format), C.int(width), C.int(height))
	return &Surface{surfaceNative}
}

func NewSurfaceFromPNG(fileName string) (*Surface, Status) {

	cstr := C.CString(fileName)
	defer C.free(unsafe.Pointer(cstr))

	surfaceNative := C.cairo_image_surface_create_from_png(cstr)

	status := Status(C.cairo_surface_status(surfaceNative))
	if status != STATUS_SUCCESS {
		return nil, status
	}

	return &Surface{surfaceNative}, STATUS_SUCCESS
}

func (s *Surface) Destroy() {
	C.cairo_surface_destroy(s.native())
}

func (s *Surface) Finish() {
	C.cairo_surface_finish(s.native())
}

func (s *Surface) WriteToPNG(fileName string) Status {

	cstr := C.CString(fileName)
	defer C.free(unsafe.Pointer(cstr))

	return Status(C.cairo_surface_write_to_png(s.native(), cstr))
}

func (s *Surface) GetWidth() int {
	return int(C.cairo_image_surface_get_width(s.native()))
}

func (s *Surface) GetHeight() int {
	return int(C.cairo_image_surface_get_height(s.native()))
}
