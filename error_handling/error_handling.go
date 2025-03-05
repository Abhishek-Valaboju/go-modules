package main

import "fmt"

type ValidationError struct {
	msg string
}

func (v *ValidationError) Error() string {
	return v.msg
}

func authentication(user_id, password string) error {

	defer fmt.Println("closing a database connection")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recoverd from ", r)
		}
	}()

	if user_id == "" {
		return &ValidationError{msg: "user_id id empty"}
	}
	if password == "" {
		return &ValidationError{msg: "password id empty"}
	}
	if user_id == "panicUser" {
		panic("critical failure ")
	}
	fmt.Println("User authentication successfull")
	return nil
}
func main() {
	user_id := "panicUser"
	password := "as"
	err := authentication(user_id, password)
	if err != nil {
		fmt.Println("error : ", err)
	}
}
