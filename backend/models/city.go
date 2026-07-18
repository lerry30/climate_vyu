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
	str := strings.ToLower(c.Name)

	// Indexing string like str[0] gives you a byte(uint8) not a
	// character. So, a byte(str[0]) can perform checks like
	// str[0] >= 'a' && str[0] <= 'z', but it would break on multi-byte
	// UTF-8 characters(accented letters, non-Latin scripts, etc.). If
	// c.Name could contain not-ASCII names, so first convert it to runes.
	runes := []rune(str)
	var prefix string

	// Why len(runes) matters: indexing str[0] on an empty string
	// panics, so guard it first.
	// For more robust checks, use unicode.IsLetter + unicode.IsLower
	if len(runes) > 0 && runes[0] >= 'a' && runes[0] <= 'z' {
		prefix = string(runes[0])
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
