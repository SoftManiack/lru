package main

type Node struct {
	previous *Node
	value    string
	next     *Node
}

type DoubleLinkedList struct {
	head  *Node
	tail  *Node
	size  int
	count int
}

func NewDoubleLinkedList(size int) *DoubleLinkedList {

	doubleLinkedList := &DoubleLinkedList{
		head: nil,
		tail: nil,
		size: size,
	}

	return doubleLinkedList
}

// Вставка в начало списка
func (l *DoubleLinkedList) insert(value string) Node {

	var Node Node

	if l.count == l.size {

		delete(hashMap, l.tail.value)

		l.tail = l.tail.next

		l.tail.previous = nil

		Node.value = value
		Node.previous = l.head
		Node.next = nil

		l.head.next = &Node

		l.head = &Node

	} else {
		if l.head == nil {
			Node.value = value
			Node.previous = nil
			Node.next = nil

			l.head = &Node
			l.tail = &Node
			l.count += 1
		} else {

			Node.value = value
			Node.previous = l.head

			l.head.next = &Node

			l.head = &Node
			l.count += 1
		}

	}

	return Node
}

func (l *DoubleLinkedList) RemoveBegin() string {

	if l.head == nil {
		return ""
	}
	var result string = l.head.value

	if l.head == l.tail {
		l.head = nil
		l.tail = nil
	} else {
		l.head = l.head.next
		l.head.previous = nil
	}

	l.count -= 1

	return result

}

// Вставка в конец
func (l *DoubleLinkedList) insertEnd(value string) {

	if l.tail == nil {
		node := Node{
			value:    value,
			previous: nil,
			next:     nil,
		}

		l.tail = &node
		l.head = &node

	} else {
		node := Node{
			value:    value,
			previous: nil,
			next:     l.tail,
		}

		l.tail.previous = &node
		l.tail = &node
	}

	l.size += 1
}
