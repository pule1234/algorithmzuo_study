package data_structure

import "math/rand"

type RandomizedSet struct {
	hash map[int]int
	arr  []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		hash: make(map[int]int),
		arr:  make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.hash[val]; ok {
		return false
	} else {
		this.hash[val] = len(this.arr)
		this.arr = append(this.arr, val)
	}
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.hash[val]; !ok {
		return false
	}

	valindex := this.hash[val]
	endvalue := this.arr[len(this.arr)-1]
	this.hash[endvalue] = valindex
	this.arr[valindex] = endvalue
	this.arr = this.arr[:len(this.arr)-1]
	delete(this.hash, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.arr[rand.Intn(len(this.arr))]
}
