package ll

type Node[T any] struct {
	Content T
	Next    *Node[T]
}

func NewNode[T any](content T) *Node[T] {
	return &Node[T]{Content: content}
}

type LinkedList[T any] struct {
	Head *Node[T]
	Tail *Node[T]
}

func NewLinkedList[T any](node *Node[T]) *LinkedList[T] {
	return &LinkedList[T]{Head: node, Tail: node}
}

func (ll *LinkedList[T]) AddTail(node *Node[T]) {
	if ll.Tail == nil {
		ll.Head = node
		ll.Tail = node
		return
	}
	ll.Tail.Next = node
	ll.Tail = node
	ll.Tail.Next = nil
}

func (ll *LinkedList[T]) GetTail() *Node[T] {
	node := ll.Head
	for ; node.Next != nil; node = node.Next {}
	return node
}

func (ll *LinkedList[T]) RemoveTail() *Node[T] {
	// only one node left
	tmp := ll.Tail
	if ll.Head == ll.Tail {
		ll.Head = nil
		ll.Tail = nil
		return tmp
	}
	// only two nodes left
	if ll.Head.Next == ll.Tail {
		ll.Head.Next = nil
		ll.Tail = ll.Head
		return tmp
	}
	for node := ll.Head; node != nil; node = node.Next {
		if node.Next.Next == nil {
			node.Next = nil
			ll.Tail = node
			break
		}
	}
	return tmp
}

func (ll *LinkedList[T]) AddFront(node *Node[T]) {
	if ll.Head == nil {
		ll.Tail = node
	}
	node.Next = ll.Head
	ll.Head = node
}

func (ll *LinkedList[T]) RemoveHead() {
	if ll.Head == nil {
		return
	}
	ll.Head = ll.Head.Next
}

func (ll *LinkedList[T]) AddNodeIdx(idx int, node *Node[T]) *Node[T] {
	if idx == 0 {
		ll.AddFront(node)
		return node
	} else if idx ==  ll.Length() {
		ll.AddTail(node)
		return node
	} else if idx > ll.Length() || idx < 0 {
		return nil
	}

	prev := ll.Head
	curr := ll.Head.Next
	for i := 1; i < idx; i++ {
		prev = prev.Next
		curr = curr.Next
	}
	prev.Next = node
	node.Next = curr
	return node
}

func (ll *LinkedList[T]) GetNodeIdx(idx int) *Node[T] {
	i := 0
	for node := ll.Head; node != nil; node = node.Next {
		if i == idx {
			return node
		}
		i++
	}
	return nil
}


func (ll* LinkedList[T]) DeleteNodeIdx(idx int) {
	if idx == 0 {
		ll.RemoveHead()
	}
	prev := ll.Head
	i := 1
	for curr := prev.Next; curr != nil; curr = curr.Next {
		if curr == ll.Tail {
			ll.RemoveTail()
			break
		} else if i == idx {
			prev.Next = curr.Next
			break
		}
		i++
		prev = curr
	}
	
}

func (ll *LinkedList[T]) DeepCopy() *LinkedList[T] {
	newLL := NewLinkedList(NewNode(ll.Head.Content))
	for node := ll.Head.Next; node != nil; node = node.Next {
		newLL.AddTail(NewNode(node.Content))
	}
	return newLL
}

func (ll *LinkedList[T]) Length() int {
	if ll == nil {
		return 0
	}
	len := 0
	for node := ll.Head; node != nil; node = node.Next {
		len++
	}
	return len
}

func CreateLinkedList[T any](data []T) *LinkedList[T] {
	if data == nil {
		return nil
	}
	ll := NewLinkedList(NewNode(data[0]))
	for _, point := range data[1:] {
		ll.AddTail(NewNode(point))
	}
	return ll
}
