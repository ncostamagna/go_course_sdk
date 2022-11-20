package main

import (
	"errors"
	"fmt"
	"os"

	userSdk "github.com/ncostamagna/go_course_sdk/user"
)

func main() {
	userTrans := userSdk.NewHttpClient("http://localhost:8081", "")

	user, err := userTrans.Get("9699dd77-b7fb-40af-9d43123213213")
	if err != nil {
		if errors.As(err, &userSdk.ErrNotFound{}) {
			fmt.Println("Not found:", err.Error())
			os.Exit(1)
		}
		fmt.Println("Internal Server Error:", err.Error())
		os.Exit(1)
	}

	fmt.Println(user)
}
