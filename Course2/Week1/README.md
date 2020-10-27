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
