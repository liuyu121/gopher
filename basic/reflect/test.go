package main

import (
	"fmt"
	"reflect"
)

const tagName = "validate"

type User struct {
	Id    int    `validate:"-"`
	Name  string `validate:"presence,min=2,max=32"`
	Email string `validate:"email,required"`
}

func main() {
	user := User{
		Id:    1,
		Name:  "liu yu",
		Email: "liuynope@gmail.com",
	}

	t := reflect.TypeOf(user)
	fmt.Println("TypeOf: ", t) // main.User
	t2 := reflect.ValueOf(user)
	fmt.Println("ValueOf", t2)                                     // {1 liu yu liuynope@gmail.com}
	fmt.Println("Interface-1: ", t2.Field(1).Interface())          // Interface-1:  liu yu
	fmt.Println("Interface-2: ", t2.Field(1).Interface().(string)) // Interface-2:  liu yu

	fmt.Println("Type: ", t.Name()) // User
	fmt.Println("Kind: ", t.Kind()) // struct

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("i=%v, field=%v\n", i, field)

		tag := field.Tag.Get(tagName)
		fmt.Printf("%d. %v(%v), tag:'%v'\n", i+1, field.Name, field.Type.Name(), tag)
	}
}
