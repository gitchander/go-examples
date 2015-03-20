package cairo

// #cgo pkg-config: cairo cairo-gobject
// #include <stdlib.h>
// #include <cairo.h>
// #include <cairo-gobject.h>
import "C"

type Pattern struct {
	pattern *C.cairo_pattern_t
}

func (p *Pattern) native() *C.cairo_pattern_t {

	if p == nil {
		return nil
	}

	return p.pattern
}

func NewPatternLinear(x0, y0, x1, y1 float64) *Pattern {
	p := C.cairo_pattern_create_linear(C.double(x0), C.double(y0), C.double(x1), C.double(y1))
	return &Pattern{p}
}

func NewPatternRadial(cx0, cy0, radius0, cx1, cy1, radius1 float64) *Pattern {
	p := C.cairo_pattern_create_radial(
		C.double(cx0), C.double(cy0), C.double(radius0),
		C.double(cx1), C.double(cy1), C.double(radius1))
	return &Pattern{p}
}

func NewPatternForSurface(s *Surface) *Pattern {
	p := C.cairo_pattern_create_for_surface(s.native())
	return &Pattern{p}
}

func (p *Pattern) Reference() *Pattern {
	reference := C.cairo_pattern_reference(p.native())
	return &Pattern{reference}
}

func (p *Pattern) Destroy() {
	C.cairo_pattern_destroy(p.native())
}

func (p *Pattern) AddColorStopRGB(offset, red, green, blue float64) {
	C.cairo_pattern_add_color_stop_rgb(p.native(), C.double(offset), C.double(red), C.double(green), C.double(blue))
}

func (p *Pattern) AddColorStopRGBA(offset, red, green, blue, alpha float64) {
	C.cairo_pattern_add_color_stop_rgba(p.native(), C.double(offset), C.double(red), C.double(green), C.double(blue), C.double(alpha))
}

func (p *Pattern) SetExtend(extend Extend) {
	C.cairo_pattern_set_extend(p.native(), C.cairo_extend_t(extend))
}

func (p *Pattern) SetMatrix(m Matrix) {
	C.cairo_pattern_set_matrix(p.native(), m.native())
}
