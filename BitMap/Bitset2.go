package BitMap

import (
	"fmt"
	"strings"
)

// 测试链接 : https://leetcode-cn.com/problems/design-bitset/

type Bitset struct {
	set     []int
	reverse bool
	ones    int
	zeros   int
	Size    int
}

func Constructor(size int) Bitset {
	arr := make([]int, (size+31)/32)
	return Bitset{
		set:     arr,
		reverse: false,
		ones:    0,
		zeros:   size,
		Size:    size,
	}
}

// 将下标为 idx 的位上的值更新为 1 。如果值已经是 1 ，则不会发生任何改变。
func (this *Bitset) Fix(idx int) {
	index := idx / 32
	bit := idx % 32
	if !this.reverse {
		if this.set[index]&(1<<bit) == 0 {
			this.set[index] |= (1 << bit)
			this.zeros--
			this.ones++
		}
	} else {
		if this.set[index]&(1<<bit) != 0 {
			this.zeros--
			this.ones++
			this.set[index] ^= (1 << bit)
		}
	}
	fmt.Printf("Fix   ones = %d, zeros = %d\n", this.ones, this.zeros)
}

func (this *Bitset) Unfix(idx int) {
	index := idx / 32
	bit := idx % 32
	if !this.reverse {
		if this.set[index]&(1<<bit) != 0 {
			this.set[index] ^= (1 << bit)
			this.ones--
			this.zeros++
		}
	} else {
		if this.set[index]&(1<<bit) == 0 {
			this.set[index] |= (1 << bit)
			this.ones--
			this.zeros++
		}
	}
	fmt.Printf("Unfix   ones = %d, zeros = %d\n", this.ones, this.zeros)
}

func (this *Bitset) Flip() {
	this.reverse = !this.reverse
	this.ones, this.zeros = this.zeros, this.ones
	fmt.Printf("Filp   ones = %d, zeros = %d\n", this.ones, this.zeros)
}

func (this *Bitset) All() bool {
	fmt.Printf("All   ones = %d, zeros = %d\n", this.ones, this.zeros)
	return this.ones == this.Size

}

func (this *Bitset) One() bool {
	return this.ones > 0
}

func (this *Bitset) Count() int {
	fmt.Printf("Count   ones = %d, zeros = %d\n", this.ones, this.zeros)
	return this.ones
}

func (this *Bitset) ToString() string {
	var sb strings.Builder
	for i := 0; i < this.Size; i++ {
		if this.reverse {
			bit := this.Get(i)
			if bit {
				sb.WriteByte('0')
			} else {
				sb.WriteByte('1')
			}
		} else {
			bit := this.Get(i)
			if bit {
				sb.WriteByte('1')
			} else {
				sb.WriteByte('0')
			}
		}

	}
	return sb.String()
}

func (this *Bitset) Get(idx int) bool {
	index := idx / 32
	bit := idx % 32
	mask := 1 << bit
	return this.set[index]&mask != 0
}
