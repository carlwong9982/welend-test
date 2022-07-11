package main

import (
	"errors"
	"fmt"
)

func removeString(targets []any) []int {
	result := []int{}
	for _, target := range targets {
		intValue, isInt := target.(int)
		if isInt {
			result = append(result, intValue)
		}
	}
	return result
}

func transform(targets []any) ([]any, error) {
	results := []any{}
	if len(targets) == 0 {
		return results, nil
	}

	intValue, isInt := targets[0].(int)
	if !isInt || intValue < 0 {
		return results, errors.New("Size must be positive integer")
	}
	length := intValue
	pointer := 0
	targets = targets[1:]
	if length > len(targets) {
		return results, errors.New("Size can't bigger than array length")
	}

	for pointer <= length-1 {
		if pointer == 0 || pointer == length-1 {
			results = append(results, targets[pointer])
		} else {
			results = append(results, targets[length-1-pointer])
		}
		pointer++
	}

	return results, nil
}

func main() {
	question1 := removeString([]any{1, 3, 'a', 'b', 4, 7})
	fmt.Println(question1)
	question2, err := transform([]any{5, "a", "b", "c", "d", "e", "f", "g"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(question2)
}
