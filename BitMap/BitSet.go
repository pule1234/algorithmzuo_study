package BitMap

type bitmap struct {
	set []int
}

func (this *bitmap) Con(n int) *bitmap {
	n = (n + 31) / 32
	arr := make([]int, n)
	return &bitmap{
		set: arr,
	}
}

// 先将该数 /32 确定该数一个存储在哪一个下标， 在对该数%32 判断该数应该存储在哪一位
func (this *bitmap) add(n int) {
	this.set[n/32] |= 1 << (n % 32)
}

// 将该位置置0
func (this *bitmap) remove(n int) {
	this.set[n/32] &= ^(1 << (n % 32))
}

// 若存在该数，则将该数置0， 若不存在该数，则将该数置1
func (this *bitmap) reverse(n int) {
	this.set[n/32] ^= 1 << (n % 32)
}

// 判断该数是否存在
func (this *bitmap) contains(n int) bool {
	return (this.set[n/32]>>(n%32))&1 == 1
}
