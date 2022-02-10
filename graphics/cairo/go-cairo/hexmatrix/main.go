package main

import (
	"log"

	"github.com/gitchander/cairo"
	"github.com/gitchander/heuristic/math/hexm"
)

func main() {
	err := run()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {

	m := hexm.NewMatrix(hexm.Coord{X: 3, Y: 3, Z: 3})

	const cellRadius = 56.2

	texture, err := cairo.NewSurfaceFromPNG("hexagone.png")
	if err != nil {
		return err
	}
	defer texture.Destroy()

	surface, err := cairo.NewSurface(cairo.FORMAT_ARGB32, 512, 512)
	if err != nil {
		return err
	}
	defer surface.Destroy()

	canvas, err := cairo.NewCanvas(surface)
	if err != nil {
		return err
	}
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
	if err != nil {
		return err
	}

	surface.Finish()

	return nil
}

func getCenter(surface *cairo.Surface) hexm.Vector {
	return hexm.Vector{
		X: float64(surface.GetWidth()) / 2,
		Y: float64(surface.GetHeight()) / 2,
	}
}
