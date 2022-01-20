package main

func finallyAsembled() (speech string) {
	// From the very beginning there are no words at all
	speech = ""
	for _, word := range words {
		// Each iteration you should
		speech += word // Concatanate "word" into "speech" variable
	}
	return
}
