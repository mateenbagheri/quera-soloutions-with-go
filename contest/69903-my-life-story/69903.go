package main

import "fmt"

type date struct {
	Day   int
	Month int
}

func dateDiff(firstDate date, secondDate date) int {
	endDate, stDate := sortDate(firstDate, secondDate)
	stDateDays := monthToDay(stDate.Month) + stDate.Day
	endDateDays := monthToDay(endDate.Month) + endDate.Day
	return endDateDays - stDateDays
}

func sortDate(firstDate date, secondDate date) (date, date) {
	if firstDate.Month > secondDate.Month {
		return firstDate, secondDate
	} else if secondDate.Month > firstDate.Month {
		return secondDate, firstDate
	} else if firstDate.Day > secondDate.Day {
		return firstDate, secondDate
	} else {
		return secondDate, firstDate
	}
}

func monthToDay(monthCount int) int {
	var result int
	if 1 <= monthCount && monthCount <= 12 {
		if monthCount <= 7 {
			result = (monthCount - 1) * 31
		} else if monthCount <= 11 {
			result = (6 * 31) + ((monthCount - 6) * 30)
		} else {
			result = (6 * 31) + (5 * 30)
		}
	}
	return result
}

func main() {
	var firstMonth, firstDay, secondMonth, secondDay int
	fmt.Scanf("%d %d\n", &firstMonth, &firstDay)
	fmt.Scanf("%d %d\n", &secondMonth, &secondDay)

	firstDate := date{firstDay, firstMonth}
	secondDate := date{secondDay, secondMonth}
	fmt.Println(dateDiff(firstDate, secondDate))
}
