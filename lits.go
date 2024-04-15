package lru

type Node struct {
	previous *Node
	value    string
	next     *Node
}

type DoubleLinkedList struct {
	heade *Node
	tail  *Node
	size  int
}

func (l *DoubleLinkedList) NewDoubleLinkedList(size int) *DoubleLinkedList {

	return &DoubleLinkedList{
		heade: nil,
		tail:  nil,
		size:  size,
	}
}

func (l *DoubleLinkedList) Add(value string) Node {

	var Node Node

	if l.heade == nil {
		Node.value = value
		Node.previous = nil
		Node.next = nil

		l.heade = &Node
		l.tail = &Node
		l.size += 1
	}

	return Node
}

func (l *DoubleLinkedList) Remove() {

}

func (l *DoubleLinkedList) insertEnd(value string) {
	if l.tail == nil {
		node := Node{
			value:    value,
			previous: nil,
			next:     nil,
		}

		l.tail = &node
		l.heade = &node

	} else {
		node := Node{
			value:    value,
			previous: l.tail,
			next:     nil,
		}

		l.tail.next = &node
		l.tail = &node
	}

	l.size += 1
}
