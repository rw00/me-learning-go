### Info

`GOPATH` is the path to the directory containing the `src`, `pkg`, and `bin` directories.
You shouldn't have to worry about it in the newer versions of `go` because it is set automatically.

### Commands
`go mod init github.com/rw00/<module>` initializes a new module in the current directory with that module path.

`go run <file.go>` executes the main in the file.

`go build` compiles the program in the current directory, producing an executable file with the same name if it's the main package.

`go install` compiles the program in the current directory and installs it in the local environment path.

`go get <package>` downloads the package and installs it in the local environment path.

### Conventions

Functions that start with a capital letter are exported and can be used outside of the package.

Don't export functions from the main package.
