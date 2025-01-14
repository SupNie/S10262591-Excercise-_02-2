package main

import (
	"fmt"
)

func main() {
	// Define the maps for rainfall data
	yr2017 := map[string][2]int{
		"Jan": {220, 18},
		"Feb": {10, 2},
		"Mar": {20, 3},
		"Apr": {10, 3},
		"May": {22, 6},
		"Jun": {24, 4},
		"Jul": {168, 13},
		"Aug": {198, 15},
		"Sep": {188, 15},
		"Oct": {213, 18},
		"Nov": {260, 19},
		"Dec": {278, 19},
	}

	yr2018 := map[string][2]int{
		"Jan": {245, 18},
		"Feb": {41, 6},
		"Mar": {201, 13},
		"Apr": {197, 16},
		"May": {175, 15},
		"Jun": {172, 14},
		"Jul": {166, 14},
		"Aug": {198, 18},
		"Sep": {185, 15},
		"Oct": {216, 16},
		"Nov": {256, 19},
		"Dec": {278, 20},
	}

	yr2019 := map[string][2]int{
		"Jan": {250, 17},
		"Feb": {40, 7},
		"Mar": {200, 12},
		"Apr": {195, 15},
		"May": {175, 15},
		"Jun": {170, 13},
		"Jul": {165, 13},
		"Aug": {190, 15},
		"Sep": {180, 14},
		"Oct": {210, 16},
		"Nov": {250, 19},
		"Dec": {275, 19},
	}

	// Store all the years' data in a map
	rainfallData := map[string]map[string][2]int{
		"2017": yr2017,
		"2018": yr2018,
		"2019": yr2019,
	}

	// Part (a): Update total number of days of rainfall for each year
	rainfallYearList := map[string]int{
		"2017": 0,
		"2018": 0,
		"2019": 0,
	}

	// Calculate the total days of rain for each year
	for year, months := range rainfallData {
		totalDays := 0
		for _, data := range months {
			totalDays += data[1]
		}
		rainfallYearList[year] = totalDays
	}

	// Display the updated total number of days of rainfall for each year
	fmt.Println("Total number of days of rainfall for each year:", rainfallYearList)

	// Part (b): Find the month with the lowest rainfall for each year
	rainfallLowestList := map[string]string{
		"2017": "",
		"2018": "",
		"2019": "",
	}

	// Find the month with the lowest rainfall for each year
	for year, months := range rainfallData {
		lowestRainfall := int(^uint(0) >> 1) // Set to maximum int value
		lowestMonth := ""
		for month, data := range months {
			if data[0] < lowestRainfall {
				lowestRainfall = data[0]
				lowestMonth = month
			}
		}
		rainfallLowestList[year] = lowestMonth
	}

	// Display the updated year and month in which there was the lowest rainfall
	fmt.Println("Month in which there was lowest rainfall for each year:", rainfallLowestList)
	fmt.Println("\nProgram exited.")
}
