package Dual_pointers

// 设定快慢指针，如果有重复的数字，那么这两个指针一定会相遇，并且快指针回到起点重新走，两者会在相同的数字的节点处相遇
func findDuplicate(nums []int) int {
	if nums == nil || len(nums) < 2 {
		return -1
	}

	slow := nums[0]
	fast := nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]] //相当向前两部
	}

	fast = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}
