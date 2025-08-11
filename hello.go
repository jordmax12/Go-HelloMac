package main

import "fmt"

func checkIfPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	for left < right {
		fmt.Printf("%c %c\n", s[left], s[right])
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func reverseString(s string) string {
	if checkIfPalindrome(s) {
		return s
	}

	runes := []rune(s)
	fmt.Println(runes)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func fibonacci(n int) int {
	fmt.Println(n)
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fmt.Println("Hello, Mac!")

	// fmt.Println(checkIfPalindrome("racecar"))
	// fmt.Println(checkIfPalindrome("hello"))
	// fmt.Println(checkIfPalindrome("madam"))
	// fmt.Println(checkIfPalindrome("level"))
	// fmt.Println(checkIfPalindrome("rotor"))
	// fmt.Println(checkIfPalindrome("deified"))
	// fmt.Println(checkIfPalindrome("civic"))
	// fmt.Println(checkIfPalindrome("radar"))

	// fmt.Println(reverseString("hello"))
	// fmt.Println(reverseString("racecar"))

	fmt.Println(fibonacci(5))
}
