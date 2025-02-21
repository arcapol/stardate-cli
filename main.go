package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var referenceYear = 2025

// isLeapYear checks if a given year is a leap year
func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}

// CalculateStardate computes the Stardate based on the TNG formula
func CalculateStardate(year int, dayOfYear int) float64 {
	totalDays := 365
	if isLeapYear(year) {
		totalDays = 366
	}

	// Compute Stardate using the corrected formula
	stardate := 1000*float64(year-referenceYear) + (float64(dayOfYear)/float64(totalDays))*1000

	return stardate
}

// ParseDate converts DD-MM-YYYY format into year and day of the year
func ParseDate(dateStr string) (int, int, error) {
	parts := strings.Split(dateStr, "-")
	if len(parts) != 3 {
		return 0, 0, fmt.Errorf("invalid date format, expected DD-MM-YYYY")
	}

	day, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid date, must be numbers in DD-MM-YYYY format")
	}
	month, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid date, must be numbers in DD-MM-YYYY format")
	}
	year, err := strconv.Atoi(parts[2])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid date, must be numbers in DD-MM-YYYY format")
	}

	// Convert to a time.Time object
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	// Get the day of the year
	dayOfYear := date.YearDay()

	return year, dayOfYear, nil
}

func main() {
	// Define command-line flags
	dateFlag := flag.String("date", "", "Date in DD-MM-YYYY format (default: today)")

	flag.Parse()

	// Get the date input
	var year, dayOfYear int
	var err error

	if *dateFlag == "" {
		// Use today's date if no input is given
		now := time.Now()
		year = now.Year()
		dayOfYear = now.YearDay()
	} else {
		// Parse the input date
		year, dayOfYear, err = ParseDate(*dateFlag)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	}

	// Calculate and print Stardate
	stardate := CalculateStardate(year, dayOfYear)
	fmt.Printf("Stardate: %.2f\n", stardate)
}
