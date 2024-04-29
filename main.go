package main

import (
	"fmt"
)

var hashMap map[string]string

var doubleLinkedList DoubleLinkedList

func main() {

	doubleLinkedList := NewDoubleLinkedList(10)

	fmt.Println(doubleLinkedList.head)
	fmt.Println(doubleLinkedList.tail)
}

func insertBeginList(key string, value string) {

	hashMap[key] = value
	doubleLinkedList.insert(key)
}
