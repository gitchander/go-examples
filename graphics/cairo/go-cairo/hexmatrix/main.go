package main

import (
	"log"

	"github.com/gitchander/cairo"
	"github.com/gitchander/heuristic/math/hexm"
)

func main() {

	m := hexm.NewMatrix(hexm.Coord{X: 3, Y: 3, Z: 3})

	const cellRadius = 56.2

	texture, err := cairo.NewSurfaceFromPNG("hexagone.png")
	checkError(err)
	defer texture.Destroy()

	surface, err := cairo.NewSurface(cairo.FORMAT_ARGB32, 512, 512)
	checkError(err)
	defer surface.Destroy()

	canvas, err := cairo.NewCanvas(surface)
	checkError(err)
	defer canvas.Destroy()

	shift := getCenter(surface).Sub(getCenter(texture))
	for I := hexm.NewIterator(m); I.HasValue(); I.Next() {
		coord := I.Coord()
		v := hexm.CoordToVector(hexm.Flat, coord)
		v = v.MulScalar(cellRadius)
		v = v.Add(shift)
		canvas.SetSourceSurface(texture, v.X, v.Y)
		canvas.Paint()
	}

	err = surface.WriteToPNG("result.png")
	checkError(err)

	surface.Finish()
}

func getCenter(surface *cairo.Surface) hexm.Vector {
	return hexm.Vector{
		float64(surface.GetWidth()) * 0.5,
		float64(surface.GetHeight()) * 0.5,
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
