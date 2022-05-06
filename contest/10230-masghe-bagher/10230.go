package main

import "fmt"

func main() {
	var triangleAngles [3]int

	fmt.Scanf("%d%d%d", &triangleAngles[0], &triangleAngles[1], &triangleAngles[2])
	if isTriangle(triangleAngles) == true{
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}

func isTriangle(angleArray [3]int) bool{
	const sumOfAngles int = 180
	
	if angleArray[0] == 0 || angleArray[1] == 0 || angleArray[2] == 0 {
		return false
	} else if angleArray[0] + angleArray[1] + angleArray[2] != sumOfAngles {
		return false
	} else {
		return true
	}
}
