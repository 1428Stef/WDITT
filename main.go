package main

import (
	"fmt"
	"time"
)

func main() {
	var today time.Time = time.Now()
	var day string = today.Weekday().String()

	fmt.Printf("Wow, it's %s today.\n", day)
	fmt.Println("Be sure to write it down so you don't forget.")
}
