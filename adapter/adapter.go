package adapter

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"strings"
)

func minmax(a, b int) (int, int) {
	if a < b {
		return a, b
	} else {
		return b, a
	}
}

type line struct {
	X1, Y1, X2, Y2 int
}

type vectorImage struct {
	lines []line
}

func newRectangle(width, height int) *vectorImage {
	width -= 1
	height -= 1
	return &vectorImage{[]line{
		line{0, 0, width, 0},
		line{0, 0, 0, height},
		line{width, 0, width, height},
		line{0, height, width, height},
	}}
}

type point struct {
	X, Y int
}

type rasterImage interface {
	getPoints() []point
}

func drawPoints(owner rasterImage) string {
	maxX, maxY := 0, 0
	points := owner.getPoints()
	for _, pixel := range points {
		if pixel.X > maxX {
			maxX = pixel.X
		}
		if pixel.Y > maxY {
			maxY = pixel.Y
		}
	}
	maxX += 1
	maxY += 1

	// preallocate

	data := make([][]rune, maxY)
	for i := 0; i < maxY; i++ {
		data[i] = make([]rune, maxX)
		for j := range data[i] {
			data[i][j] = ' '
		}
	}

	for _, point := range points {
		data[point.Y][point.X] = '*'
	}

	b := strings.Builder{}
	for _, line := range data {
		b.WriteString(string(line))
		b.WriteRune('\n')
	}

	return b.String()
}

type vectorToRasterAdapter struct {
	points []point
}

func (a *vectorToRasterAdapter) addLine(line line) {
	left, right := minmax(line.X1, line.X2)
	top, bottom := minmax(line.Y1, line.Y2)
	dx := right - left
	dy := line.Y2 - line.Y1

	if dx == 0 {
		for y := top; y <= bottom; y++ {
			a.points = append(a.points, point{left, y})
		}
	} else if dy == 0 {
		for x := left; x <= right; x++ {
			a.points = append(a.points, point{x, top})
		}
	}

	fmt.Println("generated", len(a.points), "points")
}

var pointCache = map[[16]byte][]point{}

func (a *vectorToRasterAdapter) addLineCached(line line) {
	hash := func(obj interface{}) [16]byte {
		bytes, _ := json.Marshal(obj)
		return md5.Sum(bytes)
	}

	h := hash(line)
	if pts, ok := pointCache[h]; ok {
		for _, pt := range pts {
			a.points = append(a.points, pt)
		}
		return
	}

	left, right := minmax(line.X1, line.X2)
	top, bottom := minmax(line.Y1, line.Y2)
	dx := right - left
	dy := line.Y2 - line.Y1

	if dx == 0 {
		for y := top; y <= bottom; y++ {
			a.points = append(a.points, point{left, y})
		}
	} else if dy == 0 {
		for x := left; x <= right; x++ {
			a.points = append(a.points, point{x, top})
		}
	}

	pointCache[h] = a.points
	fmt.Println("we have", len(a.points), "points")
}

func (a vectorToRasterAdapter) getPoints() []point {
	return a.points
}

func vectorToRaster(vi *vectorImage) rasterImage {
	adapter := vectorToRasterAdapter{}
	for _, line := range vi.lines {
		adapter.addLineCached(line)
	}
	return adapter
}

func Run() {
	rc := newRectangle(6, 4)
	a := vectorToRaster(rc)
	_ = vectorToRaster(rc)
	fmt.Print(drawPoints(a))
}
