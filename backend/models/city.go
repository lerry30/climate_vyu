package models

import (
	"fmt"
	"strings"

	"backend/utils"
)

// main
type City struct {
	Name string `json:"city_name"`
}

type FilteredCities struct {
	Data []string `json:"data"`
}

// ---
// methods

func (c *City) Filter() (FilteredCities, error) {
	alphabet := [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	var prefix string

	str := strings.ToLower(c.Name)

	for _, a := range alphabet {
		if strings.HasPrefix(str, a) {
			prefix = a
			break
		}
	}

	if prefix == "" {
		return FilteredCities{}, fmt.Errorf("City name not found.")
	}

	filepath := "cities/" + prefix + "_city.csv"

	row, err := utils.ReadCSV(filepath)
	if err != nil {
		return FilteredCities{}, fmt.Errorf("%w", err)
	}

	var matchedStr FilteredCities
	for _, r := range row {
		for _, d := range r {
			if strings.HasPrefix(d, str) {
				matchedStr.Data = append(matchedStr.Data, d)
			}
		}
	}

	return matchedStr, nil
}