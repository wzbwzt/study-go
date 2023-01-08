//go:build ignore
// +build ignore

package main

import "container/list"

/*
设计哈希映射
*/

const base = 769

type entry struct {
	key, value int
}

type MyHashMap struct {
	data []list.List
}

func Constructor() MyHashMap {
	return MyHashMap{make([]list.List, base)}
}

func (m *MyHashMap) hash(key int) int {
	return key % base
}

func (m *MyHashMap) Put(key, value int) {
	h := m.hash(key)
	for e := m.data[h].Front(); e != nil; e = e.Next() {
		if et := e.Value.(entry); et.key == key {
			e.Value = entry{key, value}
			return
		}
	}
	m.data[h].PushBack(entry{key, value})
}

func (m *MyHashMap) Get(key int) int {
	h := m.hash(key)
	for e := m.data[h].Front(); e != nil; e = e.Next() {
		if et := e.Value.(entry); et.key == key {
			return et.value
		}
	}
	return -1
}

func (m *MyHashMap) Remove(key int) {
	h := m.hash(key)
	for e := m.data[h].Front(); e != nil; e = e.Next() {
		if e.Value.(entry).key == key {
			m.data[h].Remove(e)
		}
	}
}
