package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	var numOfChickens, numOfRoofCanCarry int
	fmt.Print("numOfChickens: ")
	fmt.Scanf("%d\n", &numOfChickens)
	checkInvalidOfInputNumer(numOfChickens, 1000000)
	fmt.Print("numOfRoofCanCarry: ")
	fmt.Scanf("%d\n", &numOfRoofCanCarry)
	checkInvalidOfInputNumer(numOfRoofCanCarry, 1000000)

	var positionOfChickens []int
	fmt.Println("positionOfChickens: ", positionOfChickens)
	for i := 0; i < numOfChickens; i++ {
		var position int
		fmt.Printf("position %d: ", (i + 1))
		fmt.Scanf("%d\n", &position)

		checkInvalidOfInputNumer(position, 1000000000)
		positionOfChickens = append(positionOfChickens, position)
	}

	var maxOfChicken int = getMaxNumOfChickenProtect(positionOfChickens, numOfRoofCanCarry)
	fmt.Println("Max Of Chickens: ", maxOfChicken)
}

func checkInvalidOfInputNumer(nubmerCheck int, limitOfLength int) {
	if nubmerCheck < 1 || nubmerCheck > limitOfLength {
		fmt.Println("Input out of length.")
		os.Exit(0)
	}
}

func getMaxNumOfChickenProtect(positionOfChickens []int, numOfRoofCanCarry int) int {
	sort.Ints(positionOfChickens)
	maxNumberAbleToProtect := 0

	for i := 0; i < len(positionOfChickens); i++ {
		count := 1
		for j := i + 1; j < len(positionOfChickens) && positionOfChickens[j] <= positionOfChickens[i]+numOfRoofCanCarry-1; j++ {
			count++
		}
		if count > maxNumberAbleToProtect {
			maxNumberAbleToProtect = count
		}
	}

	return maxNumberAbleToProtect
}
