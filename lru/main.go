package main

import "fmt"

// 这个自己瞎写 👇  很垃圾
type Chain struct {
	Pre  string
	Now  string
	Next string
}

type Hash map[string]interface{}

type LRU struct {
	Chain []*Chain
	Data  Hash
}

func NewLRU(cap int) LRU {
	return LRU{
		Chain: make([]*Chain, 0, cap),
		Data:  make(map[string]interface{}, cap),
	}
}

func (l *LRU) Get(key string) interface{} {
	val, ok := l.Data[key]
	if ok {
		node := l.getNode(key)
		if node.Pre == "" {
			return val
		}
		if node.Next == "" {
			start := l.getStart()
			pre := l.getNode(node.Pre)
			pre.Next = ""
			node.Pre = ""
			node.Next = start.Now
			start.Pre = key
			return val
		}

		pre := l.getNode(node.Pre)
		next := l.getNode(node.Next)
		pre.Next = node.Next
		next.Pre = node.Pre
		start := l.getStart()
		start.Pre = key
		node.Pre = ""
		node.Now = key
		node.Next = start.Now

		return val
	}
	return nil
}

func (l *LRU) Remove(num int) {
	if num > len(l.Chain) {
		panic("太长了")
	}
	removed := l.Chain[len(l.Chain)-1-num+1:]
	l.Chain = l.Chain[:len(l.Chain)-1-num]
	for _, key := range removed {
		delete(l.Data, key.Now)
	}
}

func (l *LRU) Put(key string, data interface{}) {
	pre := ""
	preNext := ""
	if len(l.Chain) != 0 {
		pre = l.Chain[len(l.Chain)-1].Now
		preNext = key
		l.Chain[len(l.Chain)-1].Next = preNext
	}
	for _, chain := range l.Chain {
		if chain.Now == key {
			l.Data[key] = data
			return
		}
	}
	l.Chain = append(l.Chain, &Chain{
		Pre:  pre,
		Now:  key,
		Next: "",
	})
	l.Data[key] = data
}

func (l *LRU) GetAll() []interface{} {
	data := make([]interface{}, 0, len(l.Chain))

	order := make([]string, 0, len(l.Chain))
	start := l.getStart().Now
	order = append(order, start)
	for {
		node := l.getNode(start)
		start = node.Next
		if start == "" {
			break
		}
		order = append(order, start)
	}
	for _, key := range order {
		data = append(data, l.Data[key])
	}
	return data
}

func (l *LRU) getNode(now string) *Chain {
	for _, chain := range l.Chain {
		if chain.Now == now {
			return chain
		}
	}
	return nil
}
func (l *LRU) getStart() *Chain {
	for _, chain := range l.Chain {
		if chain.Pre == "" {
			return chain
		}
	}
	return nil
}
func main() {
	lru := NewLRU(10)
	lru.Put("1", 1)
	lru.Put("2", 2)
	lru.Put("3", 3)
	lru.Put("4", 4)
	lru.Put("5", 5)
	lru.Put("6", 6)
	all := lru.GetAll()
	fmt.Println(all)
	val := lru.Get("6")
	fmt.Println(val)
	all = lru.GetAll()
	fmt.Println(all)
	lru.Get("6")
	lru.Get("2")
	lru.Get("3")
	all = lru.GetAll()
	fmt.Println(all)
}
