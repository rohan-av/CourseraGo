# COURSE 1

# Week 2

## Pointers

-   & operator returns the address of a variable/function
-   \* operator returns data at an address (dereferencing)

```Go
var x int = 1
vay y int
var ip *int // ip is pointer to int

ip = &x
y = *ip (y is now equal to the data at the address that IP is pointing to)
```

`new()`: creates a variable and returns a pointer to the variable; variable is initialized to 0

```Go
ptr := new(int)
*ptr = 3
```

## Variable Scope

**Block**: a sequence of declarations and statements within matching brackets {}

-   including function definitions

Hierarchy of _implicit blocks_ also

Examples:

-   universe block - all Go source
-   package block - all source in a package
-   file block - all source in a file
-   "if", "for", "switch" and the clauses in "switch" or "select"

Go is lexically scoped using blocks.

## Deallocating Memory & Garbage Collection

Compiler determines whether to store memory on the heap or the stack.
Go has garbage collection built into the compiler.
Active garbage collection done in the background.

## Comments, Printing, Integers

Comments: `// This is a comment`

```Go
/* this is a
multi-line comment */
```

Printing: `fmt.Printf()`

-   Format strings
    -   E.g. `fmt.Printf("Hi %s", x)`

Integers

-   Generic int declaration: `var x int`
-   Different lengths and signs: `int8, int16, int32, int64, uint8, uint16, uint32, uint64`
-   Binary operators (same as other languages)

## Ints, Floats, Strings

Type Conversions

```Go
var x int32 = 1
vay y int16 = 2
x = int32(y)
```

Floating Points

-   `float32`: ~6 digits of precision
-   `float64`: ~15 digits of precision

Complex Numbers

-   `complex128` etc.
-   `var z complex128 = complex(2,3)`

Strings

-   Default is UTF-8
-   code point: rune
-   Strings are read-only

## String Packages

Unicode Package

-   Runes are divided into many different categories
-   Provides a set of functions
    -   `IsDigit(r rune)`
    -   `IsSpace(r rune)`
    -   `IsLetter(r rune)`
    -   `IsLower(r rune)`
    -   `IsPunct(r rune)`
    -   `ToUpper(r rune)`
    -   `ToLower(r rune)`

Strings Package

-   String search functions
    -   `Compare(a, b)`: returns an integer comparing two strings lexicographically (-1 if a < b)
    -   `Contains(s, substr)`
    -   `HasPrefix(s, prefix)`
    -   `Index(s, substr)`: returns the index of the first instance of substr in s
-   String Manipulation
    -   `Replace(s, old, new, n)`
    -   `ToLower(s)`
    -   `ToUpper(s)`
    -   `TrimSpace(s)`

Strconv Package

-   `Atoi(s)`: converts s to int
-   `Itoa(d)`: converts d to string
-   `FormatFloat(f, fmt, prec, bitSize)`: converts floating point number to a string
-   `ParseFloat(s, bitSize)`: converts string to a floating point number

## Constants

Type is inferred from righthand.

```Go
const x = 1.3
const (
  y = 4
  z = "Hi"
)
```

`iota`: generates a set of related but distinct constants

-   often represents a property which has several distinct possible values (e.g. days of the week, months of the year)
-   actual value is not important
-   similar to enum

```Go
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

```Go
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

```Go
var appleNum int
fmt.Printf("Number of apples?")

num, err := fmt.Scan(&appleNum)
```

# Week 3

## Arrays

-   Arrays are fixed length.
-   Elements accessed usin []
-   Elements initialized to zero value

```Go
var x [5]int
x[0] = 2
fmt.Printf(x[1]) // 0
```

**_Array Literal_**: an array wpre-defined with values

-   `var x [5]int = [...]{1, 2, 3, 4, 5}`
-   `...` to infer size

**Iterating through arrays**

```Go
x := [3]int {1, 2, 3}

for i, v range x {
    fmt.Printf("index %d, value %d", i, v)
}
```

-   `i` is index, `v` is bounded to value at that index
-   `range` returns two values

