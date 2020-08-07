package main

import(
    "fmt"
    "math/rand"
)

func main() {
    capstone5()
}

func c2() {
    var speed = 56000000/28/24
    fmt.Println(speed)
}

func c3() {
    var scene = "xxooyu"
    if scene == "hole" {
        fmt.Println("hole")
    } else if scene == "cave" {
        fmt.Println("cave")
    } else {
        fmt.Println("nothing")
    }

    switch(scene) {
    case "hole", "cave":
        fmt.Println("bingo")
    case "test":
        fmt.Println("test")
        // 下降至下一分支 必须显示的下降，所以没有 break
        fallthrough
    case "xxooyu":
        fmt.Println("liuduoyu")
    }

}


func getList() {
    arr := [5][4]string{
        {"zhangs a", "23", "round", "96"},
        {"space b", "13", "single", "58"},
        {"space b", "38", "round", "98"},
        {"space XX", "27", "round", "91"},
        {"space XX", "17", "signle", "51"},
    }

    fmt.Printf("%-15v d%-15v %-15v $%-15v\n", "company", "days", "type", "price")
    for _, a := range arr {
        fmt.Printf("%-15v d%-15v %-15v $%-15v\n", a[0], a[1], a[2], a[3])
    }
}

const secondsPerDay = 86400

func capstone5() {
    distance := 62100000
    company := ""
    trip := ""

    fmt.Println("")
    rate := 16 + rand.Intn(15)
}
