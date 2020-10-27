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
type struct Person {
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
