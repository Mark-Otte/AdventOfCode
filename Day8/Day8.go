package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	treeGrid := getInput()

	numVisibleTrees := calculateVisibleTrees(treeGrid)
	fmt.Printf("Number of Visible Trees: %d\n", numVisibleTrees)

	highestScenicScore := calculateHighestScenicScore(treeGrid)
	fmt.Printf("Highest Scenic Score: %d\n", highestScenicScore)
}

func getInput() [][]int {
	file, _ := os.Open("Day8Input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var treeGrid [][]int
	for scanner.Scan() {
		currentRow := scanner.Text()

		var stack []int
		stringStack := strings.Split(currentRow, "")
		for _, stringVal := range stringStack {
			intVal, err := strconv.Atoi(stringVal)
			if err != nil {
				fmt.Printf("Error parsing int from string: %v\n", err)
			}
			stack = append(stack, intVal)
		}
		treeGrid = append(treeGrid, stack)
	}

	return treeGrid
}

func calculateVisibleTrees(treeGrid [][]int) int {
	numVisibleTrees := 0

	for y, treeRow := range treeGrid {
		// If the below statement is true, then the tree is at the edge so it must be visible.
		if y == 0 || y == len(treeGrid)-1 {
			numVisibleTrees += len(treeRow)
			continue
		}
		for x, val := range treeRow {
			// If the below statement is true, then the tree is at the edge so it must be visible.
			if x == 0 || x == len(treeRow)-1 {
				numVisibleTrees += 1
				continue
			}
			if checkFromTop(treeGrid, y, x, val) {
				numVisibleTrees += 1
				continue
			}
			if checkFromBottom(treeGrid, y, x, val) {
				numVisibleTrees += 1
				continue
			}
			if checkFromLeft(treeGrid, y, x, val) {
				numVisibleTrees += 1
				continue
			}
			if checkFromRight(treeGrid, y, x, val) {
				numVisibleTrees += 1
				continue
			}
			continue
		}
	}
	return numVisibleTrees
}

func checkFromTop(treeGrid [][]int, y int, x int, val int) bool {
	for yy, treeRow := range treeGrid {
		if yy == y {
			return true
		}
		if treeRow[x] >= val {
			return false
		}
	}

	// This should never happen
	return false
}

func checkFromBottom(treeGrid [][]int, y int, x int, val int) bool {
	for i := len(treeGrid) - 1; i >= 0; i-- {
		treeRow := treeGrid[i]
		if i == y {
			return true
		}
		if treeRow[x] >= val {
			return false
		}
	}

	// This should never happen
	return false
}

func checkFromLeft(treeGrid [][]int, y int, x int, val int) bool {
	treeRow := treeGrid[y]
	for xx, treeVal := range treeRow {
		if xx == x {
			return true
		}
		if treeVal >= val {
			return false
		}
	}

	// This should never happen
	return false
}

func checkFromRight(treeGrid [][]int, y int, x int, val int) bool {
	treeRow := treeGrid[y]
	for i := len(treeRow) - 1; i >= 0; i-- {
		if i == x {
			return true
		}
		if treeRow[i] >= val {
			return false
		}
	}

	// This should never happen
	return false
}

func calculateHighestScenicScore(treeGrid [][]int) int {
	highestScenicScore := 0

	for y, treeRow := range treeGrid {
		// If the below statement is true, then the tree is at the edge so scenic score will be 0.
		if y == 0 || y == len(treeGrid)-1 {
			continue
		}
		for x := range treeRow {
			// If the below statement is true, then the tree is at the edge so scenic score will be 0.
			if x == 0 || x == len(treeRow)-1 {
				continue
			}
			upScore := lookUp(treeGrid, y, x)
			downScore := lookDown(treeGrid, y, x)
			leftScore := lookLeft(treeGrid, y, x)
			rightScore := lookRight(treeGrid, y, x)

			currentScenicScore := upScore * downScore * leftScore * rightScore
			if currentScenicScore > highestScenicScore {
				highestScenicScore = currentScenicScore
			}

		}
	}
	return highestScenicScore
}

func lookUp(treeGrid [][]int, y int, x int) int {
	scenicValue := 0
	treeHeight := treeGrid[y][x]
	for i := y - 1; i >= 0; i-- {
		if treeHeight > treeGrid[i][x] {
			scenicValue++
		} else {
			scenicValue++
			return scenicValue
		}
	}
	return scenicValue
}

func lookDown(treeGrid [][]int, y int, x int) int {
	scenicValue := 0
	treeHeight := treeGrid[y][x]
	for i := y + 1; i <= len(treeGrid)-1; i++ {
		if treeHeight > treeGrid[i][x] {
			scenicValue++
		} else {
			scenicValue++
			return scenicValue
		}
	}
	return scenicValue
}

func lookLeft(treeGrid [][]int, y int, x int) int {
	scenicValue := 0
	treeHeight := treeGrid[y][x]
	for i := x - 1; i >= 0; i-- {
		if treeHeight > treeGrid[y][i] {
			scenicValue++
		} else {
			scenicValue++
			return scenicValue
		}
	}
	return scenicValue
}

func lookRight(treeGrid [][]int, y int, x int) int {
	scenicValue := 0
	treeHeight := treeGrid[y][x]
	for i := x + 1; i <= len(treeGrid)-1; i++ {
		if treeHeight > treeGrid[y][i] {
			scenicValue++
		} else {
			scenicValue++
			return scenicValue
		}
	}
	return scenicValue
}
