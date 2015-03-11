package main

import (
	"github.com/ungerik/go-cairo"

	"github.com/gitchander/heuristic/math/graph2d"
	"github.com/gitchander/heuristic/math/hexm"
)

func getCenter(surface *cairo.Surface) graph2d.Vector {
	return graph2d.Vector{
		float32(surface.GetWidth()) * 0.5,
		float32(surface.GetHeight()) * 0.5,
	}
}

func main() {

	mSize, _ := hexm.NewSize(3, 3, 3)
	m := hexm.NewMatrix(mSize)

	const cellRadius = float32(56.2)

	texture := cairo.NewSurfaceFromPNG("./hexagone-gray.png")
	surface := cairo.NewSurface(cairo.FORMAT_ARGB32, 512, 512)

	shift := getCenter(surface).Sub(getCenter(texture))

	for I := m.NewIterator(); !I.Done(); I.Next() {

		coord, _, _ := I.Current()

		v := hexm.CoordToVector(coord)
		v = v.MulScalar(cellRadius)
		v = v.Add(shift)

		surface.SetSourceSurface(texture, float64(v.X), float64(v.Y))
		surface.Paint()
	}

	surface.WriteToPNG("./test.png")
	surface.Finish()
}
