package main

import "fmt"

type rectangulo struct {
    altura, anchura int
}

// This area method has a receiver type of *rect.
func (r *rectangulo) area() int {
    return r.altura * r.anchura
}

// Methods can be defined for either pointer or value receiver types. Here’s an example of a value receiver.

func (r rectangulo) perimetro() int {
    return 2*r.altura + 2*r.anchura
}


func main() {
    r := rectangulo{altura: 10, anchura: 5}

    // Here we call the 2 methods defined for our struct.
    fmt.Println("área: ", r.area())
    fmt.Println("perímetro", r.perimetro())

    // Go automatically handles conversion between values and pointers for method calls. You may want to use a pointer receiver type to avoid copying on method calls or to allow the method to mutate the receiving struct.

    rp := &r
    fmt.Println("área: ", rp.area())
    fmt.Println("perímetro:", rp.perimetro())
}
