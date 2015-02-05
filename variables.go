package main

import "fmt"

func main() {

    // var declares 1 or more variables.
    var cadena string = "esto es una cadena"
    fmt.Println(cadena)

    // You can declare multiple variables at once.
    var b, c int = 1, 2
    fmt.Println(b + c)

    // Go will infer the type of initialized variables.
    var d = true
    fmt.Println(d)

    // Variables declared without a corresponding initialization are zero-valued. For example, the zero value for an int is 0.
    var e1 int
    fmt.Println(e1)

    var e2 string
    fmt.Println(e2)

    var e3 float32
    fmt.Println(e3)

    // The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "short" in this case.
    f := "esto lo he creado con un atajo"
    fmt.Println(f)
}
