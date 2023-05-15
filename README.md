# Commands

>Initialize a `go.mod` file
>```go
> go mod init <module>
>```


>Compile and run code
>```go
>go run <file-name>
>```

# Data types

### Numerics
`int`  	Depends on platform: `int32` in 32 bit systems and int64 bit in 64 bit systems</br>
`int8  : byte`</br>
`int16 : short`</br>
`int32 : int`</br>
`int64 : long`</br>
`uint` only positive numbers. Be careful, if you have an `uint` and you make an operation that will be negative the number will "reset" to its maximum value!</br>
`float32` 32 bit floating-point numbers</br>
`float64` 64 bit floating-point numbers</br>
`complex64` Complex numbers with float32 floating-point numbers and imaginary parts</br>
`complex128` Complex numbers with float32 floating-point numbers and imaginary parts</br>

### Structs

>
>Similar to Java classes does not support inheritance
>```go
>type UserData struct {
>	firstName       string
>	lastName        string
>	email           string
>	numberOfTickets uint
>}
>```

>Struct methods
>```go
>func (user UserData) fullName() string {
>   return user.firstName + " " + user.lastName
>}
>```

>Struct pointers<br/>
> Go is a pass by value language, if you want to change values of the struct you must point to each reference
>```go
>func (user *UserData) changeName(firstName, lastName string) {
>   user.firstName = firstName
>   user.lastName = lastName
>}
>```

# Syntax
> Add element at the end of the slice
>```go
>append(<slice>, <element>)
>```

> Iterate elements for different data structures (not only arrays and slices)
> For arrays and slices, <b>range</b> provides the index and value for each element
>```go
>index, element := range <slice/array>
>```

> Split strings by white space in a slice
>```go
>fields := strings.Fields(element)
>```

> Create slice or maps...
>```go
>sliceInitialized := make([]string, 0 *initial slice size*)
>mapInitialized := make(map[string]string)
>```

# Packages
> Variables and functions defined outside any function, can be accessed in all other files within the same package

# Generics
> Defining the type we infer what type of generic we are going to use, the compiler will generate multiple versions of the generic function.
>```go
> func main() {
>    fmt.Println(sum(1, 2))
>    fmt.Println(sum[float64](1.3, 2.2))
> }
>
> func sum[T int | float64](a, b T) T {
>     return a + b
> }
>```

> The "~" will let you use also underlying variables like `type price int`
> ```go
> func sum[T int | float64](a, b T) T {
>     return a + b
> }
> ```

# Documentation
[Formatting expressions](https://pkg.go.dev/fmt#hdr-Printing)<br/>
[Strconv](https://pkg.go.dev/strconv)<br/>
[Rand](https://pkg.go.dev/math/rand)<br/>
[How to Write Go Code](https://go.dev/doc/code#Overview)<br/>
[Effective Go](https://go.dev/doc/effective_go)<br/>
