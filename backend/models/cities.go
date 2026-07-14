package models

import (
	"fmt"
	"os"
	"bytes"
	"strings"
	"slices"

	"encoding/json"

	"backend/utils"
)

type Record struct {
	ID string `json:"owm_city_id"`
	CityName string `json:"owm_city_name"`
	Lat string `json:"owm_latitude"`
	Long string `json:"owm_longitude"`
	CountryShort string `json:"country_short"`
	Country string `json:"country_long"`
	PostalCode string `json:"postal_code"`
}

type Cities struct {
	Records []Record `json:"RECORDS"`
}

var filepath = "files/owm_city_list.json"

func (c *Cities) Load() error {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("Error: %w\n", err)
	}

	reader := bytes.NewReader(data)
	decoder := json.NewDecoder(reader)

	if err := decoder.Decode(c); err != nil {
		return fmt.Errorf("Error: %w\n", err)
	}

	return nil
}

func (c *Cities) SaveToCSV() {
	if err := c.Load(); err != nil {
		fmt.Fprintf(os.Stderr, "Error %w\n", err)
	}

	var cities []string
	for _, c := range c.Records {
		name := strings.ToLower(strings.TrimSpace(c.CityName))
		if len(name) == 0 {
			continue
		}

		cities = append(cities, name)
	}

	slices.Sort(cities)
	alphabet := [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	sortedCities := [26][]string{}

	for i, c := range cities {
		for j, a := range alphabet {
			if strings.HasPrefix(c, a) {
				sortedCities[j] = append(sortedCities[j], cities[i])
				break
			}
		}
	}

	//fmt.Printf("%#v\n", sortedCities)

	for i, s := range sortedCities {
		cityNames := strings.Join(s, ",")
		fileName := "cities/" + alphabet[i] + "_city.csv"
		if err := utils.WriteFile(cityNames, fileName); err != nil {
			fmt.Fprintf(os.Stderr, "%w", err)
		}
	}

	/*allCities := strings.Join(cities, ",")
	if err := utils.WriteFile(allCities, "cities/cities.csv"); err != nil {
		fmt.Fprintf(os.Stderr, "%w", err)
	}*/
}