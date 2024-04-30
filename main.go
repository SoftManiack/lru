package main

import "fmt"

var hashMap map[string]string

var doubleLinkedList DoubleLinkedList

func main() {

	hashMap = make(map[string]string)

	doubleLinkedList = *NewDoubleLinkedList(10)

	insertLRU("a", "1")
	insertLRU("b", "11")
	insertLRU("c", "1")
	insertLRU("d", "100")
	insertLRU("e", "10")
	insertLRU("g", "2")
	insertLRU("c", "31")
	insertLRU("k", "3")
	insertLRU("f", "90")
	insertLRU("r", "7")
	insertLRU("t", "333")
	insertLRU("i", "333")

	fmt.Println(readLRU("a"))
	fmt.Println(readLRU("b"))
	fmt.Println(readLRU("c"))
	fmt.Println(readLRU("t"))
}

func insertLRU(key string, value string) {

	doubleLinkedList.insert(key)

	hashMap[key] = value
}

func readLRU(key string) string {

	return hashMap[key]
}
