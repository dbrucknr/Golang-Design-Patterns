# Golang-Design-Patterns

## Notes:
- If you encounter a `&` prefix to a variable, it indicates the creation of a new pointer to that variable's location in memory.
  - In Go, the `&` is the **address-of** operator, which returns the memory address of a variable

```go
x := 10
p := &x // p now holds the memory address of x
```

- If you encounter a `*` prefix to a variable, it indicates the use of a pointer to access the value it points to.
- When used in an expression (e.g., *p), it retrieves the value stored at the address the pointer holds.
- When used in a type declaration (e.g., var p *int), it indicates that the variable is a pointer to that type.
  - In Go, the `*` is the **dereference** operator, which returns the value stored at the memory address pointed to by a pointer.


```go
x := 10
p := &x
fmt.Println(*p) // prints 10
*p = 20         // updates x through the pointer

```

- To run: `go run ./cmd/web`
  - To make use of template cache: `go run ./cmd/web -use-cache`
- Visit: http://localhost:4000/

- `go test -v -count=1 .` (Run tests without a cache)
- `go test -v .`
