package cairo

// #cgo pkg-config: cairo cairo-gobject
// #include <stdlib.h>
// #include <cairo.h>
// #include <cairo-gobject.h>
import "C"

import (
	"unsafe"
)

type Canvas struct {
	canvas *C.cairo_t
}

func (c *Canvas) native() *C.cairo_t {

	if c == nil {
		return nil
	}

	return c.canvas
}

func NewCanvas(s *Surface) (*Canvas, Status) {

	canvasNative := C.cairo_create(s.native())
	status := Status(C.cairo_status(canvasNative))
	if status != STATUS_SUCCESS {
		return nil, status
	}

	return &Canvas{canvasNative}, STATUS_SUCCESS
}

func (c *Canvas) Destroy() {
	C.cairo_destroy(c.native())
}

func (c *Canvas) Status() Status {
	return Status(C.cairo_status(c.native()))
}

func (c *Canvas) MoveTo(x, y float64) {
	C.cairo_move_to(c.native(), C.double(x), C.double(y))
}

func (c *Canvas) LineTo(x, y float64) {
	C.cairo_line_to(c.native(), C.double(x), C.double(y))
}

func (c *Canvas) RelLineTo(dx, dy float64) {
	C.cairo_rel_line_to(c.native(), C.double(dx), C.double(dy))
}

func (c *Canvas) Rectangle(x, y, width, height float64) {
	C.cairo_rectangle(c.native(),
		C.double(x), C.double(y),
		C.double(width), C.double(height))
}

func (c *Canvas) Save() {
	C.cairo_save(c.native())
}

func (c *Canvas) Restore() {
	C.cairo_restore(c.native())
}

func (c *Canvas) Paint() {
	C.cairo_paint(c.native())
}

func (c *Canvas) PaintWithAlpha(alpha float64) {
	C.cairo_paint_with_alpha(c.native(), C.double(alpha))
}

func (c *Canvas) Fill() {
	C.cairo_fill(c.native())
}

func (c *Canvas) FillPreserve() {
	C.cairo_fill_preserve(c.native())
}

func (c *Canvas) Stroke() {
	C.cairo_stroke(c.native())
}

func (c *Canvas) Clip() {
	C.cairo_clip(c.native())
}

func (c *Canvas) NewPath() {
	C.cairo_new_path(c.native())
}

func (c *Canvas) NewSubPath() {
	C.cairo_new_sub_path(c.native())
}

func (c *Canvas) ClosePath() {
	C.cairo_close_path(c.native())
}

func (c *Canvas) SetSource(p *Pattern) {
	C.cairo_set_source(c.native(), p.native())
}

func (c *Canvas) SetSourceSurface(s *Surface, x, y float64) {
	C.cairo_set_source_surface(c.native(), s.native(), C.double(x), C.double(y))
}

func (c *Canvas) SetSourceRGB(red, green, blue float64) {
	C.cairo_set_source_rgb(c.native(), C.double(red), C.double(green), C.double(blue))
}

func (c *Canvas) SetSourceRGBA(red, green, blue, alpha float64) {
	C.cairo_set_source_rgba(c.native(), C.double(red), C.double(green), C.double(blue), C.double(alpha))
}

func (c *Canvas) SetLineWidth(width float64) {
	C.cairo_set_line_width(c.native(), C.double(width))
}

func (c *Canvas) SetLineJoin(lineJoin LineJoin) {
	C.cairo_set_line_join(c.native(), C.cairo_line_join_t(lineJoin))
}

func (c *Canvas) SetLineCap(lineCap LineCap) {
	C.cairo_set_line_cap(c.native(), C.cairo_line_cap_t(lineCap))
}

func (c *Canvas) SetDash(dashes []float64, offset float64) {

	numDashes := C.int(len(dashes))
	ptrDashes := (*C.double)(&dashes[0])

	C.cairo_set_dash(c.native(),
		ptrDashes,
		numDashes,
		C.double(offset))
}

func (c *Canvas) SetFillRule(fillRule FillRule) {
	C.cairo_set_fill_rule(c.native(), C.cairo_fill_rule_t(fillRule))
}

func (c *Canvas) Arc(xc, yc, radius, angle1, angle2 float64) {
	C.cairo_arc(c.native(), C.double(xc), C.double(yc), C.double(radius), C.double(angle1), C.double(angle2))
}

func (c *Canvas) ArcNegative(xc, yc, radius, angle1, angle2 float64) {
	C.cairo_arc_negative(c.native(), C.double(xc), C.double(yc), C.double(radius), C.double(angle1), C.double(angle2))
}

func (c *Canvas) CurveTo(x1, y1, x2, y2, x3, y3 float64) {

	C.cairo_curve_to(c.native(),
		C.double(x1), C.double(y1),
		C.double(x2), C.double(y2),
		C.double(x3), C.double(y3))
}

// Transformations

func (c *Canvas) Scale(sx, sy float64) {
	C.cairo_scale(c.native(), C.double(sx), C.double(sy))
}

func (c *Canvas) Translate(tx, ty float64) {
	C.cairo_translate(c.native(), C.double(tx), C.double(ty))
}

func (c *Canvas) Rotate(angle float64) {
	C.cairo_rotate(c.native(), C.double(angle))
}

func (c *Canvas) Transform(matrix *Matrix) {
	C.cairo_transform(c.native(), matrix.native())
}

func (c *Canvas) SetMatrix(matrix *Matrix) {
	C.cairo_set_matrix(c.native(), matrix.native())
}

func (c *Canvas) GetMatrix(matrix *Matrix) {
	C.cairo_get_matrix(c.native(), matrix.native())
}

func (c *Canvas) IdentityMatrix() {
	C.cairo_identity_matrix(c.native())
}

// Font

func (c *Canvas) SelectFontFace(family string, fontSlant FontSlant, fontWeight FontWeight) {

	cstrFamily := C.CString(family)
	defer C.free(unsafe.Pointer(cstrFamily))

	C.cairo_select_font_face(c.native(), cstrFamily,
		C.cairo_font_slant_t(fontSlant),
		C.cairo_font_weight_t(fontWeight))
}

func (c *Canvas) SetFontSize(size float64) {
	C.cairo_set_font_size(c.native(), C.double(size))
}

// Text

func (c *Canvas) ShowText(text string) {

	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))

	C.cairo_show_text(c.native(), cstr)
}

func (c *Canvas) TextPath(text string) {

	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))

	C.cairo_text_path(c.native(), cstr)
}

type TextExtents struct {
	BearingX float64
	BearingY float64
	Width    float64
	Height   float64
	AdvanceX float64
	AdvanceY float64
}

func (c *Canvas) TextExtents(text string, textExtents *TextExtents) {

	if textExtents == nil {
		return
	}

	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))

	var extents C.cairo_text_extents_t

	C.cairo_text_extents(c.native(), cstr, &extents)

	textExtents.BearingX = float64(extents.x_bearing)
	textExtents.BearingY = float64(extents.y_bearing)
	textExtents.Width = float64(extents.width)
	textExtents.Height = float64(extents.height)
	textExtents.AdvanceX = float64(extents.x_advance)
	textExtents.AdvanceY = float64(extents.y_advance)
}
