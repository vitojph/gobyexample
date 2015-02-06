package main

import "fmt"


// a struct with PoS features
type word struct {
    token string
    pos_tag string
    gender string
    number string
    proper  bool
}


func main() {

    //This syntax creates a new struct.
    fmt.Println(word{"amigo", "nn", "masc", "sing", false})

    // Omitted fields will be zero-valued.
    fmt.Println(word{token: "amiga", pos_tag: "nn"})


    // Access struct fields with a dot.
    azul := word{token: "azul", pos_tag: "jj"}
    fmt.Println(azul.pos_tag)

    // Structs are mutable
    azul.gender = "unsp_g"
    azul.number = "sing"
    azul.proper = false
    fmt.Println(azul)
}
