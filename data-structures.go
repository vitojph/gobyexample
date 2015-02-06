package main

import "fmt"

func main() {
    // arrays

    fmt.Println("ARRAYS")
    // Here we create an array `a` that will hold exactly
    // 5 `int`s. The type of elements and length are both
    // part of the array's type. By default an array is
    // zero-valued, which for `int`s means `0`s.
    var a [5]int
    fmt.Println("empty array:", a)

    // We can set a value at an index using the
    // `array[index] = value` syntax, and get a value with
    // `array[index]`.
    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])

    // The builtin `len` returns the length of an array.
    fmt.Println("len:", len(a))

    // slices
    fmt.Println("SLICES")

    // Unlike arrays, slices are typed only by the
    // elements they contain (not the number of elements).
    // To create an empty slice with non-zero length, use
    // the builtin `make`. Here we make a slice of
    // `string`s of length `6` (initially zero-valued).
    s := make([]string, 6)
    fmt.Println("empty slice:", s)

    // We can set and get just like with arrays.
    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])

    // We can declare and initialize a variable for slice
    // in a single line as well.
    t := []string{"a", "b", "c"}
    fmt.Println("dcl:", t)

    // `len` returns the length of the slice as expected.
    fmt.Println("len:", len(s))

    // In addition to these basic operations, slices
    // support several more that make them richer than
    // arrays. One is the builtin `append`, which
    // returns a slice containing one or more new values.
    // Note that we need to accept a return value from
    // append as we may get a new slice value.
    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)

    // Slices can also be `copy`'d. Here we create an
    // empty slice `c` of the same length as `s` and copy
    // into `c` from `s`.
    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("copy:", c)

    // Slices support a "slice" operator with the syntax
    // `slice[low:high]`. For example, this gets a slice
    // of the elements `s[2]`, `s[3]`, and `s[4]`.
    slice := make([]string, 5)
    slice[0] = "uno"
    slice[1] = "dos"
    slice[2] = "tres"
    slice[3] = "cuatro"
    slice[4] = "cinco"
    // slice[7] = "hola" errors raise panic message
    fmt.Println("los tres primeros:", slice[:3])
    fmt.Println("del segundo al cuarto:", slice[1:4])
    fmt.Println("desde el tercero:", slice[2:])

    // Slices can be composed into multi-dimensional data
    // structures. The length of the inner slices can
    // vary, unlike with multi-dimensional arrays.
    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
     }
    fmt.Println("2d: ", twoD)}

    // maps
    fmt.Println("MAPS")
    // To create an empty map, use the builtin `make`:
    // `make(map[key-type]val-type)`.
    m := make(map[string]int)

    // Set key/value pairs using typical `name[key] = val`
    // syntax.
    m["k1"] = 7
    m["k2"] = 13

    // Printing a map with e.g. `Println` will show all of
    // its key/value pairs.
    fmt.Println("map:", m)

    // Get a value for a key with `name[key]`.
    v1 := m["k1"]
    fmt.Println("v1: ", v1)

    // The builtin `len` returns the number of key/value
    // pairs when called on a map.
    fmt.Println("len:", len(m))

    // The builtin `delete` removes key/value pairs from
    // a map.
    delete(m, "k2")
    fmt.Println("map:", m)

    // The optional second return value when getting a
    // value from a map indicates if the key was present
    // in the map. This can be used to disambiguate
    // between missing keys and keys with zero values
    // like `0` or `""`. Here we didn't need the value
    // itself, so we ignored it with the _blank identifier_
    // `_`.
    _, prs := m["k2"]
    fmt.Println("prs:", prs)

    // You can also declare and initialize a new map in
    // the same line with this syntax.
    pepe := map[string]string{"nombre": "pepe", "apellido": "sánchez", "tel": "32234234234"}
    fmt.Println("map:", pepe)
}
