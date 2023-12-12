package Sliding_window

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	// 车辆尝试从0~n-1出发，看能不能走一圈，l
	// r : 窗口即将进来数字的位置
	// len : 窗口大小
	// sum : 窗口累加和

	for l, r, length, sum := 0, 0, 0, 0; l < n; l++ {
		for sum >= 0 {
			if length == n {
				return l
			}
			// r = 窗口即将进来的位置
			r = (l + length) % n
			sum += gas[r] - cost[r]
			length++
		}

		length--
		sum -= gas[l] - cost[l]
	}
	return -1
}
