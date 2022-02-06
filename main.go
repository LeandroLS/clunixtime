package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Insert the date in this format: YYYY-MM-DD hh:mm. Example: %s \n", time.Now().Format("2006-02-01 15:04"))
	fmt.Print("Insert the date: ")
	var date string
	var hoursAndMinutes string
	fmt.Scanln(&date, &hoursAndMinutes)
	timeParsed, _ := time.Parse("2006-02-01 15:04", date+" "+hoursAndMinutes)
	fmt.Println("Seconds:", timeParsed.Unix())
	fmt.Println("Milliseconds:", timeParsed.UnixMilli())
	fmt.Println("Nanoseconds:", timeParsed.UnixNano())
}
