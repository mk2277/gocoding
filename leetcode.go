package main

//Two sum
func twoSum(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}

}

//Palindrome Number

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	y := 0
	z := x

	for x != 0 {

		y = y*10 + x%10
		x = x / 10

	}

	return z == y
}

//roman to integer
func romanToInt(s string) int {

	var romanMap = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	var result = romanMap[s[len(s)-1]]

	for i := len(s) - 2; i >= 0; i-- {
		if romanMap[s[i]] < romanMap[s[i+1]] {
			result -= romanMap[s[i]]

		} else {
			result += romanMap[s[i]]

		}
	}
	return result
}
