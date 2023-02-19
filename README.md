# Coordinate Converter for Go

A simple coordinate converter, intended to be a relatively simple first program to write in the Go language.

I anticpate this standalone coordinate conversion program to cover the following topics:

1. Writing simple functions to perform a calculation based on user input
2. Writing compound functions to use simple functions to return a complex value
3. Writing unit tests for that function
4. Defining data structures with limitations and error checking
5. Running this program as a [Cloud Function](https://docs.digitalocean.com/products/functions/)

In addition to these program goals, I want to learn more about the Go language like:

- Basic Go syntax like functions, variables, and math operations
- Initializing modules
- How to manage dependencies in Go
- What local development tools I need for Go
- Understanding Go's common data types

## Usage

### Conversion
```
DmsToDd(30, 15, 50)

    returns 30.26388888888889
```

### Testing
```
go test
```

