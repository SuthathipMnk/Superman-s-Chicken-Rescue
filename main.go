package main

import (
	"fmt"
	"sort"
)

func main() {
	var numOfChickens, numOfRoofCanCarry int
	fmt.Print("numOfChickens: ")
	fmt.Scanf("%d\n", &numOfChickens)
	fmt.Print("numOfRoofCanCarry: ")
	fmt.Scanf("%d\n", &numOfRoofCanCarry)

	var positionOfChickens []int
	fmt.Println("positionOfChickens: ", positionOfChickens)
	for i := 0; i < numOfChickens; i++ {
		var position int
		fmt.Printf("position %d: ", (i + 1))
		fmt.Scanf("%d\n", &position)
		positionOfChickens = append(positionOfChickens, position)
	}

	var maxOfChicken int = getMaxNumOfChickenProtect(positionOfChickens, numOfRoofCanCarry)
	fmt.Println("Max Of Chickens: ", maxOfChicken)
}

func getMaxNumOfChickenProtect(positionOfChickens []int, numOfRoofCanCarry int) int {
	sort.Ints(positionOfChickens)
	maxCount := 0

	for i := 0; i < len(positionOfChickens); i++ {
		count := 1
		for j := i + 1; j < len(positionOfChickens) && positionOfChickens[j] <= positionOfChickens[i]+numOfRoofCanCarry-1; j++ {
			count++
		}
		if count > maxCount {
			maxCount = count
		}
	}

	return maxCount
}
