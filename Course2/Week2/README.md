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
