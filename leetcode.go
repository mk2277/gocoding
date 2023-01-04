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
