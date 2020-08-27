package main
import "fmt"
func main() {
    rates:=[]int32{1,2,3,4,5,6}
    for star, rate := range rates {
        if star+2 < 1 {
            panic("")
        }
        fmt.Println(star, rate)
    }
}
