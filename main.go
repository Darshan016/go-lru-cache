package main

import (
	"fmt"
)

const SIZE = 5

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

type Cache struct {
	Queue Queue
	Hash  Hash
}

type Hash map[string]*Node

func NewCache() Cache {
	return Cache{Queue: NewQueue(), Hash: Hash{}}
}

func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}
	head.Right = tail
	tail.Left = head

	return Queue{Head: head, Tail: tail}
}

func (cache *Cache) Check(str string) {
	node := &Node{}

	if val, ok := cache.Hash[str]; ok {
		node = cache.Remove(val)
	} else {
		node = &Node{Val: str}
	}
	cache.Add(node)
	cache.Hash[str] = node
}

func (cache *Cache) Remove(node *Node) *Node {
	fmt.Printf("remove: %s\n", node.Val)
	left := node.Left
	right := node.Right

	left.Right = right
	right.Left = left

	cache.Queue.Length -= 1
	delete(cache.Hash, node.Val)
	return node
}

func (cache *Cache) Add(node *Node) {
	fmt.Printf("add: %s\n", node.Val)
	temp := cache.Queue.Head.Right

	cache.Queue.Head.Right = node
	node.Left = cache.Queue.Head
	node.Right = temp
	temp.Left = node

	cache.Queue.Length++
	if cache.Queue.Length > SIZE {
		cache.Remove(cache.Queue.Tail.Left)
	}
}

func (cache *Cache) Display() {
	cache.Queue.Display()
}

func (queue *Queue) Display() {
	node := queue.Head.Right
	fmt.Printf("%d - [", queue.Length)
	for i := 0; i < queue.Length; i++ {
		fmt.Printf("{%s}", node.Val)
		if i < queue.Length-1 {
			fmt.Printf("<-->")
		}
		node = node.Right
	}
	fmt.Printf("]\n")
}

func main() {
	fmt.Println("cache start")

	cache := NewCache()
	for _, word := range []string{"oregairu", "aot", "naruto", "cote", "code geass", "ditf", "aot", "death note"} {
		cache.Check(word)
		cache.Display()
	}
}
