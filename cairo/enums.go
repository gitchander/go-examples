package cairo

// #cgo pkg-config: cairo cairo-gobject
// #include <stdlib.h>
// #include <cairo.h>
// #include <cairo-gobject.h>
import "C"

type Status int

const (
	STATUS_SUCCESS                   Status = C.CAIRO_STATUS_SUCCESS
	STATUS_NO_MEMORY                 Status = C.CAIRO_STATUS_NO_MEMORY
	STATUS_INVALID_RESTORE           Status = C.CAIRO_STATUS_INVALID_RESTORE
	STATUS_INVALID_POP_GROUP         Status = C.CAIRO_STATUS_INVALID_POP_GROUP
	STATUS_NO_CURRENT_POINT          Status = C.CAIRO_STATUS_NO_CURRENT_POINT
	STATUS_INVALID_MATRIX            Status = C.CAIRO_STATUS_INVALID_MATRIX
	STATUS_INVALID_STATUS            Status = C.CAIRO_STATUS_INVALID_STATUS
	STATUS_NULL_POINTER              Status = C.CAIRO_STATUS_NULL_POINTER
	STATUS_INVALID_STRING            Status = C.CAIRO_STATUS_INVALID_STRING
	STATUS_INVALID_PATH_DATA         Status = C.CAIRO_STATUS_INVALID_PATH_DATA
	STATUS_READ_ERROR                Status = C.CAIRO_STATUS_READ_ERROR
	STATUS_WRITE_ERROR               Status = C.CAIRO_STATUS_WRITE_ERROR
	STATUS_SURFACE_FINISHED          Status = C.CAIRO_STATUS_SURFACE_FINISHED
	STATUS_SURFACE_TYPE_MISMATCH     Status = C.CAIRO_STATUS_SURFACE_TYPE_MISMATCH
	STATUS_PATTERN_TYPE_MISMATCH     Status = C.CAIRO_STATUS_PATTERN_TYPE_MISMATCH
	STATUS_INVALID_CONTENT           Status = C.CAIRO_STATUS_INVALID_CONTENT
	STATUS_INVALID_FORMAT            Status = C.CAIRO_STATUS_INVALID_FORMAT
	STATUS_INVALID_VISUAL            Status = C.CAIRO_STATUS_INVALID_VISUAL
	STATUS_FILE_NOT_FOUND            Status = C.CAIRO_STATUS_FILE_NOT_FOUND
	STATUS_INVALID_DASH              Status = C.CAIRO_STATUS_INVALID_DASH
	STATUS_INVALID_DSC_COMMENT       Status = C.CAIRO_STATUS_INVALID_DSC_COMMENT
	STATUS_INVALID_INDEX             Status = C.CAIRO_STATUS_INVALID_INDEX
	STATUS_CLIP_NOT_REPRESENTABLE    Status = C.CAIRO_STATUS_CLIP_NOT_REPRESENTABLE
	STATUS_TEMP_FILE_ERROR           Status = C.CAIRO_STATUS_TEMP_FILE_ERROR
	STATUS_INVALID_STRIDE            Status = C.CAIRO_STATUS_INVALID_STRIDE
	STATUS_FONT_TYPE_MISMATCH        Status = C.CAIRO_STATUS_FONT_TYPE_MISMATCH
	STATUS_USER_FONT_IMMUTABLE       Status = C.CAIRO_STATUS_USER_FONT_IMMUTABLE
	STATUS_USER_FONT_ERROR           Status = C.CAIRO_STATUS_USER_FONT_ERROR
	STATUS_NEGATIVE_COUNT            Status = C.CAIRO_STATUS_NEGATIVE_COUNT
	STATUS_INVALID_CLUSTERS          Status = C.CAIRO_STATUS_INVALID_CLUSTERS
	STATUS_INVALID_SLANT             Status = C.CAIRO_STATUS_INVALID_SLANT
	STATUS_INVALID_WEIGHT            Status = C.CAIRO_STATUS_INVALID_WEIGHT
	STATUS_INVALID_SIZE              Status = C.CAIRO_STATUS_INVALID_SIZE
	STATUS_USER_FONT_NOT_IMPLEMENTED Status = C.CAIRO_STATUS_USER_FONT_NOT_IMPLEMENTED
	STATUS_DEVICE_TYPE_MISMATCH      Status = C.CAIRO_STATUS_DEVICE_TYPE_MISMATCH
	STATUS_DEVICE_ERROR              Status = C.CAIRO_STATUS_DEVICE_ERROR
)

