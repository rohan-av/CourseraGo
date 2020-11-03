# Week 2

## Pointers

- & operator returns the address of a variable/function
- \* operator returns data at an address (dereferencing)

```
var x int = 1
vay y int
var ip *int // ip is pointer to int

ip = &x
y = *ip (y is now equal to the data at the address that IP is pointing to)
```

`new()`: creates a variable and returns a pointer to the variable; variable is initialized to 0

```
ptr := new(int)
*ptr = 3
```

## Variable Scope

**Block**: a sequence of declarations and statements within matching brackets {}

- including function definitions

Hierarchy of _implicit blocks_ also

Examples:

- universe block - all Go source
- package block - all source in a package
- file block - all source in a file
- "if", "for", "switch" and the clauses in "switch" or "select"

Go is lexically scoped using blocks.

## Deallocating Memory & Garbage Collection

Compiler determines whether to store memory on the heap or the stack.
Go has garbage collection built into the compiler.
Active garbage collection done in the background.

## Comments, Printing, Integers

Comments: `// This is a comment` `/* this is a multi-line comment`
Printing: `fmt.Printf()`

- Format strings
  - E.g. `fmt.Printf("Hi %s", x)`

Integers

- Generic int declaration: `var x int`
- Different lengths and signs: `int8, int16, int32, int64, uint8, uint16, uint32, uint64`
- Binary operators (same as other languages)

## Ints, Floats, Strings

Type Conversions

```
var x int32 = 1
vay y int16 = 2
x = int32(y)
```

Floating Points

- `float32`: ~6 digits of precision
- `float64`: ~15 digits of precision

Complex Numbers

- `complex128` etc.
- `var z complex128 = complex(2,3)`

Strings

- Default is UTF-8
- code point: rune
- Strings are read-only

## String Packages

Unicode Package

- Runes are divided into many different categories
- Provides a set of functions
  - `IsDigit(r rune)`
  - `IsSpace(r rune)`
  - `IsLetter(r rune)`
  - `IsLower(r rune)`
  - `IsPunct(r rune)`
  - `ToUpper(r rune)`
  - `ToLower(r rune)`

Strings Package

- String search functions
  - `Compare(a, b)`: returns an integer comparing two strings lexicographically (-1 if a < b)
  - `Contains(s, substr)`
  - `HasPrefix(s, prefix)`
  - `Index(s, substr)`: returns the index of the first instance of substr in s
- String Manipulation
  - `Replace(s, old, new, n)`
  - `ToLower(s)`
  - `ToUpper(s)`
  - `TrimSpace(s)`

Strconv Package

- `Atoi(s)`: converts s to int
- `Itoa(d)`: converts d to string
- `FormatFloat(f, fmt, prec, bitSize)`: converts floating point number to a string
- `ParseFloat(s, bitSize)`: converts string to a floating point number

## Constants

Type is inferred from righthand.

```
const x = 1.3
const (
  y = 4
  z = "Hi"
)
```

`iota`: generates a set of related but distinct constants

- often represents a property which has several distinct possible values (e.g. days of the week, months of the year)
- actual value is not important
- similar to enum

```
type Grades int
const (
  A Grades = iota
  B
  C
  D
  F
)
```

Each constant is assigned to a unique integer.

## Control Flow

```
if x > 5 {
  // do thing
} else {
  // do other thing
}

for i:=0; i<10; i++ {
  // do thing
}

i = 0
for i < 10 {
  // do thing
  i++;
}

for {
  // do infinite thing
}

switch x {
case 1:
  // do thing
  // break is not necessary !!
case 2:
  // do thing
default:
  // do default thing
}

switch {
case x > 1:
  // do thing
case x < -1:
  // do thing
default:
  // do thing
}

```

Break and continue exist just as in other languages.

## Scan

Scan reads user input.
Takes a pointer as an argument.
Types data is writtern to pointer.
Returns number of scanned items.

```
var appleNum int
fmt.Printf("Number of apples?")

num, err := fmt.Scan(&appleNum)
```

# Week 3

## Arrays

- Arrays are fixed length.
- Elements accessed usin []
- Elements initialized to zero value

```
var x [5]int
x[0] = 2
fmt.Printf(x[1]) // 0
```

**_Array Literal_**: an array wpre-defined with values

- `var x [5]int = [...]{1, 2, 3, 4, 5}`
- `...` to infer size

**Iterating through arrays**

```
x := [3]int {1, 2, 3}

for i, v range x {
    fmt.Printf("index %d, value %d", i, v)
}
```

