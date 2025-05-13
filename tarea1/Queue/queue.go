package queue

import (
	"errors"
)

// Queue Definition (FIFO)
type Queue struct {
	items []int
}

// Agregar un nuevo item Enqueue
// Haciendo una rebanada de enteros que ae agregan al principio
func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// Eliminar un item Deuque, el primer elemento de la cola
// Devolver el elemento eliminado o error si esta vacía
func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("empty queue")
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

// Verificar si la lista esta vacía y devolver un error
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}
