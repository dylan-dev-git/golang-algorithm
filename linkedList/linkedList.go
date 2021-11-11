package main

import "fmt"

type Node struct {
	data string
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (obj *LinkedList) append(data string) {
	tailNode := Node{data,nil};
	if (obj.head == nil) {
		obj.head = &tailNode;
		obj.tail = &tailNode;
		return
	}
	obj.tail.next = &tailNode;
	obj.tail = &tailNode;
}

func (obj *LinkedList) printList() {
	listData := obj.head
	for listData != nil {
		fmt.Printf("%s -> ", listData.data)
		listData = listData.next
	}
}


func main() {
	println("Linked List")
	stList := LinkedList{}
	stList.append("aa")
	stList.append("bb")
	stList.append("cc")
	stList.append("dd")
	stList.append("ee")
	stList.printList()
}
