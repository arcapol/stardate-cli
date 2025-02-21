package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// Constants for default values and config filename.
const defaultBaseYear = 2323
const configFileName = ".stardate-cli-config"

// getPersistentBaseYear reads the base year from the config file in the user's home directory.
// If not found or any error occurs, it returns the defaultBaseYear.
func getPersistentBaseYear() int {
	home, err := os.UserHomeDir()
	if err != nil {
		return defaultBaseYear
	}
	configPath := filepath.Join(home, configFileName)
	data, err := os.ReadFile(configPath)
	if err != nil {
		return defaultBaseYear
	}
	year, err := strconv.Atoi(strings.TrimSpace(string(data)))
	if err != nil {
		return defaultBaseYear
	}
	return year
}

// setPersistentBaseYear writes the new base year to the config file.
func setPersistentBaseYear(newYear int) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	configPath := filepath.Join(home, configFileName)
	return os.WriteFile(configPath, []byte(strconv.Itoa(newYear)), 0644)
}

// parseDate converts a string in DD-MM-YYYY format to a time.Time object.
func parseDate(dateStr string) (time.Time, error) {
	parts := strings.Split(dateStr, "-")
	if len(parts) != 3 {
		return time.Time{}, fmt.Errorf("invalid date format, expected DD-MM-YYYY")
	}
	day, err1 := strconv.Atoi(parts[0])
	month, err2 := strconv.Atoi(parts[1])
	year, err3 := strconv.Atoi(parts[2])
	if err1 != nil || err2 != nil || err3 != nil {
		return time.Time{}, fmt.Errorf("invalid date, must be numbers in DD-MM-YYYY format")
	}
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local), nil
}

// calculateStardate computes the stardate for a given time and base year.
// Formula: Stardate = 1000*(Year - BaseYear) + (DayOfYear/TotalDays)*1000
func calculateStardate(t time.Time, baseYear int) float64 {
	year := t.Year()
	dayOfYear := t.YearDay()
	totalDays := 365
	if isLeapYear(year) {
		totalDays = 366
	}
	return 1000*float64(year-baseYear) + (float64(dayOfYear)/float64(totalDays))*1000
}

// stardateToDate converts a stardate value to a human date (time.Time) given a base year.
// It reverses the formula used in calculateStardate.
func stardateToDate(sd float64, baseYear int) time.Time {
	yearOffset := int(sd) / 1000
	year := baseYear + yearOffset
	fraction := sd - float64(yearOffset*1000)
	totalDays := 365
	if isLeapYear(year) {
		totalDays = 366
	}
	dayOfYear := int((fraction/1000)*float64(totalDays) + 0.5)
	return time.Date(year, 1, 1, 0, 0, 0, 0, time.Local).AddDate(0, 0, dayOfYear-1)
}

// isLeapYear determines if a year is a leap year.
func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}

func main() {
	// Define flags with both long and short versions.
	var dateStr string
	flag.StringVar(&dateStr, "date", "", "Human date in DD-MM-YYYY format to convert to stardate (defaults to current date).")
	flag.StringVar(&dateStr, "d", "", "Human date (shorthand).")

	var stardateValue float64
	flag.Float64Var(&stardateValue, "stardate", -1, "Stardate value to convert to human date.")
	flag.Float64Var(&stardateValue, "s", -1, "Stardate value (shorthand).")

	var baseValue int
	flag.IntVar(&baseValue, "base", 0, "Temporary base year for this conversion only (does not update persistent configuration).")
	flag.IntVar(&baseValue, "b", 0, "Temporary base year (shorthand).")

	var setBaseValue int
	flag.IntVar(&setBaseValue, "set-base", 0, "Set and persist a new base year for all future conversions.")

	var showBaseFlag bool
	flag.BoolVar(&showBaseFlag, "show-base", false, "Display the current persistent base year.")

	// Help flag: note that -h is already provided by flag package, but we alias --help as well.
	var helpFlag bool
	flag.BoolVar(&helpFlag, "help", false, "Display help information.")
	// The -h flag is already reserved for help by the flag package.

	// Override the default Usage function to include our custom message.
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: stardate [options]")
		fmt.Fprintln(os.Stderr, "For detailed help, run: stardate -h or --help")
	}

	// Parse the flags.
	flag.Parse()	

	// The flag package will automatically exit with an error if an undefined flag is passed,
	// so our custom usage message will be shown in those cases.

	// Load persistent base year from config.
	persistentBase := getPersistentBaseYear()

	// If no flags are passed, display the persistent base year, current date, and a hint to use help.
	if flag.NFlag() == 0 {
		currentDate := time.Now().Local()
		currentStardate := calculateStardate(currentDate, persistentBase)
		fmt.Printf("Current Date: %s\n", currentDate.Format("02-01-2006"))
		fmt.Printf("Current Stardate (using base year %d): %.2f\n", persistentBase, currentStardate)
		fmt.Println("\nFor more details on available commands and usage, run:")
		fmt.Println("  stardate -h or --help")
		os.Exit(0)
	}

	// If help flag is provided, display full help and exit.
	if helpFlag {
		fmt.Println("Usage: stardate [options]")
		flag.PrintDefaults()
		fmt.Println("\nExamples:")
		fmt.Println("  Convert current date to stardate:")
		fmt.Println("    stardate -date 21-02-2025")
		fmt.Println("  Convert a specific date to stardate using a temporary base year:")
		fmt.Println("    stardate -date 21-02-2025 -base 2300")
		fmt.Println("  Convert a stardate to human date:")
		fmt.Println("    stardate -stardate 45000")
		fmt.Println("  Update the reference base year:")
		fmt.Println("    stardate -set-base 2300")
		fmt.Println("  Show the current referece base year:")
		fmt.Println("    stardate -show-base")
		os.Exit(0)
	}

	// If show-base flag is used, display and exit.
	if showBaseFlag {
		fmt.Printf("Current Reference base year: %d\n", persistentBase)
		os.Exit(0)
	}

	// If set-base flag is used, update the persistent base year.
	if setBaseValue != 0 {
		if err := setPersistentBaseYear(setBaseValue); err != nil {
			fmt.Printf("Error setting reference base year: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Reference base year updated to %d\n", setBaseValue)
		// Use the updated base year for further conversion in this run.
		persistentBase = setBaseValue
	}

	// Determine which base year to use for this conversion.
	// If temporary -base flag is provided (non-zero), use it; otherwise, use the persistent base year.
	var baseYear int
	if baseValue != 0 {
		baseYear = baseValue
	} else {
		baseYear = persistentBase
	}

	// If stardate flag is provided (>= 0), perform stardate -> human date conversion.
	if stardateValue >= 0 {
		humanDate := stardateToDate(stardateValue, baseYear)
		fmt.Printf("Converted stardate %.2f to human date: %s (using base year %d)\n",
			stardateValue, humanDate.Format("02-01-2006"), baseYear)
		os.Exit(0)
	}

	// Otherwise, convert a human date to stardate.
	// Use current date if no date flag is provided.
	var dateToConvert time.Time
	var err error
	if dateStr == "" {
		dateToConvert = time.Now().Local()
	} else {
		dateToConvert, err = parseDate(dateStr)
		if err != nil {
			fmt.Printf("Error parsing date: %v\n", err)
			os.Exit(1)
		}
	}
	calculatedStardate := calculateStardate(dateToConvert, baseYear)
	fmt.Printf("Converted date %s to stardate: %.2f (using base year %d)\n",
		dateToConvert.Format("02-01-2006"), calculatedStardate, baseYear)
}