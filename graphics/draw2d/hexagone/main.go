package main

import (
	"bufio"
	"image"
	"image/png"
	"log"
	"os"

	"github.com/llgcode/draw2d/draw2dimg"

	"github.com/gitchander/heuristic/math/graph2d"
	"github.com/gitchander/heuristic/math/hexm"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func saveToPngFile(filePath string, m image.Image) error {

	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	bw := bufio.NewWriter(f)
	defer bw.Flush()

	return png.Encode(bw, m)
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {

	ps := hexm.HexVertexes(hexm.Angled)

	var (
		width  = 256
		height = 256
	)

	m := image.NewRGBA(image.Rect(0, 0, width, height))

	gc := draw2dimg.NewGraphicContext(m)

	var cellRadius = float32(minInt(width, height)) * 0.4

	var centerX = float32(width) * 0.5
	var centerY = float32(height) * 0.5

	T := graph2d.NewTransform()
	T.Scale(cellRadius, cellRadius)
	T.Move(centerX, centerY)

	v := graph2d.Vector{float32(ps[0].X), float32(ps[0].Y)}
	v = T.Apply(v)

	x0 := float64(v.X)
	y0 := float64(v.Y)

	gc.MoveTo(x0, y0)
	for i := 1; i < len(ps); i++ {

		v = graph2d.Vector{float32(ps[i].X), float32(ps[i].Y)}
		v = T.Apply(v)

		x := float64(v.X)
		y := float64(v.Y)

		gc.LineTo(x, y)
	}
	gc.LineTo(x0, y0)

	gc.Stroke()

	err := saveToPngFile("test_path.png", m)
	checkError(err)
}
