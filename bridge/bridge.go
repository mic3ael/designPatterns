package bridge

import "fmt"

type renderer interface {
	renderCircle(radius float32)
}

type vectorRenderer struct {
	renderCircle (float32)
}

func (v *vectorRenderer) renderCicle(radius float32) {
	fmt.Println("Deawing a circle of radius", radius)
}

type rasterRenderer struct {
	dpi int
}

func (r *rasterRenderer) renderCircle(radius float32) {
	fmt.Println("Drawing pixels for circle of radius", radius)
}

type circle struct {
	render renderer
	radius float32
}

func newCircle(render renderer, radius float32) *circle {
	return &circle{render, radius}
}

func (c *circle) draw() {
	c.render.renderCircle(c.radius)
}

func (c *circle) resize(factor float32) {
	c.radius *= factor
}

func Run() {
	raster := rasterRenderer{}
	// vector := vectorRenderer{}
	circle := newCircle(&raster, 5)
	circle.draw()
	circle.resize(2)
	circle.draw()
}
