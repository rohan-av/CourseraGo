# Week 1

## Why Use Concurrency?

### Parallel Execution

Two programs execute in parallel if they execute at exactly the same time

Why use parallel execution?

- Tasks may complete more quickly
- Some tasks must be performed sequentially
- Some tasks are parallelizable and some are not

### Von Neumann Bottleneck

We can achieve speedup without Parallelism.

- Design faster processors

Von Neumann Bottleneck

- CPU needs to read and modify memory
- Memory access is always slower than the CPU
- Memory speedup is not as fast as clock rate speedup: **_Von Neumann Bottleneck_**
- Can be resolved with more on-chip cache, until now

Moore's Law

- Predicted that transistor density would double every two years
- Smaller transistors switch faster
- Exponential increase in density led to exponential increase in speed
- Not valid anymore

### Power Wall

Transistors comsume power when they switch

- Increasing transistor density leads to increased power consumption
- High power leads to high temperature

Dynamic Power

- P = aCFV^2
- a is percent of time switching
- C is capacitance
- F is the clock frequency
- V is voltage swing (from low to high)

Dennard Scaling

- Voltage should scale with transistor size
- Keeps power consumption, and temperature, low
- **Problem #1**: Voltage cannot go too low
  - Must stay above threshold voltage
  - Noise problems occur
- **Problem #2**: Does not consider _leakage power_
- Dennard Scaling can't continue

To improve performance, designers cannot increase F, but increase number of cores. In order to exploit multi-core systems, parallel execution is necessary.

## Concurrent vs. Parallel

Concurrent executionnis not necessarily the same as parallel execution.
**_Concurrent_**: start and end times overlap
**_Parallel_**: execute at exactly the same time

Concurrent tasks, unlike parallel tasks, may be executed on the same hardware. Mapping from tasks to hardware is not directly controlled by the programmer (at least not in Go). Programmer determines which tasks can be executed in parallel.

**_Concurrent Programming_** is where the programmer defines the possible concurrency.

Hiding Latency

- Concurrency improves performance, even without parallelism
- Tasks must periodically wait for something (e.g. wait for memory)
- Other concurrent tasks can operate while one task is waiting
