package cli

import (
	"fmt"
	"net/http"
	"strconv"
)

// reverse a given string
func reverse(input string) (result string) {
	for _, c := range input {
		result = string(c) + result
	}

	return result
}

// if digits is false, return length
// if digits is true, we inspect the number
func inspect(input string, digits bool) (count int, kind string) {
	if !digits {
		return len(input), "char"
	}

	return inspectNumbers(input), "digit"
}

func inspectNumbers(input string) (count int) {
	for _, c := range input {
		_, err := strconv.Atoi(string(c))
		if err == nil {
			count++
		}
	}
	return count
}

func apiHealth() error {

	resp, err := http.Get("localhost:8080/health")
	if err != nil {
		return fmt.Errorf("could not retrieve server health with error %s", err)
	}
	defer resp.Body.Close()

	return nil
}
