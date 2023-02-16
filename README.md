# CoordinateConverter-Go

# Learnings
Things that I've learned along the way and had to figure out

- [ ] dependency management
- [ ] [initializing a module](https://go.dev/doc/modules/managing-dependencies#naming_module)
- [ ] [file naming conventions](https://medium.com/@kdnotes/golang-naming-rules-and-conventions-8efeecd23b68)
- [ ] a variety of golang.org/x/tools/cmd/* tools needed to be installed in VSCode along the way
- [ ] I had to look up [Go Data Types](https://www.tutorialspoint.com/go/go_data_types.htm) when I wanted to figure out how to return my decimal degree. Then after that I had to figure out what level of precision I wanted my method to support, so I found this [table](http://mysql.rjweb.org/doc.php/latlng).
- [ ] I ran into a problem when setting my variable and found that [I was "doing it wrong"](https://deepsource.io/gh/tsenart/vegeta/issue/SCC-S1021/occurrences)
- [ ] Then I remembered that [I should also be validating user input](https://medium.com/@apzuk3/input-validation-in-golang-bc24cdec1835) prior to just passing it into my function


# Questions
- [ ] What PEP8 and PEP257 equivalent linters exist for Go?