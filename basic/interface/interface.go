package main

import (
	"log"
)

// https://colobu.com/2015/09/23/laws-of-goang-reflection/
type User struct {
	Name  string
	Email string
}

type Admin struct {
	User
	Level string
}

func (u *User) Notify() error {
	log.Printf("User: Sending User Email To %s<%s>\n",
		u.Name,
		u.Email)

	return nil
}

func (a *Admin) Notify() error {
	log.Printf("Admin: Sending Admin Email To %s<%s>\n",
		a.Name,
		a.Email)

	return nil
}

type Notifier interface {
	Notify() error
}

func SendNotification(notify Notifier) error {
	return notify.Notify()
}

func main() {
	admin := &Admin{
		User: User{
			"张三",
			"1@qq.com",
		},
		Level: "admin",
	}
	admin.Notify()
	admin.User.Notify()
	SendNotification(admin)
	SendNotification(&admin.User)
}
