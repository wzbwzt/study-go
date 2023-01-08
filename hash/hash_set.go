//go:build ignore
// +build ignore

package main

import "container/list"

/*
设计哈希集合
*/

const base = 769

type MyHashSet struct {
	data []list.List
}

func Constructor() MyHashSet {
	return MyHashSet{data: make([]list.List, base)}
}

func (this *MyHashSet) hash(key int) int {
	return key % base
}

func (this *MyHashSet) Add(key int) {
	h := this.hash(key)
	for e := this.data[h].Front(); e != nil; e = e.Next() {
		if v := e.Value.(int); v == key {
			return
		}
	}
	this.data[h].PushBack(key)
}

func (this *MyHashSet) Remove(key int) {
	h := this.hash(key)
	for e := this.data[h].Front(); e != nil; e = e.Next() {
		if v := e.Value.(int); v == key {
			this.data[h].Remove(e)
		}
	}
}

func (this *MyHashSet) Contains(key int) bool {
	h := this.hash(key)
	for e := this.data[h].Front(); e != nil; e = e.Next() {
		if v := e.Value.(int); v == key {
			return true
		}
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
