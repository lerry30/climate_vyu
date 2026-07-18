# Preparing Data for Searching City Name

<mark>Note:</mark> This code only used <strong>once</strong> to extract city names from a json file and ordered them in alphabetical order. So, it can't be found for any feature/functionality in this backend.

## Structs

### Struct for City Data

This struct represents all the available city properties that can be used for future feature of this backend system. The only important in this case is the <strong>owm_city_name</strong> property.

```go
type Record struct {
	ID           string `json:"owm_city_id"`
	CityName     string `json:"owm_city_name"`
	Lat          string `json:"owm_latitude"`
	Long         string `json:"owm_longitude"`
	CountryShort string `json:"country_short"`
	Country      string `json:"country_long"`
	PostalCode   string `json:"postal_code"`
}
```

### List of City Data

Main struct.

```go
type Cities struct {
	Records []Record `json:"RECORDS"`
}
```

## File Path

The file path is declared at global scope.

```go
var filepath = "files/owm_city_list.json"
```

## Methods

<code>Load()</code> method reads the file content and parse it to json format.
- <code>os.ReadFile(file)</code> reads the file using <code>Open</code> method and handles the closing byitself. It returns <code>[]byte</code>(slice of bytes), raw bytes.
- <code>bytes.NewReader([]byte)</code> is just a wrapper that implements the <code>io.Reader</code> that <code>json.NewDecoder(io.Reader)</code> expects as its argument. This is because the <code>os.ReadFile(file)</code> method only returns slice of raw bytes.

```go
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
```

<code>SaveToCSV</code> method does all the work.
- Load the json file using <code>Load()</code> method above. The data will be saved on <code>Cities'</code> property(<code>[]Record</code>) from json format to struct properties.
- The second step extract the city names and saved them to a slice of string <code>var cities []string</code>.
- Using <code>slices.Sort(cities []string)</code> sorts the slice values alphabetically.
- <code>alphabet [26]string</code> is an array of characters in alphabet order to group the city names by their names.
- The city names will be grouped based on their first letter and saved to array of slice of letters <code>[26][]string{}</code>, 26 from a-z.
- Last steps are to join the city names in to a single long text based on their first character for csv file content which are separated with commas. So, all city names that starts with "a" will be joined together for a single file. This I think optimized the data of city names for searching.

```go
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
```