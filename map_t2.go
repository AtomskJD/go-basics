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

	tmp := make(map[string]int)

	for _, city := range groupCity[100] {
		if p, ok := cityPopulation[city]; ok {
			tmp[city] = p
		}
	}
	cityPopulation = tmp

	fmt.Println(cityPopulation)
}
