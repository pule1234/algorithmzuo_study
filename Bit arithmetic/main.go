package main

import "fmt"

// Brian Kernighan算法
// 提取出二进制里最右侧的1
// 判断一个整数是不是2的幂
// 测试链接 : https://leetcode.cn/problems/power-of-two/
// 先提取出最右侧的1 ， 然后与源数据相&  判断是否相等
func isPowerOfTwo(n int) bool {
	return n > 0 && n == n&(-n)
}

// 判断一个整数是不是3的幂
// 测试链接 : https://leetcode.cn/problems/power-of-three/
func isPowerOfThree(n int) bool {
	// 1162261467是int型范围内，最大的3的幂，它是3的19次方
	return n > 0 && 1162261467%n == 0
}

// 已知n是非负数
// 返回大于等于n的最小的2某次方
// 如果int范围内不存在这样的数，返回整数最小值
func near2power(n int) int {
	//首先将其-1
	// 然后将-1后的值 | 该值右移1位
	n = n - 1
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	return n + 1
}

func main() {
	fmt.Println(near2power(100))
}

// 给你两个整数 left 和 right ，表示区间 [left, right]
// 返回此区间内所有数字 & 的结果
// 包含 left 、right 端点
// 测试链接 : https://leetcode.cn/problems/bitwise-and-of-numbers-range/
func rangeBitwiseAnd(left int, right int) int {
	for right > left {
		right -= right & (-right)
	}
	return right
}

// 逆序二进制的状态
// 测试链接 : https://leetcode.cn/problems/reverse-bits/
func reverseBits(num uint32) uint32 {
	//暴力解法 逐位
	res := 0
	for i := 0; i < 32; i++ {
		res |= num & 1 << (32 - i)
		num >>= 1
	}

	// 魔幻 , 分治
	num = ((num & 0xaaaaaaaa) >> 1) | ((num & 0x55555555) << 1)
	num = ((num & 0xcccccccc) >> 2) | ((num & 0x33333333) << 2)
	num = ((num & 0xf0f0f0f0) >> 4) | ((num & 0x0f0f0f0f) << 4)
	num = ((num & 0xff00ff00) >> 8) | ((num & 0x00ff00ff) << 8)
	num = (num >> 16) | (num << 16)
	return uint32(res)
}

// 返回n的二进制中有几个1
// 两个整数之间的 汉明距离 指的是这两个数字对应二进制位不同的位置的数目。
// 给你两个整数 x 和 y，计算并返回它们之间的汉明距离
// 测试链接 : https://leetcode.cn/problems/hamming-distance/
func hammingDistance(x int, y int) int {
	res := 0

	for i := 0; i < 32; i++ {
		sx := x & 1
		sy := y & 1
		if sx != sy {
			res++
		}
		x >>= 1
		y >>= 1
	}
	return res
}