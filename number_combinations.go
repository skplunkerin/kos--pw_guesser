package main

import (
	"fmt"
	"os"
	"strconv"
)

// generateCombinations generates all combinations of numbers of a given length
// from a given set of numbers.
func generateCombinations(numbers []int, length int, prefix []int, result *[][]int) {
	if length == 0 {
		*result = append(*result, append([]int(nil), prefix...))
		return
	}

	for i := 0; i < len(numbers); i++ {
		newPrefix := append(prefix, numbers[i])
		generateCombinations(numbers, length-1, newPrefix, result)
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
