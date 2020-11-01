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
