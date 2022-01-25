package main


type TreeNode struct {
	Val int
	Next *TreeNode
	Prev *TreeNode
}

func (t *TreeNode) Push(Val int) {
	NewNode := &TreeNode{Val: Val}
	NewNode.Next = t
	if t != nil {
		t.Prev = NewNode
	}
	t = NewNode
}

func (t *TreeNode) Insert(PrevNode *TreeNode, Val int) {
	if PrevNode == nil {
		return
	}
	NewNode := &TreeNode{Val: Val}
	NewNode.Next = PrevNode.Next
	PrevNode.Next = NewNode
	NewNode.Prev = PrevNode
	if NewNode.Next != nil {
		NewNode.Next.Prev = NewNode
	}
}

func (t *TreeNode) Append(Val int) {
	NewNode := &TreeNode{Val: Val}
	NewNode.Next = nil
	if t == nil {
		NewNode.Prev = nil
		t = NewNode
		return
	}
	last := t
	for last.Next != nil {
		last = last.Next
	}
	last.Next = NewNode
	NewNode.Prev = last
}

func (t *TreeNode) Show() {
	if t != nil {
		println(t.Val)
		t.Next.Show()
	}
}

func main() {
	root := &TreeNode{Val: 15}
	root.Push(3)
	root.Push(5)
	root.Insert(root.Prev, 100)
	root.Append(77)
	root.Append(80)
	root.Show()
}