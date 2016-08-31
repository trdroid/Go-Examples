# Arrays

**Properties**

* Fixed length
* Homogeneous sequence of elements (all elements are of the same type)

**Declaration**

```go
var myarray [5]int    //array of 5 integers
```

**Initialization**



```go
var myarray [5]int = [5]int{1, 2, 3, 4, 5}
```

**Accessing elements**

```go
var myarray [5]int    //array of 5 integers

fmt.Println(myarray[0])
fmt.Println(myarray[len(myarray) - 1])
```

**Iteration**



**Supporting functions**

*len*: returns the count of the number of elements in the array
