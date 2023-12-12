package Sort

import (
	"math/rand"
	"time"
)

func sort(nums []int) {

}

func quicksort(nums []int, l, r int) {
	if l == r {
		return
	}
	//随机生成一个数作为num的位置
	index := RandomInt(r-l) + l
	mid := partition1(nums, l, r, index-1)
	quicksort(nums, l, mid-1)
	quicksort(nums, mid+1, r)
}

func RandomInt(maxSize int) int {
	// 初始化随机数种子
	rand.Seed(time.Now().Unix())
	// 生成一个随机数
	return rand.Intn(maxSize)
}

// 已知arr[l....r]范围上一定有x这个值
// 划分数组 <=x放左边，>x放右边，并且确保划分完成后<=x区域的最后一个数字是x
func partition1(nums []int, l, r, x int) int {
	a := l
	xi := 0
	for i := l; i <= r; i++ {
		if nums[i] <= nums[x] {
			nums[i], nums[a] = nums[a], nums[i]
			if nums[a] == nums[x] {
				xi = a
			}
		}
	}
	nums[xi], nums[a-1] = nums[a-1], nums[xi]
	return a - 1
}

func quicksort2(nums []int, l, r int) {
	if l >= r {
		return
	}
	index := RandomInt(r-l) + l
	partition2(nums, l, r, index)
	quicksort2(nums, l, index-1)
	quicksort(nums, index+1, r)
}

func partition2(nums []int, l, r, x int) {
	first, last := l, r
	i := l
	for i <= last {
		if nums[i] == nums[x] {
			i++
		} else if nums[i] > nums[x] {
			nums[i], nums[last] = nums[last], nums[i]
			last--
		} else {
			nums[i], nums[first] = nums[first], nums[i]
			i++
			first++
		}
	}
}
