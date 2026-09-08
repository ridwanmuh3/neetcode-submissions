type Node struct {
    Value int
    Next *Node
}

type LinkedList struct {
    Head *Node
    Size int
}

func NewLinkedList() *LinkedList {
    return &LinkedList{}
}

func (ll *LinkedList) Get(index int) int {
    if index < 0 || index >= ll.Size || ll.Head == nil {
        return -1 
    }

    curr := ll.Head
    for i := 0; i < index; i++ {
        curr = curr.Next
    }

    return curr.Value

}

func (ll *LinkedList) InsertHead(val int) {
    newNode := &Node{
        Value: val,
        Next: nil,
    }
    if ll.Head == nil {
        ll.Head = newNode
    } else {
        newNode.Next = ll.Head
        ll.Head = newNode
    }
    ll.Size++
}

func (ll *LinkedList) InsertTail(val int) {
    newNode := &Node{
        Value: val,
        Next: nil,
    }
    if ll.Head == nil {
        ll.Head = newNode
    } else {
        iterationNode := ll.Head
        for iterationNode.Next != nil {
            iterationNode = iterationNode.Next
        }
        iterationNode.Next = newNode
    }
    ll.Size++
}

func (ll *LinkedList) Remove(index int) bool {
    if ll.Head == nil || index < 0 || index >= ll.Size {
        return false
    }

    if index == 0 {
        ll.Head = ll.Head.Next
        ll.Size--
        return true
    }

    current := ll.Head
    for i := 0; i < index-1; i++ {
        current = current.Next
    }

    if current.Next != nil {
        current.Next = current.Next.Next
        ll.Size--
        return true
    }

    return false
}

func (ll *LinkedList) GetValues() []int {
    var values []int
    iteratorNode := ll.Head
    for iteratorNode != nil {
        values = append(values, iteratorNode.Value)
        iteratorNode = iteratorNode.Next
    }
    return values
}
