package  main


import(
    "fmt"
    "unicode/utf8"
)

func main() {
    fmt.Println(12/5)
    fmt.Println(14/5)
    array := [5]int{10, 20, 30, 40, 50}
    fmt.Println("a[1] = ", array[1])

    for i, x := range array {
        fmt.Printf("array[%d] = %d \n", i, x)
    }


    str := "test 柳余"
    fmt.Println(str);
    fmt.Println("len(str)", len(str));
    fmt.Println("RuneCountInString", utf8.RuneCountInString(str));
    
    str2 := 'x' // print 120 rune
    fmt.Println(str2);

}

