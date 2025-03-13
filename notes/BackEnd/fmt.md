# Print()

- Prints in the same line. i.e doesn't add a newline to the end of the string.
# Println()

- As the name indicates.
# Printf()

-  Prints formatted strings
- formal specifiers

| Format specifiers | description                                  |
| ----------------- | -------------------------------------------- |
| %v                | default  and for **variables**               |
| %q                | for strings and adds double quotes           |
| %T                | detects the type of variable                 |
| %0.?f             | for integers<br>? round up to decimal points |
```go
fmt.Printf("world %v",name)
```
# Sprintf()

- Save printf. returns the formatted string.
```go
var str = Sprintf("name is %v", name)
```


  

