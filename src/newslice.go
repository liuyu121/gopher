package main

import("fmt")

func main() {
    source := []string{"apple", "orange", "plum", "banana", "grape"}
    fmt.Println(source)

    slice := source[1:2:3]
    fmt.Println(slice)

    slice2 := source[1:2:2] // 容量 = 长度，则 append 后，会生成新的底层数组
    fmt.Println(slice2)

    slice = append(slice, "x1") // 这时候，容量用完了
    fmt.Println(source)
    fmt.Println(slice)
    slice = append(slice, "x2") // 下面开始，就是使用新的底层数组
    fmt.Println(source)
    fmt.Println(slice)
    slice = append(slice, "x3")
    fmt.Println(source)
    fmt.Println(slice)
    slice = append(slice, "x4")
    fmt.Println(source)
    fmt.Println(slice)
    slice = append(slice, "x5")
    fmt.Println(source)
    fmt.Println(slice)
    fmt.Println(slice2)
    /*
    slice2 = append(slice2, "x1y2")
    fmt.Println(source)
    fmt.Println(slice)
    fmt.Println(slice2)
    */
}