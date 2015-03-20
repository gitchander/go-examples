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

func (c *Canvas) Native() uintptr {
	return uintptr(unsafe.Pointer(c.native()))
}

func NewCanvas(s *Surface) (*Canvas, Status) {

	canvasNative := C.cairo_create(s.native())
	status := Status(C.cairo_status(canvasNative))
	if status != STATUS_SUCCESS {
		return nil, status
	}

	return &Canvas{canvasNative}, STATUS_SUCCESS
}

func NewCanvasNative(ptr uintptr) (*Canvas, Status) {

	canvasNative := (*C.cairo_t)(unsafe.Pointer(ptr))
	status := Status(C.cairo_status(canvasNative))
	if status != STATUS_SUCCESS {
		return nil, status
	}
	reference := C.cairo_reference(canvasNative)
	return &Canvas{reference}, STATUS_SUCCESS
}

func (c *Canvas) Reference() *Canvas {
	reference := C.cairo_reference(c.native())
	return &Canvas{reference}
}

func (c *Canvas) Destroy() {
	C.cairo_destroy(c.native())
}

func (c *Canvas) Status() Status {
	return Status(C.cairo_status(c.native()))
}

func (c *Canvas) Save() {
	C.cairo_save(c.native())
}

func (c *Canvas) Restore() {
	C.cairo_restore(c.native())
}

func (c *Canvas) GetTarget() *Surface {

	var surfaceNative *C.cairo_surface_t
	surfaceNative = C.cairo_get_target(c.native())
	if surfaceNative == nil {
		return nil
	}
	return &Surface{surfaceNative}
}

func (c *Canvas) PushGroup() {
	C.cairo_push_group(c.native())
}

func (c *Canvas) PushGroupWithContent(content Content) {
	C.cairo_push_group_with_content(c.native(), C.cairo_content_t(content))
}

// cairo_pop_group ()

// cairo_pop_group_to_source ()

func (c *Canvas) GetGroupTarget() *Surface {
	var surfaceNative *C.cairo_surface_t
	surfaceNative = C.cairo_get_group_target(c.native())
	if surfaceNative == nil {
		return nil
	}
	return &Surface{surfaceNative}
}

func (c *Canvas) SetSourceRGB(red, green, blue float64) {
	C.cairo_set_source_rgb(c.native(), C.double(red), C.double(green), C.double(blue))
}

func (c *Canvas) SetSourceRGBA(red, green, blue, alpha float64) {
	C.cairo_set_source_rgba(c.native(), C.double(red), C.double(green), C.double(blue), C.double(alpha))
}

func (c *Canvas) SetSource(p *Pattern) {
	C.cairo_set_source(c.native(), p.native())
}

func (c *Canvas) SetSourceSurface(s *Surface, x, y float64) {
	C.cairo_set_source_surface(c.native(), s.native(), C.double(x), C.double(y))
}

func (c *Canvas) GetSource() *Pattern {

	var (
		patternNative    *C.cairo_pattern_t
		patternReference *C.cairo_pattern_t
	)

	patternNative = C.cairo_get_source(c.native())
	patternReference = C.cairo_pattern_reference(patternNative)

	return &Pattern{patternReference}
}

func (c *Canvas) SetAntialias(antialias Antialias) {
	C.cairo_set_antialias(c.native(), C.cairo_antialias_t(antialias))
}

func (c *Canvas) GetAntialias() Antialias {
	return Antialias(C.cairo_get_antialias(c.native()))
}

func (c *Canvas) SetDash(dashes []float64, offset float64) {

	numDashes := C.int(len(dashes))
	ptrDashes := (*C.double)(&dashes[0])

	C.cairo_set_dash(c.native(),
		ptrDashes,
		numDashes,
		C.double(offset))
}

func (c *Canvas) GetDashCount() int {
	return int(C.cairo_get_dash_count(c.native()))
}

// cairo_get_dash ()

func (c *Canvas) SetFillRule(fillRule FillRule) {
	C.cairo_set_fill_rule(c.native(), C.cairo_fill_rule_t(fillRule))
}

