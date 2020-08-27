package main

import(
    "fmt"
    "strconv"
)

func main() {
    n := int64(123)
    s := strconv.FormatInt(n, 10) // s == "61" (hexadecimal)

    fmt.Println(s)
    fmt.Printf("%v-%T", s, s)
}