## Slices

Slices are flexible; size can be changed.
A **_slice_** is a window on an underlying array.

-   Size up to the whole area

Three Properties:

-   Pointer to start of the slice
-   Length: number of elements in the slice
-   Capacity: max number of elements
    -   From start of slice to end of array

```Go
arr := [...]string{"a","b","c","d","e","f","g"}
s1 := arr[1:3] // elements 1 and 2
s2 := arr[2:5]
```

`len()` returns length. `cap()` returns capacity.

Accessing Slices

-   Writing to a slice changes the underlying array
-   Overlapping slices refer to the same array elements
    -   Changes to one slice will affect any overlapping slice

Slice Literals

-   Creates the underlying array and creates a slice to reference it
-   Slice points to the start of the array, length is capacity

```Go
sli := []int{1, 2, 3} // empty brackets --> slice
```

## Variable Slices

`make()` creates a slice (and array)

-   2-argument version: type and length/capacity
    -   `sli = make([]int, 10)`
-   3-argument version: type, length, capacity
    -   `sli = make([]int, 10, 15)`

`append()` increases the size of the slice

-   Adds elements to the end of the slice
-   Inserts into underlying array
-   Increases size of array if necessary

## Maps

Go's implementation of a hash table

```Go
var idMap map[string]int // map[key]value
idMap = make(map[string]int)
```

Map Literal

-   `idMap := map[string]int{"joe": 123}`

Accessing Maps

-   `fmt.Println(idMap["joe"])`
-   Adding a pair: `idMap["jane"] = 456`
-   Deleting a pair: `delete(idMap, "joe")`

Other Map Functions

-   two-value assignment: `id, p := idMap["joe"]`
    -   `p` is a boolean that is `true` if key is present
-   `len()`

Iterating through a Map

```Go
for key, val := range idMap {
  fmt.Println(key, val)
}
```

## Structs

```Go
type Person struct {
  name string
  addr string
  phone string
}

var p1 Person
```

Accessing Struct Fields

```Go
p1.name = "joe"
x= p1.addr
```

Initializing Structs

-   `p1 := new(Person)`
-   Struct literal: `p1 := Person{name: "joe", addr: "a st.", phone: "123"}`

## Other notes

`"sort"` library can be used for sorts.

# Week 4

## RFCs: Requests for Comments

**RFCs**: Definitions of Internet protocols and formats
E.g.

-   HTML
-   URI (Uniform Resource Identifier)
-   HTTP
-   JSON

Golang has protocol pakcages to decode and encode these protocols, e.g.

-   `"net/http"` : `http.Get(www.uci.edu)`
-   `"net"`: TCP/IP and socket programming (`net.Dial("tcp","uci.edu:80")`)

JSON

-   JavaScript Object Notation, RFC 7159
-   Format to represent structured information
-   Attribute-value pairs
    -   struct/map

## JSON

Properties

-   All Unicode
-   Human-readable
-   Fairly compact representation
-   Types can be combined recursively

JSON Marshalling: generating JSON representation from an object

```Go
p1 := Person(name: "Joe", addr: "A St.", phone: "123")

barr, err := json.Marshal(p1)
```

`Marshal()` returns JSON representation as `[]byte`.

```Go
var p2 Person

err := json.Unmarshal(barr, &p2)
```

`Unmarshal()` converts a JSON `[]byte` into a Go object.

## File Access (`ioutil`)

-   Linear access, not random access
    -   Mechanical delay
-   Basic operations
    -   Open
    -   Read
    -   Write
    -   Close
    -   Seek (move read/write head)

**`ioutil` File Read**

```Go
dat, e := ioutil.ReadFile("test.txt")
```

-   `dat` is `[]byte` filled with contents of entire file
-   Explicit open/close not needed
-   Large files cause a problem

**`ioutil` File Write**

```Go
dat = "Hello, world"

err := ioutil.WriteFile("outfile.txt", dat, 0777)
```

-   Creates a file
-   Third argument is permission (Unix-style permission bytes)
    -   `0777`: universal permission for read/write

