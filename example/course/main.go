package main

import (
	"errors"
	"fmt"
	"os"

	couseSdk "github.com/ncostamagna/go_course_sdk/course"
)

func main() {
	courseTrans := couseSdk.NewHttpClient("http://localhost:8082", "")

	course, err := courseTrans.Get("e5a48a1c-5837-49c1-99e0-c51cb42ce41212")
	if err != nil {
		if errors.As(err, &couseSdk.ErrNotFound{}) {
			fmt.Println("Not found:", err.Error())
			os.Exit(1)
		}
		fmt.Println("Internal Server Error:", err.Error())
		os.Exit(1)
	}

	fmt.Println(course)

}
