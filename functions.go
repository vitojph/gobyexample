// _Functions_ are central in Go. We'll learn about
// functions with a few different examples.

package main

import "fmt"

// Here's a function that takes two `int`s and returns
// their sum as an `int`.
func suma(a int, b int) int {
    // Go requires explicit returns, i.e. it won't
    // automatically return the value of the last
    // expression.
    return a + b
}

func cuadrado(n int) int {
    return n*n
}

func es_par(n int) bool {
    if n%2 == 0 {
        return true
    } else {
        return false
    }
}

func es_impar(n int) bool {
    if n%2 == 0 {
        return false
    } else {
        return true
    }
}

func esmultiplo(n1, n2 int) bool {
    if n1%n2 == 0 {
        return true
    } else {
        return false
    }
}

// Variadic functions can be called with any number of trailing arguments.
func sumatodo(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func main() {

    // Call a function just as you'd expect, with
    // `name(args)`.
    resultado := suma(1, 2)
    fmt.Println("1+2 =", resultado)


    // probemos
    for i := 0;  i < 10 ; i++ {
        fmt.Printf("el cuadrado de %d es %d\n", i, cuadrado(i))
        if es_par(i) {
            fmt.Printf("%d es par\n\n", i)
        }

        if es_impar(i) {
            fmt.Printf("%d es impar\n\n", i)
        }
    }

    // repite un bucle 5 veces
    for i, _ := range[5]int{} {
        fmt.Println("voy por el", i)
    }

    sumatodo(1, 2, 3, 4, 5, 6)
    sumatodo(12, -45, 7, 852)
}
