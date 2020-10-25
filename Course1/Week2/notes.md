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

`new()` creates a variable and returns a pointer to the variable; variable is initialized to 0.

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
