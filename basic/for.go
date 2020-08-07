package main

import("fmt")


func main() {
    // 指定声明数组
    arr := [10]int{7:7, 9:1000}
    fmt.Println(arr)

    // 类 c
    for i:=1; i<=3; i++ {
        fmt.Println("i = ", i)
    }

    fmt.Println("\n")

    // = while
    j := 1
    for true {
        if (j > 3) {
            break
        }

        fmt.Println("j = ", j)
        j++
    }

    fmt.Println("\n")

    // 数组
    array := [4]int{1, 34, 567, 900}
    for _, value := range array {
        fmt.Println("value:", value)
    }

    fmt.Println("\n")

    slice := []string{"Kobe", "AJ", "KG", "Tmac"}
    for _, value := range slice {
        fmt.Println("value:", value)
    }

    fmt.Printf("len(slice) %d - cap(slice) %d\n", len(slice), cap(slice))

    slice2 := slice[1:2:3]
    fmt.Printf("slice2 is %v\n", slice2)
    fmt.Printf("len(slice2) %d - cap(slice2) %d\n", len(slice2), cap(slice2))

    slice3 := make([]int, 3, 5)
    slice3[2] = 1222;
    fmt.Printf("slice3 is %v\n", slice3)

    //dict := make(map[string]string)

}