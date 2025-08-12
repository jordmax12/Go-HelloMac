package main

import (
	"errors"
	"fmt"
	"time"
)

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
	// fmt.Println(n)
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func retryWithFibonacciBackoff(maxAttempts int, call func() error) error {
	a := time.Duration(0)
	b := 50 * time.Millisecond  // base delay
	maxDelay := 5 * time.Minute // cap

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		if err := call(); err == nil {
			return nil
		}
		fmt.Println("attempt", attempt, "with delay", b)
		time.Sleep(b)
		// next delay = sum of previous two
		a, b = b, a+b
		if b > maxDelay {
			b = maxDelay
		}
	}
	return errors.New("max attempts reached")
}

func testRetry() {
	i := 0
	maxAttempts := 100
	err := retryWithFibonacciBackoff(maxAttempts, func() error {
		i++
		if i < 20 {
			fmt.Println("transient failure")
			return errors.New("try again")
		}
		fmt.Println("success on attempt", i)
		return nil
	})
	fmt.Println("result:", err)
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

	// fmt.Println(fibonacci(20))

	testRetry()
}
