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

```
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

-   Often need to send data to collaborate

Channels

-   Transfer data between goroutines
-   Channels are types
-   Use `make()` to create a channel: `c := make(chan int)`
-   Send and receive data using the arrow operator (`<-`)
    -   Send data: `c <- 3`
    -   Receive data: `x := <- c`

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

```
for i := range c {
  fmt.Println(i)
}
```

Iterates until sender calls `close(c)`.

Receiving from Multiple Goroutines

-   Multiple channels may be used to receive from multiple sources
-   One option: Read sequentially

```
a := <- c1
b := <- c2
fmt.Println(a*b) // waits on both channels
```

## Select

Select Statement

-   May have a choice of which data to use (first-come first-served)
-   Use the `select` statement to wait on the first data from a set of channels

```
select {
  case a = <- c1:
    fmt.println(a)
  case b = <- c2:
    fmt.Println(b)
}
```

-   May select either send or receive operations

```
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

-   Access to shared variables cannot be interleaved; should be mutually exclusive

sync.Mutex

-   A Mutex ensures mutual exclusion
-   Uses a _binary semaphore_
-   `Lock()` method puts the flag up (shared variable in use) at the beginning of the critical region
-   `Lock()` blocks until the flag is put down if lock is already taken by a goroutine
-   `Unlock()` method puts the flag down (done using shared variable)
-   When `Unlock()` is called, a blocked `Lock()` can proceed

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

-   Initialization must happen once and before everything else
-   How do you perform initialization when there are multiple goroutines?
    -   Could perform initialization before starting the goroutines
    -   sync.Once

sync.Once

-   Has one method: `once.Do(f)`
-   Function `f` is executed only one time (even if it called in multiple goroutines)
-   All calls to `once.Do()` block until the first returns (ensured that initialization executes first)

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

-   Golang runtime automatically detects when all goroutines are deadlocked
-   Cannot detect when a subset of goroutines are deadlocked

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

-   No deadlock, but philosopher 4 may starve
