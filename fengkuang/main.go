package main

/*
上机编程题
题目描述：扁平数组转Tree

给定一个扁平数组，数组内每个对象的id属性是唯一的。每个对象具有pid属性，pid属性为0表示为根节点（根节点只有一个），其它表示自己的父节点id。
编写一段程序，输入为给定的扁平数组，输出要求为一个树结构，为其中每个对象增加children数组属性（里面存放child对象）。

给定输入：
[

	{"id": 1, "name": "部门1", "pid": 0},
	{"id": 2, "name": "部门2", "pid": 1},
	{"id": 3, "name": "部门3", "pid": 1},
	{"id": 4, "name": "部门4", "pid": 3},
	{"id": 5, "name": "部门5", "pid": 4}

]

给定输出：

	{
	  "id": 1,
	  "name": "部门1",
	  "pid": 0,
	  "children": [
	    {
	      "id": 2,
	      "name": "部门2",
	      "pid": 1,
	      "children": []
	    },
	    {
	      "id": 3,
	      "name": "部门3",
	      "pid": 1,
	      "children": [
	        {
	          "id": 4,
	          "name": "部门4",
	          "pid": 3,
	          "children": [
	            {
	              "id": 5,
	              "name": "部门5",
	              "pid": 4,
	              "children": []
	            }
	          ]
	        }
	      ]
	    }
	  ]
	}
*/
func main() {
	toTree([]*Item{

		{id: 1, name: "部门1", pid: 0},
		{id: 2, name: "部门2", pid: 1},
		{id: 3, name: "部门3", pid: 1},
		{id: 4, name: "部门4", pid: 3},
		{id: 5, name: "部门5", pid: 4},
	})
}

type Item struct {
	name     string
	id       int
	pid      int
	children []*Item
}

/*
[

	{"id": 1, "name": "部门1", "pid": 0},
	{"id": 2, "name": "部门2", "pid": 1},
	{"id": 3, "name": "部门3", "pid": 1},
	{"id": 4, "name": "部门4", "pid": 3},
	{"id": 5, "name": "部门5", "pid": 4}

]
*/
func toTree(items []*Item) *Item {
	head := new(Item)
	for _, item := range items {
		if item.pid == 0 {
			head = item
			break
		}
	}
	if head == nil {
		panic("吊策划")
	}

	return head
}

func getChildren(items []*Item, pid int) {
	res := make([]*Item, 0)
	for i, item := range items {
		if item.id == pid {
			res = append(res, item)
		}

	}
}
