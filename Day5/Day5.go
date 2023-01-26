package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type step struct {
	moves, from, to int
}

func main() {
	myStacks := getInput()
	steps := getSteps()
	partOne(steps, myStacks)
	myStacks = getInput()
	PartTwo(steps, myStacks)

}

func getInput() [][]string {
	file, _ := os.Open("Day5Input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var myStacks [][]string
	for scanner.Scan() {
		currentStack := scanner.Text()
		currentStack = strings.Replace(currentStack, "[", "", -1)
		currentStack = strings.Replace(currentStack, "]", "", -1)
		currentStack = strings.Replace(currentStack, " ", "", -1)

		var stack []string
		for _, b := range []byte(currentStack) {
			stack = append(stack, string(b))
		}
		myStacks = append(myStacks, stack)
	}

	return myStacks
}

func getSteps() []step {
	var steps []step
	file, _ := os.Open("Day5Moves.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentMoveSet := scanner.Text()
		var s string
		var currentStep step
		_, err := fmt.Sscanf(currentMoveSet, "%s %d %s %d %s %d", &s, &currentStep.moves, &s, &currentStep.from, &s, &currentStep.to)
		if err != nil {
			fmt.Printf("Error parsing moves %s\n", err)
		}
		currentStep.from--
		currentStep.to--
		steps = append(steps, currentStep)
	}
	return steps
}

func partOne(steps []step, myStacks [][]string) {
	fmt.Print("Part 1: ")
	for _, step := range steps {
		for i := 0; i < step.moves; i++ {
			lastElementInFrom := myStacks[step.from][len(myStacks[step.from])-1]
			myStacks[step.from] = myStacks[step.from][:len(myStacks[step.from])-1]
			myStacks[step.to] = append(myStacks[step.to], lastElementInFrom)
		}
	}

	for _, stack := range myStacks {
		fmt.Printf("%s", stack[len(stack)-1])
	}
}

func PartTwo(steps []step, myStacks [][]string) {
	fmt.Print("\nPart 2: ")
	for _, step := range steps {
		for i := step.moves; i > 0; i-- {
			lastElementInFrom := myStacks[step.from][len(myStacks[step.from])-(i)]
			myStacks[step.to] = append(myStacks[step.to], lastElementInFrom)
		}
		myStacks[step.from] = myStacks[step.from][:len(myStacks[step.from])-(step.moves)]

	}

	for _, stack := range myStacks {
		fmt.Printf("%s", stack[len(stack)-1])
	}
}
