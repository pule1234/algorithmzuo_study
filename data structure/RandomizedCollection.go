package data_structure

import (
	"math/rand"
	"runtime"
)

type RandomizedCollection struct {
	hash map[int]map[int]struct{}
	arr  []int
}

func Constructor() RandomizedCollection {
	return RandomizedCollection{map[int]map[int]struct{}{},[]int}
}

func (this *RandomizedCollection) Insert(val int) bool {
	ids,has := this.hash[val]
	if !has {
		ids = map[int]struct{}{}
		this.hash[val] = ids
	}
	ids[len(this.arr)] = struct{}{}
	this.arr = append(this.arr,val)
	return !has
}

func (this *RandomizedCollection) Remove(val int) bool {
	ids , has :=this.hash[val]
	if !has {
		return false
	}
	var i int
	for id := range ids {
		i = id
		break
	}
	n := len(this.arr)
	this.arr[i] =this.arr[n-1]
	delete(ids,i)
	delete(this.hash[this.arr[i]],n-1)
	if i < n-1 {
		this.hash[this.arr[i]][i] = struct{}{}
	}
	if len(ids) == 0 {
		delete(this.hash,val)
	}
	this.arr = this.arr[:n-1]
	return true
}

func (this *RandomizedCollection) GetRandom() int {
	return this.arr[rand.Intn(len(this.arr))]
}
