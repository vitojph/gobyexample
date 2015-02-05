package main

import "fmt"

func main() {

    // Strings, which can be added together with +.
    fmt.Println("go" + "lang" + "es chulo")

    // Integers and floats.

    fmt.Println("-13 + 7 =", -13+7)
    fmt.Println("7/3 =", 7/3)
    fmt.Println("7/3.0 =", 7/3.0)

    // Booleans, with boolean operators as you’d expect.

    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)
}