func (c *Canvas) GetFillRule() FillRule {
	return FillRule(C.cairo_get_fill_rule(c.native()))
}

func (c *Canvas) SetLineCap(lineCap LineCap) {
	C.cairo_set_line_cap(c.native(), C.cairo_line_cap_t(lineCap))
}

func (c *Canvas) GetLineCap() LineCap {
	return LineCap(C.cairo_get_line_cap(c.native()))
}

func (c *Canvas) SetLineJoin(lineJoin LineJoin) {
	C.cairo_set_line_join(c.native(), C.cairo_line_join_t(lineJoin))
}

func (c *Canvas) GetLineJoin() LineJoin {
	return LineJoin(C.cairo_get_line_join(c.native()))
}

func (c *Canvas) SetLineWidth(width float64) {
	C.cairo_set_line_width(c.native(), C.double(width))
}

func (c *Canvas) GetLineWidth() float64 {
	return float64(C.cairo_get_line_width(c.native()))
}

func (c *Canvas) SetMiterLimit(limit float64) {
	C.cairo_set_miter_limit(c.native(), C.double(limit))
}

func (c *Canvas) GetMiterLimit() float64 {
	return float64(C.cairo_get_miter_limit(c.native()))
}

func (c *Canvas) SetOperator(operator Operator) {
	C.cairo_set_operator(c.native(), C.cairo_operator_t(operator))
}

func (c *Canvas) GetOperator() Operator {
	return Operator(C.cairo_get_operator(c.native()))
}

func (c *Canvas) SetTolerance(tolerance float64) {
	C.cairo_set_tolerance(c.native(), C.double(tolerance))
}

func (c *Canvas) GetTolerance() float64 {
	return float64(C.cairo_get_tolerance(c.native()))
}

func (c *Canvas) Clip() {
	C.cairo_clip(c.native())
}

func (c *Canvas) ClipPreserve() {
	C.cairo_clip_preserve(c.native())
}

// cairo_clip_extents ()

func (c *Canvas) InClip(x, y float64) bool {
	var b C.cairo_bool_t
	b = C.cairo_in_clip(c.native(), C.double(x), C.double(y))
	return boolGolang(b)
}

func (c *Canvas) ResetClip() {
	C.cairo_reset_clip(c.native())
}

// cairo_rectangle_list_destroy ()

// cairo_copy_clip_rectangle_list ()

func (c *Canvas) Fill() {
	C.cairo_fill(c.native())
}

func (c *Canvas) FillPreserve() {
	C.cairo_fill_preserve(c.native())
}

// cairo_fill_extents ()

func (c *Canvas) InFill(x, y float64) bool {
	var b C.cairo_bool_t
	b = C.cairo_in_fill(c.native(), C.double(x), C.double(y))
	return boolGolang(b)
}

// cairo_mask ()

// cairo_mask_surface ()

func (c *Canvas) Paint() {
	C.cairo_paint(c.native())
}

func (c *Canvas) PaintWithAlpha(alpha float64) {
	C.cairo_paint_with_alpha(c.native(), C.double(alpha))
}

func (c *Canvas) Stroke() {
	C.cairo_stroke(c.native())
}

func (c *Canvas) StrokePreserve() {
	C.cairo_stroke_preserve(c.native())
}

// cairo_stroke_extents ()

func (c *Canvas) InStroke(x, y float64) bool {
	var b C.cairo_bool_t
	b = C.cairo_in_stroke(c.native(), C.double(x), C.double(y))
	return boolGolang(b)
}

func (c *Canvas) CopyPage() {
	C.cairo_copy_page(c.native())
}

func (c *Canvas) ShowPage() {
	C.cairo_show_page(c.native())
}

func (c *Canvas) GetReferenceCount() uint {
	return uint(C.cairo_get_reference_count(c.native()))
}

// cairo_set_user_data ()

// cairo_get_user_data ()

//------------------------------------------
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

func (c *Canvas) NewPath() {
	C.cairo_new_path(c.native())
}

func (c *Canvas) NewSubPath() {
	C.cairo_new_sub_path(c.native())
}

func (c *Canvas) ClosePath() {
	C.cairo_close_path(c.native())
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
