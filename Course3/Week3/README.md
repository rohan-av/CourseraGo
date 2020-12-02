# Week 3

## Goroutines

Creating a Goroutine

- One goroutine is created automatically to execute the `main()`
- Other goroutines are created using the `go` keyword

```
a = 1
go foo() // newly created sub-goroutine
a = 2
```

`foo()` does not have to completely execute before `a = 2` can happen.

Exiting a Goroutine

- A goroutine exits when its code is complete
- When the main goroutine is complete, all other goroutines exit
  - A goroutine may not complete its execution because main completes early (Early Exit)

## Basic Synchronization

- Using global events whose execution is viewed by all threads, simultaneously

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

- Sync package contains functions synchronize between goroutines
- sync.WaitGroup forces a goroutine to wait for other goroutines
- Contains an internal counter (counting semaphore)
  - Increment counter for each goroutine to wait for
  - Decrement counter when each goroutine completes
  - Waiting goroutine cannot continue until counter is 0

Main thread:

```
var wg sync.WaitGroup
wg.Add(1)
go foo(&wg)
wg.Wait()
```

Foo thread:

```
wg.Done()
```

`Add()` increments the counter
`Done()` decrements the counter
`Wait()` will block until the counter == 0

## Goroutine Communication

Goroutines usually work together to perform a bigger task (e.g. web server uses multi-threading to handle many connections, finding the product of 4 integers).

- Often need to send data to collaborate

Channels

- Transfer data between goroutines
- Channels are types
- Use `make()` to create a channel: `c := make(chan int)`
- Send and receive data using the arrow operator (`<-`)
  - Send data: `c <- 3`
  - Receive data: `x := <- c`

E.g.

```
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

- Sending blocks until data is received
- Receiving blocks until data is sent

Blocking and Synchronization

- Channel communication is synchronous
- Blocking us the same as waiting for communication

Task 1: `c <- 3`
Task 2: `<- c`

The above is another way to implement wait.

## Buffered Channels

Channel Capacity

- Channel can contain a limited number of objects
- Optional argument to make() defines channel capacity: `c := make(chan int, 3)`
- Sending only blocks if buffer is **full**
- Receiving only blocks if buffer is **empty**

Use of Buffering

- Sender and receiver do not need to operate at exactly the same speed (temporary speed mismatch is acceptable)
- On average, the speeds have to match
