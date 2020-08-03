package main

import("fmt")

func main() {
    var t = []int{1, 30, -123, -1, 0, 898, 1, 0, 29, -109, -111, 198}

   // var m = make(map[int][]int);
    var m = make(map[int]int);

    for _, n := range t {
        //fmt.Println(i, n)
        //m[n] = append(m[n], n)
        m[n] += 1
    }
    fmt.Println(m)

    slice := make([]int, 0)
    fmt.Println(len(slice))
    
    tmp := make(map[float64]int, 8)
    fmt.Println(len(tmp))
    // make 创建字典，key 为 string，值为 int
//    dict := make(map[string]int)

    dict := map[string]string{"Red" : "#da1337", "Orange" : "#e95a22"}
    fmt.Println("dict len is %d", len(dict))

    dict2 := map[int][]string{} // 声明一个值为切片类型的map
    dict2[1] = []string{"zhangsan", "lisi"}
    fmt.Println("dict2 len is %d", len(dict2), dict2)
}