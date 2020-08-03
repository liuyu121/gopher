package main

import(
    "fmt"
    "counter"
    "user"
)

func main() {
    counter := counter.New(10)
    fmt.Printf("Counter: %d\n", counter)

    admin := user.Admin{
        //Name: "zhangsamn",
        //Email: "1@1.com",
        Rights: 123,
    }

    admin.Name = "kobe"
    admin.Email = "kobe@heaven.com"
    fmt.Printf("Admin is %v \n", admin)
    
}

