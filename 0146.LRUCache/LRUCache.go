package main

import "fmt"

type LRUCache struct {
	Capactity int
	Cache map[int]*linkedNode
	Head *linkedNode
	Tail *linkedNode
}

type linkedNode struct {
	Key int
    Val int
    Prev *linkedNode
    Next *linkedNode
}

func Constructor(capacity int) LRUCache {
	lruCache := LRUCache{
		Capactity: capacity,
		Cache: make(map[int]*linkedNode),
		Head: &linkedNode{},
		Tail: &linkedNode{},
	}
	lruCache.Head.Next = lruCache.Tail
	lruCache.Tail.Prev = lruCache.Head
    return lruCache
}

func (this *LRUCache) Get(key int) int {
    node, ok := this.Cache[key]
	if ok {
		this.removeNode(node)
		this.addNode(node)
		return node.Val
	}
    return -1
}

func (this *LRUCache) addNode(node *linkedNode) {
	this.Head.Next.Prev = node
	node.Prev = this.Head
	node.Next = this.Head.Next
	this.Head.Next = node
	this.Cache[node.Key] = node
}

func (this *LRUCache) removeNode(node *linkedNode) {
	delete(this.Cache, node.Key)
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	// released memeory
	node.Next = nil
    node.Prev = nil
}

func (this *LRUCache) Put(key int, value int)  {
	// 3(N) -> 2(C) -> 1 -> 3
	node, ok := this.Cache[key]
	if ok {
		node.Val = value
		this.removeNode(node)
		this.addNode(node)
		return
	} else {
		NewCache := linkedNode{
			Key: key,
			Val: value,
		}
		this.addNode(&NewCache)
	}
	if len(this.Cache) > this.Capactity {
		this.removeNode(this.Tail.Prev)
	}
}

var count int
func (this *LRUCache) Show(node *linkedNode) {
	if node != nil {
		fmt.Printf(">> %v %v %v\n", node.Prev, node, node.Next)
		this.Show(node.Next)
	}
	count ++
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	lRUCache := Constructor(2)
	lRUCache.Put(1, 1)
	lRUCache.Put(2, 2)
	println(lRUCache.Get(1))
	lRUCache.Put(3, 3)
	lRUCache.Put(4, 4)
	println(lRUCache.Get(2))
	lRUCache.Show(lRUCache.Head)
}
