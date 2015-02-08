package main

import (
    "fmt"
    "time"
)

func main() {
    p := fmt.Printf
    now := time.Now()

    p("Hora %d\n", now.Hour())
    p("Minutos %d\n", now.Minute())
    p("Día de la semana %s\n", time.Weekday(now.Weekday()))
    p("Día %d\n", now.Day())
    p("Mes %d\n", now.Month())
    p("Año %d\n", now.Year())


    fmt.Println(now.Clock())

    then := time.Now()
    fmt.Println(now.After(then))

    const longForm = "Jan 2, 2006 at 3:04pm (CET)"
    date := fmt.Sprintf("%s %d, %d at %d:%d (CET)",  time.Month(now.Month()), now.Day(), now.Year(), 21, 45)
    p(date)
    //t, _ := time.Parse(longForm, date)
    //  fmt.Println(t.Local())


}
