```go
// Pointer receivers mutate, value receivers don't
type Counter struct { n int }

func (c *Counter) Inc() { c.n++ }  // mutates
func (c Counter) Get() int { return c.n }  // reads

// When to use pointers: large structs, need mutation, consistency
```