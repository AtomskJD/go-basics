package main

import "fmt"

func main() {
	fmt.Println(firstBadVersion(2))
}
func isBadVersion(i int) bool {
	if i >= 2 {
		return true
	}
	return false
}
func firstBadVersion(n int) int {
	for ; n > 0; n-- {
		if !isBadVersion(n) {
			break
		}
	}
	return n + 1
}
