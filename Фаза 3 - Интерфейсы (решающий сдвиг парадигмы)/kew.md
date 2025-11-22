Фаза 3: Интерфейсы (решающий сдвиг парадигмы)
```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

// Accept interfaces, return structs
func Process(r io.Reader) error { ... }

// Empty interface for generic before Go 1.18
func Print(v interface{}) { fmt.Println(v) }
```