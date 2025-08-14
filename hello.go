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

func sortedSquares(nums []int) []int {
	var left = 0
	var right = len(nums) - 1
	var pos = len(nums) - 1
	var newArr = make([]int, len(nums))

	for left <= right {
		var leftSq = nums[left] * nums[left]
		var rightSq = nums[right] * nums[right]

		fmt.Printf("leftSq: %d \n rightSq: %d \n pos: %d \n left: %d\n right: %d", leftSq, rightSq, pos, left, right)
		fmt.Println(newArr)

		if leftSq > rightSq {
			newArr[pos] = leftSq
			left++
		} else {
			newArr[pos] = rightSq
			right--
		}

		pos--
	}

	return newArr
}

func findMaxAverage(nums []int, k int) float64 {
	currSum := 0
	for i := 0; i < k; i++ {
		currSum += nums[i]
	}

	maxSum := currSum

	for i := k; i < len(nums); i++ {
		currSum = currSum + nums[i] - nums[i-k]
		maxSum = max(maxSum, currSum)
	}

	return float64(maxSum) / float64(k)
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
