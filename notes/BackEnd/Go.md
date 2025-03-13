- Is a statically typed, compiled programming language
- Has a different variable declaration 
	```go
	var varaibleName Type= variable
	variableName := variable // only inside finctions and for declarations
	var variableName = variable // Automatically infers the type

	var variableName int = variable
	var varibleName int8 // 8 bits
	var varibleName int16
	var varibleName int32
	var varibleName uint  // only positive 
	var variableName string= "string" // Double quotes
	
```
- variable types
	- string ( double quotes only)
	- int, int8/16/32/64
	- uint, uint8/16/32/64
	- float, float8/16/32/64 
- Variable values can be updated later on

# Arrays

```go
var name [N]Type = [N]Type{ val1, val2, ..... , valN}
var name = [N]Type{ val1, val2, ..... , valN} // shorthand notation
name:= [N]Type{ val1, val2, ..... , valN}

```


# Slices

- Basically a dynamic array like JavaScript
- The difference in declaration is the length of the array are not specified
```go
var name []Type = []Type{ val1, val2, ..... }
```
**To add an element to the array**

```go
name = append(name, value)
```
- The append function returns a new slice.

**Slice ranges**
```go
name := slice[START, END]
name := slice[:] //from start to end
```