package main

import (
	"fmt"
	"time"
)

func main() {
	// Get the current date
	currentDate := time.Now()

	// Get the first day of the current month
	firstDayOfMonth := time.Date(currentDate.Year(), currentDate.Month(), 1, 0, 0, 0, 0, time.Local)

	// Get the last day of the current month
	lastDayOfMonth := firstDayOfMonth.AddDate(0, 1, -1)

	// Get the number of days in the current month
	numDaysInMonth := lastDayOfMonth.Day()

	// Get the day of the week for the first day of the month
	firstDayOfWeek := firstDayOfMonth.Weekday()

	// Print the calendar header
	fmt.Printf("Calendar for %s\n", currentDate.Month())

	// Print the days of the week
	fmt.Println("Su Mo Tu We Th Fr Sa")

	// Print leading spaces for the first week
	for i := 0; i < int(firstDayOfWeek); i++ {
		fmt.Printf("   ")
	}

	// Print the days of the month
	day := 1
	for day <= numDaysInMonth {
		fmt.Printf("%2d ", day)

		// Move to the next line if it's Saturday
		if (firstDayOfWeek+time.Weekday(day-1))%7 == 6 {
			fmt.Println()
		}

		day++
	}

	fmt.Println()
}
