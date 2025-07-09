package main

import "fmt"

//////////////////////////////////////////////////////
// Step 1: Define the Interface (Contract)
//////////////////////////////////////////////////////

// IShape defines the behavior expected from any shape
type IShape interface {
	Area() float32
	Perimeter() float32
}

//////////////////////////////////////////////////////
// Step 2: Implement a Concrete Type (Rectangle)
//////////////////////////////////////////////////////

// Rectangle is a concrete type that has length and breadth
type Rectangle struct {
	length  float32
	breadth float32
}

// Method to calculate Area - part of IShape
func (r *Rectangle) Area() float32 {
	return r.length * r.breadth
}

// Method to calculate Perimeter - part of IShape
func (r *Rectangle) Perimeter() float32 {
	return 2 * (r.length + r.breadth)
}

//////////////////////////////////////////////////////
// Step 3: ShapeService - depends on IShape interface
//////////////////////////////////////////////////////

// ShapeService depends on an IShape — not on Rectangle directly
type ShapeService struct {
	shape IShape // <-- This is Dependency Injection
}

// PrintDetails uses the injected shape to get area and perimeter
func (s *ShapeService) PrintDetails() {
	fmt.Println("Area:", s.shape.Area())
	fmt.Println("Perimeter:", s.shape.Perimeter())
}

//////////////////////////////////////////////////////
// Step 4: Main - Inject dependencies
//////////////////////////////////////////////////////

func main() {
	// Create a rectangle
	rect := &Rectangle{length: 10, breadth: 5}

	// Inject it into the ShapeService
	service := ShapeService{shape: rect}

	// Use the service
	service.PrintDetails()
}
