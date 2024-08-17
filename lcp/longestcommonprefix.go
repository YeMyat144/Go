package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// Start with the first string as the prefix
	prefix := strs[0]

	// Compare the prefix with each string in the array
	for _, str := range strs[1:] {
		for len(str) < len(prefix) || prefix != str[:len(prefix)] {
			// Reduce the prefix by one character each time
			prefix = prefix[:len(prefix)-1]
			// If the prefix becomes empty, return an empty string
			if prefix == "" {
				return ""
			}
		}
	}

	return prefix
}

func main() {
	strs1 := []string{"flower", "flow", "flight"}
	strs2 := []string{"dog", "racecar", "car"}

	fmt.Println("Example 1:", longestCommonPrefix(strs1)) 
	fmt.Println("Example 2:", longestCommonPrefix(strs2)) 
}
