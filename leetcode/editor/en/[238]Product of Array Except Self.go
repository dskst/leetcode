//Given an integer array nums, return an array answer such that answer[i] is
//equal to the product of all the elements of nums except nums[i].
//
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit
//integer.
//
// You must write an algorithm that runs in O(n) time and without using the
//division operation.
//
//
// Example 1:
// Input: nums = [1,2,3,4]
//Output: [24,12,8,6]
//
// Example 2:
// Input: nums = [-1,1,0,-3,3]
//Output: [0,0,9,0,0]
//
//
// Constraints:
//
//
// 2 <= nums.length <= 10⁵
// -30 <= nums[i] <= 30
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit
//integer.
//
//
//
// Follow up: Can you solve the problem in O(1) extra space complexity? (The
//output array does not count as extra space for space complexity analysis.)
//
// Related Topics Array Prefix Sum 👍 21418 👎 1302

package main

// leetcode submit region begin(Prohibit modification and deletion)
func productExceptSelf(nums []int) []int {
	answers := make([]int, 0, len(nums))

	for i := 0; i < len(nums); i++ {
		answer := 1
		for idx, num := range nums {
			if idx == i {
				continue
			}
			answer *= num
		}
		answers = append(answers, answer)
	}

	return answers
}

//leetcode submit region end(Prohibit modification and deletion)