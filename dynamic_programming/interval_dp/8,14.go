package interval_dp

func countEval(s string, result int) int {
	str := []byte(s)
	n := len(str)
	dp := make([][][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][]int, n)
	}
	ft := f1(str, 0, n-1, dp)

	return ft[result]
}

// s[l...r]是表达式的一部分，且一定符合范式
// 0/1  逻  0/1   逻       0/1
//
//	l  l+1  l+2  l+3........r
//
// s[l...r]  0 : ?
//
//	1 : ?
func f1(str []byte, l, r int, dp [][][]int) []int {
	if dp[l][r] != nil {
		return dp[l][r]
	}

	f := 0 // 为0
	t := 0 // 为1

	if l == r { // 只剩一个字符，0/1
		if str[l] == '0' {
			f = 1
		} else {
			f = 0
		}

		if str[l] == '1' {
			t = 1
		} else {
			t = 0
		}
	} else {
		tmp := []int{}
		for k := l + 1; k < r; k += 2 {
			// l ... r
			// 枚举每一个逻辑符号最后执行 k = l+1 ... r-1  k+=2
			tmp = f1(str, l, k-1, dp)
			a := tmp[0]
			b := tmp[1]
			tmp = f1(str, k+1, r, dp)
			c := tmp[0]
			d := tmp[1]
			if str[k] == '&' {
				f += a*c + a*d + b*c
				t += b * d
			} else if str[k] == '|' {
				f += a * c
				t += a*d + b*c + b*d
			} else {
				f += a*c + b*d
				t += a*d + b*c
			}
		}
	}
	ft := []int{f, t}
	dp[l][r] = ft
	return ft
}