## File Access (`os`)

`os.Open()`: opens a file (returns a file descriptor)
`os.Close()`: closes a file
`Read()`: reads from a file into a `[]byte`

-   Controls the amount read (fills the `[]byte`)
-   returns no. of bytes read

```Go
f, err := os.Open("dt.txt")
barr := make([]byte, 10)
nb, err := f.Read(barr)
f.close()
```

`Write()`: writes a `[]byte`
`WriteString()`: writes a string (any Unicode sequence)

```Go
f, err := os.Create("outfile.txt")

barr := []byte{1, 2, 3}
nb, err := f.Write(barr)
nb, err := f.WriteString("Hi")
```

# COURSE 2

# Week 1

## Function Parameters and Return Values

**Parameters** are listed in parenthesis after function name.

```Go
func foo(x, y int) {
  fmt.Printf(x * y)
}
```

**Return values** are after pramaters in declaration.

```Go
func foo(x int) int {
  return x + 1
}

func foo2(x int) (int, int) {
  return x, x + 1
}

a, b := foo2(3)
```

## Call by Value / References

Call by Value

-   Passed arguments are copied to parameters
-   Modifying parameters has no effect outside the function
-   Data Encapsulation (limits the propagation of errors)
-   Needs copying time

Call by Reference

-   Pass a pointer as an argument

```Go
func foo(y *int) {
  *y m= *y + 1
}

func main() {
  x := 2
  foo(&x)
  fmt.Print(x)
}
```

-   Advantage: no copying time
-   Disadvantage: does not follow data encapsulation

## Passing Arrays & Slices

Passing Array Arguments

-   Array arguments are copied
-   Arrays can be big, so this can be a problem
-   Pass array pointers instead (unnecessary in Golang)

```Go
func foo (x *[3]int) int {
  (*x)[0] = (*x)[0] + 1
}

func main() {
  a := [3]int{1, 2, 3}
  foo(&a)
  fmt.Print(a)
}
```

-   Pass slices instead
    -   Slices contain a pointer to the array --> Passing a slice copies the pointer

```Go
func foo(sli []int) int {
  sli[0] = sli[0] + 1
}

func main() {
  a := []int{1, 2, 3}
  foo(a)
  fmt.Print(a)
}
```

# Week 2

## First-Class Values

Functions are **_first-class_**: being able to treat a function like any other type\

-   Variables can be declared with a function type
-   Can be created dynamically
-   Can be passed as argumetns and returned as values
-   Can be stored in data structures

```Go
var funcVar func(int) int

func incFn(x int)  int  {
  return x + 1
}

func main() {
  funcVar = incFn
  fmt.Print(funcvar(1))
}
```

```Go
func applyIt(afunct func (int) int, val int) int  {
  return afunct(val)
}
```

Anonymous functions

```Go
func main() {
  v := applyIt(func (x int) int {return x + 1}, 2)
  fmt.Println(v)
}
```

## Returning Functions

Can be used to create a function with controllable parameters.

```Go
func MakeDistOrigin (o_x, o_y float64) func (float64, float64) float64 {
  fn := func(x, y float64) float64 {
    return math.Sqrt(math.Pow(x - o_x, 2) + math.Pow(y - o_y, 2))
  }
  return fn
}
```

**_Environment_** of a function: the set of all names that are valid inside a function

-   Lexical Scoping --> environment includes names defined in the block where the function is defined

**_Closure_**: function + its environment

-   When functions are passed/returned, their environment comes with them

## Variadic and Differed

Variable Argument Number (in Variadic Functions)

-   Treated as a slice inside function
-   Can pass a slice to a variadic function as well

```Go
func getMax(vals ...int) int {
  maxV := 1
  for _, v := range vals {
    if v > maxV {
      maxV = v
    }
  }
  return maxV
}
```

Deferred Function Calls

-   Call can be deferred until the surrounding function completes
-   Typically used for cleanup activities

```Go
func main() {
  defer fmt.Println("Bye!")
  fmt.Println("Hello!")
}
```

