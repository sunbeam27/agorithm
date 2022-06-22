package main

/*
使用双向链表+hashmap
双向链表 写的复杂度O(1)
map 读的复杂度O(1)
*/
// doublyLinkedNode defines a node for double-linked-list
type doublyLinkedNode struct {
	prev, next *doublyLinkedNode
	key, val   int
}

// LRUCache defines a object for cache
type LRUCache struct {
	len, cap    int
	first, last *doublyLinkedNode         //head,tail
	nodes       map[int]*doublyLinkedNode //hashtable,find node in O(1)
}

// Constructor creates a cache object
func Constructor(capacity int) LRUCache {
	return LRUCache{
		len:   0,
		cap:   capacity,
		first: nil,
		last:  nil,
		nodes: make(map[int]*doublyLinkedNode, capacity),
	}
}

// Get returns value by key
func (l *LRUCache) Get(key int) int {
	if node, ok := l.nodes[key]; ok {
		//key exist
		// move the node to the head of double-linked-list
		l.moveToFirst(node)
		return node.val
	}

	//key not exist,return -1
	return -1
}

// Put puts key-value pair into LRUCache
func (l *LRUCache) Put(key int, value int) {
	if node, ok := l.nodes[key]; ok {
		//update value of old node
		node.val = value
		// move the node to the head of double-linked-list
		l.moveToFirst(node)
	} else {
		if l.len == l.cap {
			delete(l.nodes, l.last.key)
			l.removeLast()
		} else {
			l.len++
		}
		node := &doublyLinkedNode{
			prev: nil,
			next: nil,
			key:  key,
			val:  value,
		}
		l.nodes[key] = node
		l.insertToFirst(node)
	}
}

func (l *LRUCache) removeLast() {
	if l.last.prev != nil {
		//双向链表长度>1
		l.last.prev.next = nil
	} else {
		//双向链表长度 == 1,first == last
		l.first = nil
	}
	l.last = l.last.prev
}
func (l *LRUCache) moveToFirst(node *doublyLinkedNode) {
	switch node {
	case l.first:
		return
	case l.last:
		l.removeLast()
	default:
		//在双向链中，删除 node 节点
		node.prev.next = node.next
		node.next.prev = node.prev
	}
	// 策略是
	// 如果是移动node
	// 先删除,再插入
	l.insertToFirst(node)
}

func (l *LRUCache) insertToFirst(node *doublyLinkedNode) {
	if l.last == nil {
		//空的双向链表
		l.last = node
	} else {
		node.next = l.first
		l.first.prev = node
	}
	l.first = node
}

func main() {
	lruCache := Constructor(10)
	lruCache.Put(1, 1)
	lruCache.Put(2, 2)
	lruCache.Put(3, 3)
	lruCache.Put(4, 4)
	lruCache.Get(3)

}
