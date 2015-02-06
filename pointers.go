package main

import "fmt"


func sumauno1(x int) {
    x += 1
}

// con punteros
func sumauno2(xp *int) {
    *xp += 1
}


func main() {

    // primero, sin punteros
    x1 := 5
    for i := 0;  i < 5; i++ {
        sumauno1(x1)
        fmt.Println(x1) // x1 is still 5
    }

    // luego, con punteros
    x2 := 5
    for i := 0;  i < 5; i++ {
        sumauno2(&x2)
        fmt.Println(x2) // x2 is summing up
    }

}