-   Arguments of a deferred call are evaluated immediately

# Week 3

## Support for Classes

Go has no `class` keyword.

-   Method has a **_receiver type_** that it is associated with
-   Use dot notation to call the method

```Go
type MyInt int

func (mi MyInt) Double () int { // (mi MyInt) is the receiver type
  return int(mi*2)
}

func main() {
  v := MyInt(3)
  fmt.Println(v.Double()) // Object v is an implicit argument to the method (call by value)
}
```

The receiver type can be a struct of some kind.

```Go
func (p Point) DistToOrig() {
  t := math.Pow(p.x, 2) + math.Pow(p.y, 2)
  return math.Sqrt(t)
}

func main() {
  p1 := Point(3,4)
  fmt.Println(p1.DistToOrigin())
}
```

## Encapsulation

Controlling Access

-   Can define public functions to allow access to hidden data

```Go
package data

var x int = 1
func PrintX() {fmt.Println(x)}
```

```Go
package main

import "data"

func main() {
  data.PrintX()
}
```

-   For structs, hide fields of structs by starting field name with a lower-case letter

```Go
package main
func main() {
  var p data.Point
  p.InitMe(3, 4)
  p.Scale(2)
  p.PrintMe() // 3, 4
}
```

## Point Receivers

The receiver object is implicity passed as an argument, and argument passing in Go is call by value. Hence, the method cannot directly modify the data inside the receiver object. The solution is to make the receiver type a pointer.

```Go
func (p *Point) OffsetX(v float64) {
  p.x = p.x + v
}
```

For pointer receivers, dereferencing is not needed inside the method body, unlike how it is normally done with pointers. Additionally, there is no need to reference in when the method is called (i.e. no `&p`).

Good programming practice is for all methods for a type to have pointer receivers or for all methods for a type to have non-pointer receivers. Mixing can result in confusion.

# Week 4

## Polymorphism

Ability for an object to have different "forms" depending on the context.

-   Identical at a high level of abstraction

Implemented with inheritance and overriding (with same signature).

Go does not have inheritance.

## Interfaces

-   Set of method signatures (name, parameters, return values)
    -   Implementation is NOT defined
-   Used to express conceptual similarity between types
-   E.g. Shape2D interface
    -   All 2D shapes must have `Area()` and `Perimeter()`

A type _satisfies an interface_ if it defines all methods specified in the interface.

-   Additional methods are OK
-   Similar to inheritance with overriding

Defining an interface

```Go
type Shape2D interface {
  Area() float64
  Perimeter() float64
}

type Triangle { ... }
func (t Triangle) Area() float64 { ... }
func (t Triangle) Perimeter() float64 { ... }
```

-   There is no need to state explicity that a type satisfies an interface (unlike in Java)

## Interface vs. Concrete Types

Concrete Types

-   Specify the exact representation of the data and methods
-   Complete method implementation is included

Interface Types

-   Specifies some method signatures
-   Implementations are abstracted

Interface Values

-   Can be treated like other values
    -   Assigned to variables
    -   Passed, returned
-   Interface values have two components:
    -   Dynamic Type: Concrete type to which it is assigned to
    -   Dynamic Value: Value of the dynamic type
-   An interface can have a nil dynamic value (check whether dynamic value is nil before calling method!)

```Go
var s1 Speaker
var d1 *Dog // Speak() method defined in Dog
s1 = d1
```

-   If interface has no dynamic type, then no methods can be called from it

## Using Interfaces

Ways to use an interface

-   Need a function which takes multiple types of parameter

```Go
func FitInYard(s Shape2D) bool {
  if (s.Area() > 100 && s.Perimeter <> 100) {
    return true
  }
  return false
}
```

Empty Interface

-   Empty interface specifies no methods
-   All types satisfy the empty interface
-   Use it to have a function accept any type as a parameter

```Go
func PrintMe(val interface{}) {
  fmt.Println(val)
}
```

## Type Assertions

Interfaces hide the differences between types. Sometimes you need to treat different types in different ways.

