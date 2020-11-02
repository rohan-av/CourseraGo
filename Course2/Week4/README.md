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
