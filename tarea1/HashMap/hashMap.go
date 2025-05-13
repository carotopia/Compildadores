package hashMap

import (
	_ "errors"
)

// Definicion de Hashmap
// Implementar un mapa para hacer el par clave-valor
type HashMap struct {
	Data map[string]int
}

// Agregar une elemento al hashmap como llave y valor
// Si ya existe el valor se actualiza
func (h *HashMap) Put(key string, value int) {
	h.Data[key] = value

}

// Obtener un valor con la llave

func (h *HashMap) Get(key string) (int, bool) {
	value, exists := h.Data[key]
	return value, exists
}

// Checar si el valor existe
func (h *HashMap) HasKey(key string) bool {
	_, exists := h.Data[key]
	return exists
}

func (h *HashMap) Remove(key string) {
	delete(h.Data, key)
}
