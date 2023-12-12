package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const MAXN = 40
const MAXM = 1 << 20

var arr [MAXN]int64
var lsum [MAXM]int64
var rsum [MAXM]int64
var n int
var w int64

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	for scanner.Scan() {
		n, _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		w, _ = strconv.ParseInt(scanner.Text(), 10, 64)

		scanner.Scan()
		words := scanner.Text()
		for i, word := range words {
			arr[i], _ = strconv.ParseInt(string(word), 10, 64)
		}

		fmt.Fprintln(writer, compute())
		writer.Flush()
	}
}

func compute() int64 {
	lsize := f(0, n>>1, 0, w, lsum[:], 0)
	rsize := f(n>>1, n, 0, w, rsum[:], 0)

	sort.Slice(lsum[:lsize], func(i, j int) bool {
		return lsum[i] < lsum[j]
	})
	sort.Slice(rsum[:rsize], func(i, j int) bool {
		return rsum[i] < rsum[j]
	})

	var ans int64 = 0
	for i, j := lsize-1, 0; i >= 0; i-- {
		for j < rsize && lsum[i]+rsum[j] <= w {
			j++
		}
		ans += int64(j)
	}
	return ans
}

func f(i, e int, s, w int64, ans []int64, j int) int {
	if s > w {
		return j
	}

	if i == e {
		ans[j] = s
		j++
	} else {
		j = f(i+1, e, s, w, ans, j)
		j = f(i+1, e, s+arr[i], w, ans, j)
	}
	return j
}
