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

func (l *DoubleLinkedList) Add(value string) *Node {

	Node Node{}

	if l.heade == nil {
		l.heade = Node
		Node.value = value
		Node.previous = nil
		Node.next = nil
	}

	return Node
}
