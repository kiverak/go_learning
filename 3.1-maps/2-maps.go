package maps

import "fmt"

func DoMap2() {
	groupCity := map[int][]string{
		10:   []string{"qwerty"},             // города с населением 10-99 тыс. человек
		100:  []string{"Engels", "Marks"},    // города с населением 100-999 тыс. человек
		1000: []string{"Moscow", "New-York"}, // города с населением 1000 тыс. человек и более
	}

	cityPopulation := map[string]int{
		"qwerty":   12345,
		"Engels":   12345,
		"Marks":    12345,
		"Moscow":   12345,
		"New-York": 12345,
	}

	for group, cities := range groupCity {
		if group != 100 {
			for _, city := range cities {
				delete(cityPopulation, city)
			}
		}
	}

	fmt.Println(cityPopulation)
}
