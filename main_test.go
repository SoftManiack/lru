package main

import (
	"fmt"
	"testing"
)

func TestListInsertEnd(t *testing.T) {

	doubleLinkedList := NewDoubleLinkedList(10)

	doubleLinkedList.insertEnd("1")
	doubleLinkedList.insertEnd("2")
	doubleLinkedList.insertEnd("3")

	if doubleLinkedList.head.value != "1" {
		t.Errorf("head error")
	}

	if doubleLinkedList.tail.value != "3" {
		t.Errorf("tail error")
	}

	if doubleLinkedList.tail.next.value != "2" {
		t.Errorf("tail error")
	}
}

func TestListInsertBegin(t *testing.T) {
	doubleLinkedList := NewDoubleLinkedList(10)

	doubleLinkedList.insert("1")
	doubleLinkedList.insert("2")
	doubleLinkedList.insert("4")

	if doubleLinkedList.head.value != "4" {
		t.Errorf("head error")
	}

	if doubleLinkedList.tail.value != "1" {
		t.Errorf("tail error")
	}

	if doubleLinkedList.tail.next.value != "2" {
		t.Errorf("tail error")
	}
}

/*func TestListInsertBeginAndEnd(t *testing.T) {

	doubleLinkedList := NewDoubleLinkedList(10)

	doubleLinkedList.insertBegin("1")
	doubleLinkedList.insertEnd("2")
	doubleLinkedList.insertBegin("4")
	doubleLinkedList.insertEnd("5")

	if doubleLinkedList.head.value != "4" {
		t.Errorf("head error")
	}

	if doubleLinkedList.tail.value != "5" {
		t.Errorf("tail error")
	}

} */

func TestListOverflow(t *testing.T) {

	doubleLinkedList := NewDoubleLinkedList(4)

	doubleLinkedList.insert("2")
	doubleLinkedList.insert("3")
	doubleLinkedList.insert("4")
	doubleLinkedList.insert("5")
	doubleLinkedList.insert("6")
	doubleLinkedList.insert("7")

	fmt.Println(doubleLinkedList.count)
	fmt.Println(doubleLinkedList.tail.value)
	fmt.Println(doubleLinkedList.head.value)
	fmt.Println(doubleLinkedList.size)

}
