package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	mySlice, err := getInput()
	if err != nil {
		fmt.Printf("Error getting input %s \n", err)
		os.Exit(3)
	}

	greedyElf, mostCalories := findElfWithMostFood(mySlice)

	fmt.Printf("Elf number %d is carrying the most calories with %d calories\n", greedyElf, mostCalories)

	calorieTotal, e1, e2, e3 := findTop3Elves(mySlice)

	fmt.Printf("The top three elves are %d, %d and %d. They carry a combined %d calories\n", e1, e2, e3, calorieTotal)
}

func getInput() ([][]int, error) {
	file, _ := os.Open("Day1Input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var mySlice [][]int
	var innerSlice []int
	for scanner.Scan() {
		if scanner.Text() == "" {
			mySlice = append(mySlice, innerSlice)
			innerSlice = nil
		} else {
			calorieVal, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Unable to parse text input to integer")
				return nil, err
			}
			innerSlice = append(innerSlice, calorieVal)
		}
	}
	return mySlice, nil
}

func findElfWithMostFood(mySlice [][]int) (int, int) {
	mostCalories := 0
	greedyElf := 0

	for elf, foods := range mySlice {
		currentCaloriesTotal := 0
		for _, calories := range foods {
			currentCaloriesTotal += calories
		}
		if currentCaloriesTotal > mostCalories {
			mostCalories = currentCaloriesTotal
			greedyElf = elf + 1
		}
	}

	return greedyElf, mostCalories
}

func findTop3Elves(mySlice [][]int) (int, int, int, int) {
	caloriesTop := 0
	calories2nd := 0
	calories3rd := 0
	elf1 := 0
	elf2 := 0
	elf3 := 0

	for elf, foods := range mySlice {
		currentCaloriesTotal := 0
		for _, calories := range foods {
			currentCaloriesTotal += calories
		}
		//fmt.Printf("Elf %d has %d Calories\n", elf+1, currentCaloriesTotal)

		if currentCaloriesTotal > caloriesTop {
			if caloriesTop > calories2nd {
				if calories2nd > calories3rd {
					calories3rd = calories2nd
					elf2 = elf3
				}
				calories2nd = caloriesTop
				elf2 = elf1
			}
			caloriesTop = currentCaloriesTotal
			elf1 = elf + 1
		} else if currentCaloriesTotal > calories2nd {
			if calories2nd > calories3rd {
				calories3rd = calories2nd
				elf2 = elf3
			}
			calories2nd = currentCaloriesTotal
			elf2 = elf + 1
		} else if currentCaloriesTotal > calories3rd {
			calories3rd = currentCaloriesTotal
			elf3 = elf + 1
		}
	}

	//fmt.Printf("%d, %d, %d, = %d\n", caloriesTop, calories2nd, calories3rd, caloriesTop+calories2nd+calories3rd)
	return caloriesTop + calories2nd + calories3rd, elf1, elf2, elf3
}
