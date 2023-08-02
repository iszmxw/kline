package wsdata

type Node struct {
	data interface{} // 存储的数据
	prev *Node       // 前一个节点的指针
	next *Node       // 后一个节点的指针
}

type LinkedList struct {
	head   *Node // 链表的头节点指针
	tail   *Node // 链表的尾节点指针
	Length int   // 链表的长度
}

// Enqueue 将数据入队，即从尾部插入数据
func (list *LinkedList) Enqueue(data interface{}) {
	newNode := &Node{data: data}

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.prev = list.tail
		list.tail.next = newNode
		list.tail = newNode
	}

	list.Length++
}

// Dequeue 将数据出队，即从头部删除数据，并返回删除的数据
func (list *LinkedList) Dequeue() interface{} {
	if list.head == nil {
		return nil
	}

	data := list.head.data
	list.head = list.head.next

	if list.head != nil {
		list.head.prev = nil
	} else {
		list.tail = nil
	}

	list.Length--

	return data
}
