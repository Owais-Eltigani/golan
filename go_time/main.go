package main

import (
	"fmt"
	"time"
)


func main()  {
	
	fmt.Println("go time: ")

	presentTime := time.Now()

	// the default time format is 01-02-2006 Monday 15:04:05 MST
	fmt.Println(presentTime.Format("01-02-2006"))

	//? how to create a time in future or past ?
	createTime := time.Date(2026, time.July, 14,12,12,12,12,time.UTC)

	fmt.Println("the time is", createTime.Format("01-02-2006 Monday"))
}