package main

import "fmt"

func main() {
	var trafficLightsInfo [][]int
	var numberOfTrafficLights, length int
	var distance, greenDuration, redDuration int
	fmt.Scanf("%d %d\n", &numberOfTrafficLights, &length)
	for i := 0; i < numberOfTrafficLights; i++ {
		var trafficLight []int
		fmt.Scanf("%d %d %d\n", &distance, &redDuration, &greenDuration)
		trafficLight = append(trafficLight, distance, redDuration, greenDuration)
		trafficLightsInfo = append(trafficLightsInfo, trafficLight)
	}
	fmt.Println(getBackTimeCalc(length, trafficLightsInfo))
}

func getBackTimeCalc(length int, trafficLights [][]int) int {
	var currentTime, currentLength int
	for currentLength < length {
		iso := false
		for _, trafficLight := range trafficLights {
			if trafficLight[0] == currentLength {
				if trafficLight[1] > currentTime%(trafficLight[1]+trafficLight[2]) {
					currentTime += trafficLight[1] - currentTime%(trafficLight[1]+trafficLight[2])
				}
			} else {
				iso = true
			}
		}
		if iso {
			currentTime += 1
		}
		currentLength += 1
	}

	return currentTime
}
