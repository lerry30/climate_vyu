package utils

import (
	"encoding/csv"
	"fmt"
	"os"
)

func WriteFile(data string, filename string) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("Error: %w\n", err)
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		return fmt.Errorf("Error: %w\n", err)
	}
	return nil
}

func ReadCSV(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return [][]string{}, fmt.Errorf("%w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	row, err := reader.ReadAll() // return [][]string, error

	return row, err
}
