package Dual_pointers

import (
	"math"
	"sort"
)

func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)

	ans := 0
	for i, j := 0, 0; i < len(houses); i++ {
		for !beat(houses, heaters, i, j) {
			j++
		}
		ans = max(ans, int(math.Abs(float64(houses[i]-heaters[j]))))
	}
	return ans
}

// 这个函数含义：
// 当前的地点houses[i]由heaters[j]来供暖是最优的吗？
// 当前的地点houses[i]由heaters[j]来供暖，产生的半径是a
// 当前的地点houses[i]由heaters[j + 1]来供暖，产生的半径是b
// 如果a < b, 说明是最优，供暖不应该跳下一个位置
// 如果a >= b, 说明不是最优，应该跳下一个位置

func beat(houses []int, heaters []int, i, j int) bool {
	return j == len(heaters)-1 || math.Abs(float64(heaters[j]-houses[i])) < math.Abs(float64(heaters[j+1]-houses[i]))
}
