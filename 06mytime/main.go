package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("--------> My TIMe <--------")

	currentTime := time.Now()
	fmt.Printf("Current time is: Type: %T || Value: %v \n", currentTime, currentTime)

	formattedTime := currentTime.Format("02-01-2006 15:04:15")
	fmt.Printf("Formatted time is: Type: %T || Value: %v \n", formattedTime, formattedTime)

	customDate := time.Date(2020, time.January, 11, 14, 50, 20, 0, time.UTC)
	formattedCustomDate := customDate.Format("02-01-2006 15:04:15")
	fmt.Printf("Created time is: Type: %T || Value: %v \n", customDate, customDate)
	fmt.Printf("Formatted Created time is: Type: %T || Value: %v \n", formattedCustomDate, formattedCustomDate)
}
