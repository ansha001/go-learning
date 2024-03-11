package main

import (
	"fmt"
)

type Entry struct {
	key   string
	value int
	next  *Entry
}
type HashTable struct {
	size    int
	buckets []*Entry
}

func newHashTable(size int) *HashTable {
	return &HashTable{
		size:    size,
		buckets: make([]*Entry, size),
	}
}

func (ht *HashTable) hash(key string) int {
	hash := 0
	for _, char := range key {
		hash += int(char)
	}
	return hash % ht.size
}

func (ht *HashTable) insert(key string, value int) {
	index := ht.hash(key)
	entry := &Entry{key: key, value: value}

	// Quadratic probing
	for i := 0; i < ht.size; i++ {
		probeIndex := (index + i*i) % ht.size
		if ht.buckets[probeIndex] == nil {
			ht.buckets[probeIndex] = entry
			return
		}
	}
	fmt.Println("Hash table is full. Unable to insert:", key)
}

func (ht *HashTable) search(key string) (int, bool) {
	index := ht.hash(key)
	current := ht.buckets[index]

	for current != nil {
		if current.key == key {
			return current.value, true
		}
		current = current.next
	}
	return 0, false
}

func main() {
	hashTable := newHashTable(5)

	var num, option int
	for option != 3 {
		fmt.Println("\n1. Insert\n2. Search\n3. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&option)

		switch option {
		case 1:
			fmt.Print("Enter key: ")
			var key string
			fmt.Scan(&key)
			fmt.Print("Enter value: ")
			fmt.Scan(&num)
			hashTable.insert(key, num)
		case 2:
			fmt.Print("Enter key to search: ")
			var key string
			fmt.Scan(&key)
			if value, found := hashTable.search(key); found {
				fmt.Printf("Value for '%s': %d\n", key, value)
			} else {
				fmt.Printf("Key '%s' not found.\n", key)
			}
		}
	}
}
