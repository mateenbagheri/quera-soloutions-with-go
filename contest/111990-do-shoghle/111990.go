package main

import "fmt"

const (
	zojDaysBank  = "perspolis"
	fardDaysBank = "bahman"
	tatilDay     = "tatil"
)

func getBank(day string) string {
	dayValueMap := map[string]int{
		"shanbe":       0,
		"yekshanbe":    1,
		"doshanbe":     2,
		"seshanbe":     3,
		"chaharshanbe": 4,
		"panjshanbe":   5,
		"jome":         6,
	}

	if dayValueMap[day] == 6 {
		return tatilDay
	} else if dayValueMap[day]%2 == 0 {
		return zojDaysBank
	} else {
		return fardDaysBank
	}
}

func main() {
	var day string
	fmt.Scanf("%s\n", &day)
	fmt.Println(getBank(day))
}
