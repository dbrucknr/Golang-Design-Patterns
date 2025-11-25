Chapter 1: Setup
- `go mod init github.com/dbrucknr/go-design-patterns`
- `mkdir cmd`
- `mkdir cmd/web` - this is a convention in Go for projects that serve web pages
    - If you create an API it goes in the `cmd/api` directory.
- `touch cmd/web/main.go` - Here we will create a simple web server using the net/http package.
