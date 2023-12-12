package dynamic_programming

func numDecodings(s string) int {
	return f1([]byte(s), 0)
}

func f1(str []byte, i int) int {
	if i == len(str) {
		return 1
	}

	if str[i] == '0' {
		return 0
	}

	var ans int
	//自己单独转
	if str[i] == '*' {
		ans = f1(str, i+1) * 9
	} else {
		ans = f1(str, i+1)
	}

	if i+1 < len(str) {
		if str[i] != '*' {
			if str[i+1] != '*' {
				if (str[i]-'0')*10+str[i+1]-'0' <= 26 {
					ans += f1(str, i+2)
				}
			} else {
				// num  *
				//  i  i+1
				if str[i] == '1' {
					ans += f1(str, i+2) * 9
				} else {
					ans += f1(str, i+2) * 6
				}
			}
		} else {
			if str[i+1] != '*' {
				if str[i+1] < '6' {
					ans += f1(str, i+2) * 2
				} else {
					ans += f1(str, i+2)
				}
			} else {
				ans += f1(str, i+2) * 15
			}
		}
	}
	return ans
}

func numDecodings2(s string) int {
	var mod = 1000000007

	str := []byte(s)
	dp := make([]int, len(str)+1)
	n := len(s)
	dp[n] = 1

	for i := n - 1; i >= 0; i-- {
		if str[i] != '0' {
			if str[i] == '*' {
				dp[i] = 9 * dp[i+1]
			} else {
				dp[i] = dp[i+1]
			}
			if i+1 < n {
				if str[i] != '*' {
					if str[i+1] != '*' {
						if (str[i]-'0')*10+str[i+1]-'0' <= 26 {
							dp[i] += dp[i+2]
						}
					} else {
						if str[i] == '1' {
							dp[i] += 9 * dp[i+2]
						}
						if str[i] == '2' {
							dp[i] += 6 * dp[i+2]
						}
					}
				} else {
					if s[i+1] != '*' {
						if s[i+1] <= '6' {
							dp[i] += dp[i+2] * 2
						} else {
							dp[i] += dp[i+2]
						}
					} else {
						dp[i] += dp[i+2] * 15
					}
				}
			}
			dp[i] %= mod
		}
	}

	return dp[0]
}
