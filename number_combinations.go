package main

import (
	"fmt"
	"os"
	"strconv"
)

// createIntCombinations generates all combinations of numbers of a given length
// from a given set of numbers.
//
// dirt: glitch/travels/labnotes/TOPSECRT/
func createIntCombinations(srcFilePath string, numbers []int, length int) {
	// TODO:
	//  - add logic to call this if the `1combinations.txt` file doesn't exist or
	//    is empty.
	//  - limit combinations to including each number 1+ times.

	// generate combinations of numbers and save them to a file
	var combinations [][]int
	generateMinimumCombinations(numbers, length, nil, &combinations)

	err := saveCombinationsToFile(srcFilePath, "", combinations)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Combinations saved to %s", srcFilePath)
}

// generateMaximumCombinations generates all possible combinations of
// numbers of the given length using the given set of numbers.
//
// This will generate a much larger set of possible combinations as it will
// include combinations that do not contain at least one of each int in numbers.
// ie:
// when numbers = [3, 4, 5, 7] & length = 5, this will generate 1024 results;
// some examples:
//   - Valid: [3, 3, 3, 3, 3]
//   - Valid: [3, 4, 5, 7, 3]
func generateMaximumCombinations(numbers []int, length int, prefix []int, result *[][]int) {
	if length == 0 {
		*result = append(*result, append([]int(nil), prefix...))
		return
	}
	for i := 0; i < len(numbers); i++ {
		newPrefix := append(prefix, numbers[i])
		generateMaximumCombinations(numbers, length-1, newPrefix, result)
	}
}

// generateMinimumCombinations generates a minimum list of combinations of
// numbers of the given length using the given set of numbers.
//
// This will generate a much smaller set of possible combinations as it will
// only include combinations that contain at least one of each int in numbers.
// ie:
// when numbers = [3, 4, 5, 7] & length = 5, this will generate 240 results;
// some examples:
//   - Invalid: [3, 3, 3, 3, 3]
//   - Valid: [3, 4, 5, 7, 3]
func generateMinimumCombinations(numbers []int, length int, prefix []int, result *[][]int) {
	// Check if the generated combination contains at least one of each int in
	// numbers
	if length == 0 {
		// create a lookup map of numbers to check against
		lookup := map[int]struct{}{}
		for _, num := range numbers {
			lookup[num] = struct{}{}
		}
		// check if the generated combination contains at least one of each int in
		// numbers lookup map
		matches := []int{}
		for _, num := range prefix {
			if _, found := lookup[num]; found {
				matches = append(matches, num)
				// remove the number from the lookup map so we don't count it twice
				delete(lookup, num)
			}
		}
		// if the number of matches equals the number of numbers, then we have a
		// combination that contains at least one of each int in numbers
		if len(matches) == len(numbers) {
			// add the combination to the result
			*result = append(*result, append([]int(nil), prefix...))
		}
		return
	}
	// generate all available combinations
	for i := 0; i < len(numbers); i++ {
		newPrefix := append(prefix, numbers[i])
		generateMinimumCombinations(numbers, length-1, newPrefix, result)
	}
}

// saveCombinationsToFile saves all combinations to a file.
func saveCombinationsToFile(filename string, separator string, combinations [][]int) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	for _, combo := range combinations {
		combinationStr := ""
		for _, num := range combo {
			combinationStr += strconv.Itoa(num) + separator
		}
		if separator != "" {
			combinationStr = combinationStr[:len(combinationStr)-1] // Remove the last separator
		}
		_, err := fmt.Fprintln(file, combinationStr)
		if err != nil {
			return err
		}
	}
	return nil
}