- `i` is index, `v` is bounded to value at that index
- `range` returns two values

## Slices

Slices are flexible; size can be changed.
A **_slice_** is a window on an underlying array.

- Size up to the whole area

Three Properties:

- Pointer to start of the slice
- Length: number of elements in the slice
- Capacity: max number of elements
  - From start of slice to end of array

```
arr := [...]string{"a","b","c","d","e","f","g"}
s1 := arr[1:3] // elements 1 and 2
s2 := arr[2:5]
```

`len()` returns length. `cap()` returns capacity.

Accessing Slices

- Writing to a slice changes the underlying array
- Overlapping slices refer to the same array elements
  - Changes to one slice will affect any overlapping slice

Slice Literals

- Creates the underlying array and creates a slice to reference it
- Slice points to the start of the array, length is capacity

```
sli := []int{1, 2, 3} // empty brackets --> slice
```

## Variable Slices

`make()` creates a slice (and array)

- 2-argument version: type and length/capacity
  - `sli = make([]int, 10)`
- 3-argument version: type, length, capacity
  - `sli = make([]int, 10, 15)`

`append()` increases the size of the slice

- Adds elements to the end of the slice
- Inserts into underlying array
- Increases size of array if necessary

## Maps

Go's implementation of a hash table

```
var idMap map[string]int // map[key]value
idMap = make(map[string]int)
```

Map Literal

- `idMap := map[string][int] { "joe": 123}`

Accessing Maps

- `fmt.Println(idMap["joe"])`
- Adding a pair: `idMap["jane"] = 456`
- Deleting a pair: `delete(idMap, "joe")`

Other Map Functions

- two-value assignment: `id, p := idMap["joe"]`
  - `p` is a boolean that is `true` if key is present
- `len()`

Iterating through a Map

```
for key, val := range idMap {
  fmt.Println(key, val)
}
```

## Structs

```
type Person struct {
  name string
  addr string
  phone string
}

var p1 Person
```

Accessing Struct Fields

```
p1.name = "joe"
x= p1.addr
```

Initializing Structs

- `p1 := new(Person)`
- Struct literal: `p1 := Person{name: "joe", addr: "a st.", phone: "123"}`

## Other notes

`"sort"` library can be used for sorts.

# Week 4

## RFCs: Requests for Comments

**RFCs**: Definitions of Internet protocols and formats
E.g.

- HTML
- URI (Uniform Resource Identifier)
- HTTP
- JSON

Golang has protocol pakcages to decode and encode these protocols, e.g.

- `"net/http"` : `http.Get(www.uci.edu)`
- `"net"`: TCP/IP and socket programming (`net.Dial("tcp","uci.edu:80")`)

JSON

- JavaScript Object Notation, RFC 7159
- Format to represent structured information
- Attribute-value pairs
  - struct/map

## JSON

Properties

- All Unicode
- Human-readable
- Fairly compact representation
- Types can be combined recursively

JSON Marshalling: generating JSON representation from an object

```
p1 := Person(name: "Joe", addr: "A St.", phone: "123")

barr, err := json.Marshal(p1)
```

`Marshal()` returns JSON representation as `[]byte`.

```
var p2 Person

err := json.Unmarshal(barr, &p2)
```

`Unmarshal()` converts a JSON `[]byte` into a Go object.

## File Access (`ioutil`)

- Linear access, not random access
  - Mechanical delay
- Basic operations
  - Open
  - Read
  - Write
  - Close
  - Seek (move read/write head)

**`ioutil` File Read**

```
dat, e := ioutil.ReadFile("test.txt")
```

- `dat` is `[]byte` filled with contents of entire file
- Explicit open/close not needed
- Large files cause a problem

**`ioutil` File Write**

```
dat = "Hello, world"

err := ioutil.WriteFile("outfile.txt", dat, 0777)
```

- Creates a file
- Third argument is permission (Unix-style permission bytes)
  - `0777`: universal permission for read/write

## File Access (`os`)

`os.Open()`: opens a file (returns a file descriptor)
`os.Close()`: closes a file
`Read()`: reads from a file into a `[]byte`

- Controls the amount read (fills the `[]byte`)
- returns no. of bytes read

```
f, err := os.Open("dt.txt")
barr := make([]byte, 10)
nb, err := f.Read(barr)
f.close()
```

`Write()`: writes a `[]byte`
`WriteString()`: writes a string (any Unicode sequence)

```
f, err := os.Create("outfile.txt")

barr := []byte{1, 2, 3}
nb, err := f.Write(barr)
nb, err := f.WriteString("Hi")
```
