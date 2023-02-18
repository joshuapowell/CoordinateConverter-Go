# CoordinateConverter-Go

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

## Learnings
Things that I've learned along the way and had to figure out

- [initializing a module](https://go.dev/doc/modules/managing-dependencies#naming_module)
- [file naming conventions](https://medium.com/@kdnotes/golang-naming-rules-and-conventions-8efeecd23b68)
- a variety of golang.org/x/tools/cmd/* tools needed to be installed in VSCode along the way
- I had to look up [Go Data Types](https://www.tutorialspoint.com/go/go_data_types.htm) when I wanted to figure out how to return my decimal degree. Then after that I had to figure out what level of precision I wanted my method to support, so I found this [table](http://mysql.rjweb.org/doc.php/latlng).
- I ran into a problem when setting my variable and found that [I was "doing it wrong"](https://deepsource.io/gh/tsenart/vegeta/issue/SCC-S1021/occurrences)
- Then I remembered that [I should also be validating user input](https://medium.com/@apzuk3/input-validation-in-golang-bc24cdec1835) prior to just passing it into my function
- I had to [lookup](https://gobyexample.com/structs) how to use the `struct` properly 
- I had to find a tutorial on [writing and running unit tests](https://www.digitalocean.com/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package) in go
- I wanted to limit the minimum and maximum values of a `struct` (data structure), I found this [StackOverflow](JR102339)


## Questions
- [ ] What PEP8 and PEP257 equivalent linters exist for Go?
- [ ] How can I call this module from a separate HTTP API?
- [ ] How do I handle dependencies (e.g., go mod tidy)?