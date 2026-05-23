package main

import "fmt"

type engine struct {
	horsePower int
}

func (e engine) sound() string {
	return "vroom"
}

type car struct {
	model string
	color string
	engine
}

func (c car) sound() string {
	return "car vroom"
}

func main() {
	myCar := car{
		model: "Toyota",
		color: "red",
		engine: engine{
			horsePower: 150,
		},
	}
	fmt.Println(myCar)
	fmt.Println(myCar.sound())        // car vroom
	fmt.Println(myCar.engine.sound()) // vroom    // here we achieve method overriding by calling the sound method of the engine struct using the embedded struct syntax.

}

//output
// {Toyota red {150}}
// car vroom
// vroom