var key_Status = map[Status]string{

	STATUS_SUCCESS:                   "CAIRO_STATUS_SUCCESS",
	STATUS_NO_MEMORY:                 "CAIRO_STATUS_NO_MEMORY",
	STATUS_INVALID_RESTORE:           "CAIRO_STATUS_INVALID_RESTORE",
	STATUS_INVALID_POP_GROUP:         "CAIRO_STATUS_INVALID_POP_GROUP",
	STATUS_NO_CURRENT_POINT:          "CAIRO_STATUS_NO_CURRENT_POINT",
	STATUS_INVALID_MATRIX:            "CAIRO_STATUS_INVALID_MATRIX",
	STATUS_INVALID_STATUS:            "CAIRO_STATUS_INVALID_STATUS",
	STATUS_NULL_POINTER:              "CAIRO_STATUS_NULL_POINTER",
	STATUS_INVALID_STRING:            "CAIRO_STATUS_INVALID_STRING",
	STATUS_INVALID_PATH_DATA:         "CAIRO_STATUS_INVALID_PATH_DATA",
	STATUS_READ_ERROR:                "CAIRO_STATUS_READ_ERROR",
	STATUS_WRITE_ERROR:               "CAIRO_STATUS_WRITE_ERROR",
	STATUS_SURFACE_FINISHED:          "CAIRO_STATUS_SURFACE_FINISHED",
	STATUS_SURFACE_TYPE_MISMATCH:     "CAIRO_STATUS_SURFACE_TYPE_MISMATCH",
	STATUS_PATTERN_TYPE_MISMATCH:     "CAIRO_STATUS_PATTERN_TYPE_MISMATCH",
	STATUS_INVALID_CONTENT:           "CAIRO_STATUS_INVALID_CONTENT",
	STATUS_INVALID_FORMAT:            "CAIRO_STATUS_INVALID_FORMAT",
	STATUS_INVALID_VISUAL:            "CAIRO_STATUS_INVALID_VISUAL",
	STATUS_FILE_NOT_FOUND:            "CAIRO_STATUS_FILE_NOT_FOUND",
	STATUS_INVALID_DASH:              "CAIRO_STATUS_INVALID_DASH",
	STATUS_INVALID_DSC_COMMENT:       "CAIRO_STATUS_INVALID_DSC_COMMENT",
	STATUS_INVALID_INDEX:             "CAIRO_STATUS_INVALID_INDEX",
	STATUS_CLIP_NOT_REPRESENTABLE:    "CAIRO_STATUS_CLIP_NOT_REPRESENTABLE",
	STATUS_TEMP_FILE_ERROR:           "CAIRO_STATUS_TEMP_FILE_ERROR",
	STATUS_INVALID_STRIDE:            "CAIRO_STATUS_INVALID_STRIDE",
	STATUS_FONT_TYPE_MISMATCH:        "CAIRO_STATUS_FONT_TYPE_MISMATCH",
	STATUS_USER_FONT_IMMUTABLE:       "CAIRO_STATUS_USER_FONT_IMMUTABLE",
	STATUS_USER_FONT_ERROR:           "CAIRO_STATUS_USER_FONT_ERROR",
	STATUS_NEGATIVE_COUNT:            "CAIRO_STATUS_NEGATIVE_COUNT",
	STATUS_INVALID_CLUSTERS:          "CAIRO_STATUS_INVALID_CLUSTERS",
	STATUS_INVALID_SLANT:             "CAIRO_STATUS_INVALID_SLANT",
	STATUS_INVALID_WEIGHT:            "CAIRO_STATUS_INVALID_WEIGHT",
	STATUS_INVALID_SIZE:              "CAIRO_STATUS_INVALID_SIZE",
	STATUS_USER_FONT_NOT_IMPLEMENTED: "CAIRO_STATUS_USER_FONT_NOT_IMPLEMENTED",
	STATUS_DEVICE_TYPE_MISMATCH:      "CAIRO_STATUS_DEVICE_TYPE_MISMATCH",
	STATUS_DEVICE_ERROR:              "CAIRO_STATUS_DEVICE_ERROR",
}

func StatusToString(status Status) string {

	s, ok := key_Status[status]
	if !ok {
		s = "CAIRO_STATUS_UNDEFINED"
	}

	return s
}

type Format int

const (
	FORMAT_INVALID   Format = C.CAIRO_FORMAT_INVALID
	FORMAT_ARGB32    Format = C.CAIRO_FORMAT_ARGB32
	FORMAT_RGB24     Format = C.CAIRO_FORMAT_RGB24
	FORMAT_A8        Format = C.CAIRO_FORMAT_A8
	FORMAT_A1        Format = C.CAIRO_FORMAT_A1
	FORMAT_RGB16_565 Format = C.CAIRO_FORMAT_RGB16_565
	FORMAT_RGB30     Format = C.CAIRO_FORMAT_RGB30
)

type LineJoin int

const (
	LINE_JOIN_MITER LineJoin = C.CAIRO_LINE_JOIN_MITER
	LINE_JOIN_ROUND LineJoin = C.CAIRO_LINE_JOIN_ROUND
	LINE_JOIN_BEVEL LineJoin = C.CAIRO_LINE_JOIN_BEVEL
)

type LineCap int

const (
	LINE_CAP_BUTT   LineCap = C.CAIRO_LINE_CAP_BUTT
	LINE_CAP_ROUND  LineCap = C.CAIRO_LINE_CAP_ROUND
	LINE_CAP_SQUARE LineCap = C.CAIRO_LINE_CAP_SQUARE
)

type FillRule int

const (
	FILL_RULE_WINDING  FillRule = C.CAIRO_FILL_RULE_WINDING
	FILL_RULE_EVEN_ODD FillRule = C.CAIRO_FILL_RULE_EVEN_ODD
)

type Extend int

const (
	EXTEND_NONE    Extend = C.CAIRO_EXTEND_NONE
	EXTEND_REPEAT  Extend = C.CAIRO_EXTEND_REPEAT
	EXTEND_REFLECT Extend = C.CAIRO_EXTEND_REFLECT
	EXTEND_PAD     Extend = C.CAIRO_EXTEND_PAD
)

type FontSlant int

const (
	FONT_SLANT_NORMAL  FontSlant = C.CAIRO_FONT_SLANT_NORMAL
	FONT_SLANT_ITALIC  FontSlant = C.CAIRO_FONT_SLANT_ITALIC
	FONT_SLANT_OBLIQUE FontSlant = C.CAIRO_FONT_SLANT_OBLIQUE
)

type FontWeight int

const (
	FONT_WEIGHT_NORMAL FontWeight = C.CAIRO_FONT_WEIGHT_NORMAL
	FONT_WEIGHT_BOLD   FontWeight = C.CAIRO_FONT_WEIGHT_BOLD
)
