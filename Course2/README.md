# Week 1

## Function Parameters and Return Values

**Parameters** are listed in parenthesis after function name.

```
func foo(x, y int) {
  fmt.Printf(x * y)
}
```

**Return values** are after pramaters in declaration.

```
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

- Passed arguments are copied to parameters
- Modifying parameters has no effect outside the function
- Data Encapsulation (limits the propagation of errors)
- Needs copying time

Call by Reference

- Pass a pointer as an argument

```
func foo(y *int) {
  *y m= *y + 1
}

func main() {
  x := 2
  foo(&x)
  fmt.Print(x)
}
```

- Advantage: no copying time
- Disadvantage: does not follow data encapsulation

## Passing Arrays & Slices

Passing Array Arguments

- Array arguments are copied
- Arrays can be big, so this can be a problem
- Pass array pointers instead (unnecessary in Golang)

```
func foo (x *[3]int) int {
  (*x)[0] = (*x)[0] + 1
}

func main() {
  a := [3]int{1, 2, 3}
  foo(&a)
  fmt.Print(a)
}
```

- Pass slices instead
  - Slices contain a pointer to the array --> Passing a slice copies the pointer

```
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

- Variables can be declared with a function type
- Can be created dynamically
- Can be passed as argumetns and returned as values
- Can be stored in data structures

```
var funcVar func(int) int

func incFn(x int)  int  {
  return x + 1
}

func main() {
  funcVar = incFn
  fmt.Print(funcvar(1))
}
```

```
func applyIt(afunct func (int) int, val int) int  {
  return afunct(val)
}
```

Anonymous functions

```
func main() {
  v := applyIt(func (x int) int {return x + 1}, 2)
  fmt.Println(v)
}
```

## Returning Functions

Can be used to create a function with controllable parameters.

```
func MakeDistOrigin (o_x, o_y float64) func (float64, float64) float64 {
  fn := func(x, y float64) float64 {
    return math.Sqrt(math.Pow(x - o_x, 2) + math.Pow(y - o_y, 2))
  }
  return fn
}
```

**_Environment_** of a function: the set of all names that are valid inside a function

- Lexical Scoping --> environment includes names defined in the block where the function is defined

**_Closure_**: function + its environment

- When functions are passed/returned, their environment comes with them

## Variadic and Differed

Variable Argument Number (in Variadic Functions)

- Treated as a slice inside function
- Can pass a slice to a variadic function as well

```
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

- Call can be deferred until the surrounding function completes
- Typically used for cleanup activities

```
func main() {
  defer fmt.Println("Bye!")
  fmt.Println("Hello!")
}
```

- Arguments of a deferred call are evaluated immediately

# Week 3

## Support for Classes

Go has no `class` keyword.

- Method has a **_receiver type_** that it is associated with
- Use dot notation to call the method

```
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

```
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

- Can define public functions to allow access to hidden data

```
package data

var x int = 1
func PrintX() {fmt.Println(x)}
```

```
package main

import "data"

func main() {
  data.PrintX()
}
```

- For structs, hide fields of structs by starting field name with a lower-case letter

```
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

```
func (p *Point) OffsetX(v float64) {
  p.x = p.x + v
}
```

For pointer receivers, dereferencing is not needed inside the method body, unlike how it is normally done with pointers. Additionally, there is no need to reference in when the method is called (i.e. no `&p`).

Good programming practice is for all methods for a type to have pointer receivers or for all methods for a type to have non-pointer receivers. Mixing can result in confusion.

# Week 4

## Polymorphism

Ability for an object to have different "forms" depending on the context.

- Identical at a high level of abstraction

Implemented with inheritance and overriding (with same signature).

Go does not have inheritance.

## Interfaces

- Set of method signatures (name, parameters, return values)
  - Implementation is NOT defined
- Used to express conceptual similarity between types
- E.g. Shape2D interface
  - All 2D shapes must have `Area()` and `Perimeter()`

A type _satisfies an interface_ if it defines all methods specified in the interface.

- Additional methods are OK
- Similar to inheritance with overriding

Defining an interface

```
type Shape2D interface {
  Area() float64
  Perimeter() float64
}

type Triangle { ... }
func (t Triangle) Area() float64 { ... }
func (t Triangle) Perimeter() float64 { ... }
```

- There is no need to state explicity that a type satisfies an interface (unlike in Java)

## Interface vs. Concrete Types

Concrete Types

- Specify the exact representation of the data and methods
- Complete method implementation is included

Interface Types

- Specifies some method signatures
- Implementations are abstracted

Interface Values

- Can be treated like other values
  - Assigned to variables
  - Passed, returned
- Interface values have two components:
  - Dynamic Type: Concrete type to which it is assigned to
  - Dynamic Value: Value of the dynamic type
- An interface can have a nil dynamic value (check whether dynamic value is nil before calling method!)

```
var s1 Speaker
var d1 *Dog // Speak() method defined in Dog
s1 = d1
```

- If interface has no dynamic type, then no methods can be called from it

## Using Interfaces

Ways to use an interface

- Need a function which takes multiple types of parameter

```
func FitInYard(s Shape2D) bool {
  if (s.Area() > 100 && s.Perimeter <> 100) {
    return true
  }
  return false
}
```

Empty Interface

- Empty interface specifies no methods
- All types satisfy the empty interface
- Use it to have a function accept any type as a parameter

```
func PrintMe(val interface{}) {
  fmt.Println(val)
}
```

## Type Assertions

Interfaces hide the differences between types. Sometimes you need to treat different types in different ways.

Type Assertions

```
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

```
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

- Many Go programs return error interface objects to indicate errors

```
type error interface {
  Error() string
}
```

- Correct operation: `error == nil`
- Incorrect operation: `Error()` prints error message