Type Assertions

```Go
func DrawShape(s Shape2D) bool {
  rect, ok := s.(Rectangle)
  if ok {
    DrawRect(rect)
  }
  tri, ok := s.(Triangle)
  if ok {
    DrawTri(tri)
  }
}
```

Type Switch

```Go
func DrawShape(s Shape2D) bool {
  switch sh := s.(type) {
  case Rectangle:
    DrawRect(sh)
  case Triangle:
    DrawTriangle(sh)
  }
}
```

## Error Handling

Error Interface

-   Many Go programs return error interface objects to indicate errors

```Go
type error interface {
  Error() string
}
```

-   Correct operation: `error == nil`
-   Incorrect operation: `Error()` prints error message

# COURSE 3

# Week 1

## Why Use Concurrency?

### Parallel Execution

Two programs execute in parallel if they execute at exactly the same time

Why use parallel execution?

-   Tasks may complete more quickly
-   Some tasks must be performed sequentially
-   Some tasks are parallelizable and some are not

### Von Neumann Bottleneck

We can achieve speedup without Parallelism.

-   Design faster processors

Von Neumann Bottleneck

-   CPU needs to read and modify memory
-   Memory access is always slower than the CPU
-   Memory speedup is not as fast as clock rate speedup: **_Von Neumann Bottleneck_**
-   Can be resolved with more on-chip cache, until now

Moore's Law

-   Predicted that transistor density would double every two years
-   Smaller transistors switch faster
-   Exponential increase in density led to exponential increase in speed
-   Not valid anymore

### Power Wall

Transistors comsume power when they switch

-   Increasing transistor density leads to increased power consumption
-   High power leads to high temperature

Dynamic Power

-   P = aCFV^2
-   a is percent of time switching
-   C is capacitance
-   F is the clock frequency
-   V is voltage swing (from low to high)

Dennard Scaling

-   Voltage should scale with transistor size
-   Keeps power consumption, and temperature, low
-   **Problem #1**: Voltage cannot go too low
    -   Must stay above threshold voltage
    -   Noise problems occur
-   **Problem #2**: Does not consider _leakage power_
-   Dennard Scaling can't continue

To improve performance, designers cannot increase F, but increase number of cores. In order to exploit multi-core systems, parallel execution is necessary.

## Concurrent vs. Parallel

Concurrent executionnis not necessarily the same as parallel execution.
**_Concurrent_**: start and end times overlap
**_Parallel_**: execute at exactly the same time

Concurrent tasks, unlike parallel tasks, may be executed on the same hardware. Mapping from tasks to hardware is not directly controlled by the programmer (at least not in Go). Programmer determines which tasks can be executed in parallel.

**_Concurrent Programming_** is where the programmer defines the possible concurrency.

Hiding Latency

-   Concurrency improves performance, even without parallelism
-   Tasks must periodically wait for something (e.g. wait for memory)
-   Other concurrent tasks can operate while one task is waiting

# Week 2: Concurrency Basics

## Processes

**_Process_**: an instance of a running program

Things unique to a process

-   Memory
    -   Virtual address space
    -   Code, stack, heap, shared libraries
-   Registers
    -   Program counter, data registers, stack pointer etc.

**_Operating Systems_** allow many processes to execute concurrently.

-   Processes are switched quickly
-   User has the impression of parallelism
-   OS must give processes fair access to resources

## Scheduling

**_Context switch_**: the act of control flow changing from one process to another

## Threads and Goroutines

Context switching between processes can be slow.

Threads vs. Processes

-   Threads share some context
-   Many threads can exist in one process
-   Stack, data registers and code unique to every thread
-   Virtual memory and file descriptors shared among all threads in a process
-   Context switching between threads is faster

**_Goroutine_**: like a thread, but in Go.

-   Many Goroutines execute within a single OS thread

**_Go Runtime Scheduler_** schedules goroutines inside an OS thread (like a little OS inside a single OS thread)

