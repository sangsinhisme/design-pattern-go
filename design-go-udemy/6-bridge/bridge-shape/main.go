package main

import "fmt"

// Circle, Square
// Raster,Vector

// RasterCircle, RasterSquare, VectorCircle -> VectorSquare, ...
// Expand triangle, etc

// Renderer
type Renderer interface {
	RenderCircle(radius float32)
}

type VectorRenderer struct{}

func (v *VectorRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing a circle of radius", radius)
}

type RasterRenderer struct{}

func (r *RasterRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing pixels for circle of radius", radius)
}

type Circle struct {
	renderer Renderer
	radius   float32
}

func NewCircle(renderer Renderer, radius float32) *Circle {
	return &Circle{renderer, radius}
}

func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.radius)
}
func (c *Circle) Resize(factor float32) {
	c.radius *= factor
}

func main() {
	rasterRenderer := RasterRenderer{}
	//vectorRenderer := VectorRenderer{}
	circle := NewCircle(&rasterRenderer, 9)
	circle.Draw()
	circle.Resize(0.8)
	circle.Draw()
}
