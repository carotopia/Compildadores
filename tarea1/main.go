package main

import (
	_ "errors"
	"fmt"
	_ "slices"
	_ "tarea1/HashMap"
	hashMap "tarea1/HashMap"

	"tarea1/queue"
	"tarea1/stack"
)

func main() {
	// STACK (LIFO)

	// Crear una pila vacía
	stack := stack.Stack{}

	// Agregar elementos usando Push
	stack.Push(1)
	stack.Push(2)
	stack.Push(30)

	// Imprimir el stack
	fmt.Println("Stack:", stack.Items)

	// Pop LIFO
	// Devolver el último elemento eliminado e imprimirla
	item, err := stack.Pop()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Elemento eliminado:", item)
		fmt.Println("Stack después de Pop:", stack.Items)
	}

	fmt.Println(" ")

	// Crear una queue vacía (FIFO)

	q := queue.Queue{}

	// Agregar elementos a la queue
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)

	// Imprimri la queue
	fmt.Println("Queue:", q)

	// Eliminar el primer elmento
	item, err = q.Dequeue()
	if err != nil {
		fmt.Println(" Error", err)
	} else {
		fmt.Println("Elemento eliminado:", item)
	}
	fmt.Println("Queue después :", q)
	fmt.Println(" ")

	// Hashmap
	// Crear el hashmap "h"
	h := hashMap.HashMap{Data: make(map[string]int)}

	// Agregar llaves y valores
	h.Put("key1", 1)
	h.Put("key2", 2)
	h.Put("key3", 3)

	fmt.Println("HashMap:", h)
	value, exists := h.Get("key1")
	if exists {
		fmt.Println("Valor", value)

	} else {
		fmt.Println("No encontrado")
	}
	// Eliminar una llave
	fmt.Println("Eliminar key1")
	h.Remove("key1")
	fmt.Println("HashMap:", h)

}
