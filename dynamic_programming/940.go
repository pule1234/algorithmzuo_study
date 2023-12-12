package dynamic_programming

import "fmt"

func distinctSubseqII(s string) int {
	mod := 1000000007
	str := []byte(s)
	cnt := make([]int, 26)

	all, newadd := 1, 0
	for _, v := range str {
		newadd = (all - cnt[v-'a'] + mod) % mod
		fmt.Println("newadd = ", newadd)
		cnt[v-'a'] = (cnt[v-'a'] + newadd) % mod

		all = (all + newadd) % mod
		fmt.Println("all  =", all)

	}

	return (all - 1 + mod) % mod
}
