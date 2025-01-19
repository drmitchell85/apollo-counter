package cli

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
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

func apiHealth() ([]byte, error) {

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get("http://localhost:8080/health")
	if err != nil {
		return nil, fmt.Errorf("connection failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	return b, nil
}
