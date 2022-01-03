package main

import "fmt"

func main() {
	groupCity := map[int][]string{
		10:   []string{"Moscaov", "Klinton", "kirovsk"},
		100:  []string{"CPB", "NVZ", "Briansk"},
		1000: []string{"KIpr", "Harkov", "Novosibirsk"},
	}

	cityPopulation := map[string]int{
		"Moscaov":     54,
		"Klinton":     65,
		"kirovsk":     789,
		"CPB":         345,
		"NVZ":         548,
		"Briansk":     579,
		"KIpr":        65468,
		"Harkov":      78645,
		"Novosibirsk": 57498,
	}

	for _, city := range append(groupCity[10], groupCity[1000]...) {
		if _, ok := cityPopulation[city]; ok {
			delete(cityPopulation, city)
		}
	}

	fmt.Println(cityPopulation)
}
