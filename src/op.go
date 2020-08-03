package main

import "fmt"

func main() {
    n1 := 0110
    n2 := 1011

    fmt.Printf("%v-%[1]T\n", n1)
    fmt.Printf("%v-%[1]T\n", n2)
    fmt.Printf("%b\n", n2)
    fmt.Printf("%v-%[1]T\n", 0110 &^ 1011)
    fmt.Printf("\n")
    fmt.Printf("\n")
    fmt.Printf("%d\n", 0110 &^ 1011)
    fmt.Printf("%0b\n", 0b0110 &^ 0b1011)
    fmt.Printf("%o\n", 0110 &^ 1011)
    /*
    a := 0b0110
    b := 0b0100
    c := a &^ b 
    fmt.Printf("%b\n", a)
    fmt.Printf("%d\n", a)
    fmt.Printf("%b\n", b)
    fmt.Printf("%d\n", b)
    fmt.Printf("%b\n", c)
    fmt.Printf("%d\n", c)
    fmt.Printf("\n")
    fmt.Printf("\n")
    fmt.Printf("%d\n", 0110 &^ 1011)
    fmt.Printf("%0b\n", 0b0110 &^ 0b1011)
    fmt.Printf("%o\n", 0110 &^ 1011)
    fmt.Printf("\n")
    fmt.Printf("\n")
    fmt.Printf("%d\n", 0110)
    fmt.Printf("%16b\n", 0110)
    fmt.Printf("%16b\n", 1011)
    */
}