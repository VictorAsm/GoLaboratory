package main

import (
	"GoAlgorithmLab/DataStructures/circularlinkedlist"
	"fmt"
)

func main() {
	list := circularlinkedlist.NewCircularLinkedList()

	list.Insert(10)
	list.Insert(20)
	list.Insert(30)

	fmt.Println("Lista después de insertar:")
	list.Print()

	list.Delete(20)
	fmt.Println("Lista después de eliminar 20:")
	list.Print()

	fmt.Println("Es circular?", list.IsCircular())
}
