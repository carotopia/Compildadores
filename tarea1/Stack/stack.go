package stack

import (
	"errors"
	_ "errors"
)

// Definir un nuevo tipo, esta rebanada almacenará los elemento

type Stack struct {
	Items []int
}

// Metodo push para agregar un elemento a la pila
// Recibe un valor que será el nuevo elemento a agregar
// Agrega un item al final de la rebanada

func (s *Stack) Push(item int) {
	s.Items = append(s.Items, item)
}

// Metodo para remover el último elemento
// Devuelve dos valores, el valor eliminado o error si esta vacía
// Verifica si la pila esta vacía con el metodo IsEmpty()
// Calcula el indice del último elemento
// Obtiene el último elemento y lo guarda en item
// Recorta la rebanada por el último elemento
func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")

	}
	index := len(s.Items) - 1
	item := s.Items[index]
	s.Items = s.Items[:index]
	return item, nil
}

// Verifica si el stack esta vacio
func (s *Stack) IsEmpty() bool {
	return len(s.Items) == 0
}
