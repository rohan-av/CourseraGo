# Week 2: Concurrency Basics

## Processes

**_Process_**: an instance of a running program

Things unique to a process

- Memory
  - Virtual address space
  - Code, stack, heap, shared libraries
- Registers
  - Program counter, data registers, stack pointer etc.

**_Operating Systems_** allow many processes to execute concurrently.

- Processes are switched quickly
- User has the impression of parallelism
- OS must give processes fair access to resources

## Scheduling

**_Context switch_**: the act of control flow changing from one process to another

## Threads and Goroutines

Context switching between processes can be slow.

Threads vs. Processes

- Threads share some context
- Many threads can exist in one process
- Stack, data registers and code unique to every thread
- Virtual memory and file descriptors shared among all threads in a process
- Context switching between threads is faster

**_Goroutine_**: like a thread, but in Go.

- Many Goroutines execute within a single OS thread

**_Go Runtime Scheduler_** schedules goroutines inside an OS thread (like a little OS inside a single OS thread)

- _Logical processor_ is mapped to a thread (number of logical processors can be decided by programmer; parallel execution is possible if more than 1 logical processor is used)

## Interleaving

Order of execution within a task is known, but the order of execution between concurrent tasks is unknown. Interleaving of instructions between different tasks is unknown.

Interleaving is happening at the machine code instruction level.

## Race Conditions

**_Race Condition_**: outcome depends on non-deterministic ordering

Programmer needs to make sure that the program output is deterministic.

Races occur due to _communication_.

Threads are largely independent but not completely independent, as they are sharing information. (e.g. Web server with one thread per client)
