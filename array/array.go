/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package array

type Array struct {
	data []int
	size int
}

func Constructor(capacity int) *Array {
	return &Array{
		data: make([]int, capacity),
	}
}
func (this *Array) GetSize() int {
	return this.size
}
func (this *Array) IsEmpty() bool {
	return this.size == 0
}

func (this *Array) GetCapacity() int {
	return len(this.data)
}
func (this *Array) Add(index, e int) {
	if index < 0 || index > this.size {
		panic("index is illegal")
	}
	//TODO　resize
	for i := this.size - 1; i >= index; i-- {
		this.data[i+1] = this.data[i]
	}
	this.data[index] = e
	this.size++
}
func (this *Array) AddLast(e int) {
	this.Add(this.size, e)
}
func (this *Array) AddFirst(e int) {
	this.Add(0, e)
}
func (this *Array) Get(index int) int {
	if index < 0 || index > this.size {
		panic("index is illegal")
	}
	return this.data[index]
}
func (this *Array) GetFirst() int {
	return this.Get(0)
}
func (this *Array) GetLast() int {
	return this.Get(this.size - 1)
}

//修改index索引位置的元素为e
func (this *Array) Set(index, e int) {
	if index < 0 || index > this.size {
		panic("index is illegal")
	}
	this.data[index] = e
}

//查找数组中是否有元素e O(n)
func (this *Array) Contains(e int) bool {
	for i := 0; i < this.size; i++ {
		if this.data[i] == e {
			return true
		}
	}
	return false
}

//查找数组中元素e所在的索引，如果不存在元素e，则返回-1 O(n)
func (this *Array) Find(e int) int {
	for i := 0; i < this.size; i++ {
		if this.data[i] == e {
			return i
		}
	}
	return -1
}

//从数组中删除index位置的元素, 返回删除的元素
func (this *Array) Remove(index int) int {
	if index < 0 || index > this.size {
		panic("index is illegal")
	}
	for i := index + 1; i < this.size; i++ {
		this.data[i-1] = this.data[i]
	}
	//TODO　resize
	ret := this.data[index]
	this.size--
	this.data = append(this.data[:index], this.data[index+1:]...)
	return ret
}

// 从数组中删除第一个元素, 返回删除的元素
func (this *Array) RemoveFirst() {
	this.Remove(0)
}

// 从数组中删除最后一个元素, 返回删除的元素
func (this *Array) RemovLast() {
	this.Remove(this.size - 1)
}

// 从数组中删除元素e
func (this *Array) RemoveElement(e int) {
	index := this.Find(e)
	if index != -1 {
		this.Remove(index)
	}
}

//// 将数组空间的容量变成newCapacity大小
//private function resize($newCapacity)
//{
//$newData = (new self($newCapacity))->data;
//for ($i = 0; $i < $this->size; $i++) {
//$newData[$i] = $this->data[$i];
//}
//$this->data = $newData;
//$this->capacity = $newCapacity;
//}