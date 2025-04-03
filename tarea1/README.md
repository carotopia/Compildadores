
# Data Structures implementadas ien Go


## Casos de prueba

### Stack
```go
stack.Push(1)
stack.Push(2)
stack.Push(30)
fmt.Println(stack.Items) // Output: [1, 2, 30]

item, _ := stack.Pop()
fmt.Println(item) // Output: 30
fmt.Println(stack.Items) // Output: [1, 2]
```

### Queue
```go
q.Enqueue(2)
q.Enqueue(3)
q.Enqueue(4)
fmt.Println(q) // Output: Queue with elements [2, 3, 4]

item, _ := q.Dequeue()
fmt.Println(item) // Output: 2
fmt.Println(q) // Output: Queue with elements [3, 4]
```

### HashMap
```go
h.Put("key1", 1)
h.Put("key2", 2)
h.Put("key3", 3)
fmt.Println(h) // Output: HashMap with stored key-value pairs

value, exists := h.Get("key1")
if exists {
   fmt.Println(value) // Output: 1
}

h.Remove("key1")
fmt.Println(h) // key1 removed from HashMap
```