-   _Logical processor_ is mapped to a thread (number of logical processors can be decided by programmer; parallel execution is possible if more than 1 logical processor is used)

## Interleaving

Order of execution within a task is known, but the order of execution between concurrent tasks is unknown. Interleaving of instructions between different tasks is unknown.

Interleaving is happening at the machine code instruction level.

## Race Conditions

**_Race Condition_**: outcome depends on non-deterministic ordering

Programmer needs to make sure that the program output is deterministic.

Races occur due to _communication_.

Threads are largely independent but not completely independent, as they are sharing information. (e.g. Web server with one thread per client)

# Week 3

## Goroutines

Creating a Goroutine

-   One goroutine is created automatically to execute the `main()`
-   Other goroutines are created using the `go` keyword

```Go
a = 1
go foo() // newly created sub-goroutine
a = 2
```

`foo()` does not have to completely execute before `a = 2` can happen.

Exiting a Goroutine

-   A goroutine exits when its code is complete
-   When the main goroutine is complete, all other goroutines exit
    -   A goroutine may not complete its execution because main completes early (Early Exit)

## Basic Synchronization

-   Using global events whose execution is viewed by all threads, simultaneously

Task 1:

```
x = 1
x = x + 1
GLOBAL EVENT
```

Task 2:

```
if GLOBAL EVENT
  print x
```

Synchronization is used to restrict bad interleavings (at the cost of efficiency).

## Wait Groups

Sync WaitGroup

-   Sync package contains functions synchronize between goroutines
-   sync.WaitGroup forces a goroutine to wait for other goroutines
-   Contains an internal counter (counting semaphore)
    -   Increment counter for each goroutine to wait for
    -   Decrement counter when each goroutine completes
    -   Waiting goroutine cannot continue until counter is 0

Main thread:

```Go
var wg sync.WaitGroup
wg.Add(1)
go foo(&wg)
wg.Wait()
```

Foo thread:

```Go
wg.Done()
```

`Add()` increments the counter
`Done()` decrements the counter
`Wait()` will block until the counter == 0

## Goroutine Communication

Goroutines usually work together to perform a bigger task (e.g. web server uses multi-threading to handle many connections, finding the product of 4 integers).

-   Often need to send data to collaborate

Channels

-   Transfer data between goroutines
-   Channels are types
-   Use `make()` to create a channel: `c := make(chan int)`
-   Send and receive data using the arrow operator (`<-`)
    -   Send data: `c <- 3`
    -   Receive data: `x := <- c`

E.g.

```Go
func prod(v1 int, v2 int, c chan int) {
  c <- v1 * v2
}

func main() {
  c := make(chan int)
  go prod(1, 2, c)
  go prod(3, 4, c)
  a := <- c
  b := <- c
  fmt.Println(a*b)
}
```

## Blocking on Channels

Channels are, by default, unbuffered (cannot hold data in transit).

-   Sending blocks until data is received
-   Receiving blocks until data is sent

Blocking and Synchronization

-   Channel communication is synchronous
-   Blocking us the same as waiting for communication

Task 1: `c <- 3`
Task 2: `<- c`

The above is another way to implement wait.

## Buffered Channels

Channel Capacity

-   Channel can contain a limited number of objects
-   Optional argument to make() defines channel capacity: `c := make(chan int, 3)`
-   Sending only blocks if buffer is **full**
-   Receiving only blocks if buffer is **empty**

Use of Buffering

-   Sender and receiver do not need to operate at exactly the same speed (temporary speed mismatch is acceptable)
-   On average, the speeds have to match

# Week 4: Synchronized Communication

## Blocking on Channels

Iterating Through a Channel

```Go
for i := range c {
  fmt.Println(i)
}
```

Iterates until sender calls `close(c)`.

Receiving from Multiple Goroutines

-   Multiple channels may be used to receive from multiple sources
-   One option: Read sequentially

```Go
a := <- c1
b := <- c2
fmt.Println(a*b) // waits on both channels
```

## Select

Select Statement

-   May have a choice of which data to use (first-come first-served)
-   Use the `select` statement to wait on the first data from a set of channels

