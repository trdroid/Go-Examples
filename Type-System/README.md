# Type System

Go is a statically typed language.

### Types 

* Basic Types
* Aggregate Types
* Reference Types
* Interface Types

### Basic Types

**Numeric Types**

* Integers
* Floating-Point numbers
* Complex numbers

**Integers**

Go provides both signed and unsigned integers. 

* int8
* int16
* int32
* int64
* uint8
* uint16
* uint32
* uint64


**Floating-Point Numbers**

Go provides the following floating point numbers that are represented by IEEE 754 standard implemented by all modern CPUs

* float32
* float64

### User defined types

Go allows custom types to be declared using 

* struct

**struct**

A *struct* type is declared by composing a set of fields, with each field declared with a built-in type or another user defined type. 

```go
type person struct {
  name string
  email string
  student bool
}
```
