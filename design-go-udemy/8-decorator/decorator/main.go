package main

import "fmt"

type Shape interface {
	Render() string
}

type Circle struct {
	Radius float32
}

func (c *Circle) Render() string {
	return fmt.Sprintf("Circle of radius %f", c.Radius)
}

func (c *Circle) Resize(factor float32) {
	c.Radius *= factor
}

type Square struct {
	Side float32
}

func (s *Square) Render() string {
	return fmt.Sprintf("Square with side %f", s.Side)
}

/*
Image you have these circles and squares in your system, but know you want to color it
One approach -> jump to Square and add additional member (Square -> color) but will conflict open closed principle
	-> don't know how it effects to behavior
Solve: extend type, one possibility is that you can simply aggregate and thereby extending the type
*/

/*
type ColorSquare struct {
	Square
	Color string
}
Now on the one hand, this could work if you have a small number of structs that you want to extend this way
On the other hand, if you have entire groups of structs like circles & squares & triangles -> sorts of other shapes
-> counterpart colored class for every single one of those types just doesn't make sense
*/

// Decorator
type ColoredShape struct {
	Shape Shape
	Color string
}

func (c *ColoredShape) Render() string {
	return fmt.Sprintf("%s has the color %s", c.Shape.Render(), c.Color)
}

/*
Layer decorator on decorator
*/
type TransparentShape struct {
	Shape        Shape
	Transparency float32
}

func (t *TransparentShape) Render() string {
	return fmt.Sprintf("%s has %f%% transparency", t.Shape.Render(), t.Transparency*100.0)
}

func main() {
	circle := Circle{2}

	// Once you've made a decorator, once you've put colored shape over the ordinary shape
	// You cannot say redCircle.Resize() because the resize method is no longer available
	circle.Resize(2)
	fmt.Println(circle.Render())

	redCircle := ColoredShape{&circle, "Red"}
	//redCircle -> can not call Resize()
	fmt.Println(redCircle.Render())

	//Layer decorator on decorator
	rhsCircle := TransparentShape{&redCircle, 0.5}
	fmt.Println(rhsCircle.Render())

}