```Go
select {
  case a = <- c1:
    fmt.println(a)
  case b = <- c2:
    fmt.Println(b)
}
```

-   May select either send or receive operations

```Go
select {
  case a = <- inchan:
    fmt.Println("Received a")
  case outchan <- b:
    fmt.Println("Sent b")
}
```

Select with an Abort Channel

-   Use select with a separate abort channel
-   May want to receive data until an abort signal is received

Default Select

-   May want a default operation to avoid blocking

```Go
select {
  case a = <- c1:
    fmt.Println(a)
  case b = <- c2:
    fmt.Println(b)
  default:
    fmt.Println("nop")
}
```

## Mutual Exclusion

Two goroutines writing to a shared variable can interfere with each other

**_Concurrency-Safe_**: function can be invoked concurrently without interfering with other goroutines

Don't let 2 goroutines write to a shared variable at the same time!

-   Access to shared variables cannot be interleaved; should be mutually exclusive

sync.Mutex

-   A Mutex ensures mutual exclusion
-   Uses a _binary semaphore_
-   `Lock()` method puts the flag up (shared variable in use) at the beginning of the critical region
-   `Lock()` blocks until the flag is put down if lock is already taken by a goroutine
-   `Unlock()` method puts the flag down (done using shared variable)
-   When `Unlock()` is called, a blocked `Lock()` can proceed

```Go
var i int = 0
var mut sync.Mutex
func inc() {
  mut.Lock()
  i = i + 1
  mut.Unlock()
}
```

## Once Synchronization

Synchronous Initialization

-   Initialization must happen once and before everything else
-   How do you perform initialization when there are multiple goroutines?
    -   Could perform initialization before starting the goroutines
    -   sync.Once

sync.Once

-   Has one method: `once.Do(f)`
-   Function `f` is executed only one time (even if it called in multiple goroutines)
-   All calls to `once.Do()` block until the first returns (ensured that initialization executes first)

```Go
var on sync.Once
var wg sync.WaitGroup

func setup() {
  fmt.Println("Init")
}

func dostuff() {
  on.Do(setup)
  fmt.Println("hello")
  wg.Done()
}

func main() {
  wg.Add(2)
  go dostuff()
  go dostuff()
  wg.Wait()
}
```

## Deadlock

Synchronization causes the execution of different goroutines to depend on each other. Circular synchronization dependencies cause all involved goroutines to block.

```Go
func dostuff(c1 chan int, c2 chan int) {
  <- c1
  c2 <- 1
  wg.Done()
}

func main() {
  ch1 := make(chan, int)
  ch2 := make(chan, int)
  wg.Add(2)
  go dostuff(ch1, ch2)
  go dostuff(ch2, ch1)
  wg.Wait()
}
```

In the above, each goroutine is blocked on channel read, as they both rely on the other goroutine to write to the channel it is reading from. Hence, there is a deadlock.

Deadlock Detection

-   Golang runtime automatically detects when all goroutines are deadlocked
-   Cannot detect when a subset of goroutines are deadlocked

## Dining Philosophers

```Go
type ChopS struct { sync.Mutex }

type Philo struct {
  leftCS, rightCS *ChopS
}

func (p Philo) eat() {
  for {
    p.leftCS.Lock()
    p.rightCS.Lock()

    fmt.Println("eating")

    p.rightCS.Unlock()
    p.leftCS.Unlock()
  }
}

func main() {

  // Initialization (wait group excluded)

  CSticks := make([]*ChopS, 5)
  for i := 0; i < 5; i++ {
    CSticks[i] = new(ChopS)
  }

  philos := make([]*Philo, 5)
  for i := 0; i < 5; i++ {
    philos[i] = &Philo{Csticks[i], Csticks[(i+1)%5]}
  }

  // Eating

  for i := 0; i < 5; i++ {
    go philos[i].eat()
  }
}
```

Deadlock solution (by Dijkstra): each philosopher picks up lowest number chopstick first

-   No deadlock, but philosopher 4 may starve
