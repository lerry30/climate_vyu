# City Name Search Suggestions

For searching string from CSV file and to build city name suggestions

## Structs

Main struct has property of <code>Name</code> which will hold user searched string.

```go
type City struct {
	Name string `json:"city_name"`
}
```

Struct that will hold searched suggestions

```go
type FilteredCities struct {
	Data []string `json:"data"`
}
```

## Method

Indexing string like <code>str[0]</code> gives you a <code>byte(uint8)</code> not a character. So, a <code>byte</code> can perform checks like <code>str[0] >= 'a' && str[0] <= 'z'</code>, but it would break on multi-byte UTF-8 characters(accented letters, non-Latin scripts, etc.). If c.Name could contain not-ASCII names, so first convert it to runes.

Why <code>len(runes)</code> matters: indexing <code>str[0]</code> on an empty string panics, so guard it first.

For more robust checks, use <code>unicode.IsLetter + unicode.IsLower</code>.

### <code>Filter()</code> Execution Steps

- Extract the first letter.
- Getting the first letter helps to find the right CSV file.
- The last step is to get all the city names that starts with the searched string.

```go
func (c *City) Filter() (FilteredCities, error) {
	str := strings.ToLower(c.Name)

	runes := []rune(str)
	var prefix string
	
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
```