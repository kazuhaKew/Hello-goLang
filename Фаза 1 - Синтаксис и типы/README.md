```go
// Variables, zero values, short declarations
var name string  // ""
age := 25

// Slices vs arrays
arr := [3]int{1,2,3}        // fixed size
slice := []int{1,2,3}       // dynamic
slice = append(slice, 4)

// Maps
m := make(map[string]int)
m["key"] = 42

// Structs with embedding
type Base struct { ID int }
type User struct {
    Base
    Name string
}
```