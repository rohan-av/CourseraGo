# Week 4: Synchronized Communication

## Blocking on Channels

Iterating Through a Channel

```
for i := range c {
  fmt.Println(i)
}
```

Iterates until sender calls `close(c)`.

Receiving from Multiple Goroutines

- Multiple channels may be used to receive from multiple sources
- One option: Read sequentially

```
a := <- c1
b := <- c2
fmt.Println(a*b) // waits on both channels
```

## Select

Select Statement

- May have a choice of which data to use (first-come first-served)
- Use the `select` statement to wait on the first data from a set of channels

```
select {
  case a = <- c1:
    fmt.println(a)
  case b = <- c2:
    fmt.Println(b)
}
```

- May select either send or receive operations

```
select {
  case a = <- inchan:
    fmt.Println("Received a")
  case outchan <- b:
    fmt.Println("Sent b")
}
```

Select with an Abort Channel

- Use select with a separate abort channel
- May want to receive data until an abort signal is received

Default Select

- May want a default operation to avoid blocking

```
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

- Access to shared variables cannot be interleaved; should be mutually exclusive

sync.Mutex

- A Mutex ensures mutual exclusion
- Uses a _binary semaphore_
- `Lock()` method puts the flag up (shared variable in use) at the beginning of the critical region
- `Lock()` blocks until the flag is put down if lock is already taken by a goroutine
- `Unlock()` method puts the flag down (done using shared variable)
- When `Unlock()` is called, a blocked `Lock()` can proceed

```
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

- Initialization must happen once and before everything else
- How do you perform initialization when there are multiple goroutines?
  - Could perform initialization before starting the goroutines
  - sync.Once

sync.Once

- Has one method: `once.Do(f)`
- Function `f` is executed only one time (even if it called in multiple goroutines)
- All calls to `once.Do()` block until the first returns (ensured that initialization executes first)

```
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

```
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

- Golang runtime automatically detects when all goroutines are deadlocked
- Cannot detect when a subset of goroutines are deadlocked

## Dining Philosophers

```
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

- No deadlock, but philosopher 4 may starve
