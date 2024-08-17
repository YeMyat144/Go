package main

import (
    "fmt"
    "math"
)

// Define the Shape interface
type Shape interface {
    Area() float64
}

// Rectangle struct
type Rectangle struct {
    Width  float64
    Height float64
}

// Circle struct
type Circle struct {
    Radius float64
}

// Implement the Area method for Rectangle
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Implement the Area method for Circle
func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

// Function to print the area of any Shape
func PrintArea(s Shape) {
    fmt.Printf("Area: %.2f\n", s.Area())
}

func main() {
    rectangle := Rectangle{Width: 5, Height: 3}
    circle := Circle{Radius: 2.5}

    // Call PrintArea on rectangle and circle, both of which implement the Shape interface
    PrintArea(rectangle) // Prints the area of the rectangle
    PrintArea(circle)    // Prints the area of the circle
}