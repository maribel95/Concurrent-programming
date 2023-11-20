# Concurrent-programming

Concurrent programming is a programming paradigm that deals with the simultaneous execution of multiple tasks or processes that can occur in an overlapping manner in time. Instead of running a program sequentially, where each instruction is executed one after the other, concurrent programming allows various parts of the program to run in parallel.

In this repository you can find three different works that solve different "toy" problems related to this concept of concurrent programming. Concepts are internalized through different levels of abstraction.


Different programming languages have been used, such as Java, Ada and Go.

## P1: The problem of the mixed bathroom with traffic lights.

There are 6 women and 6 men working in a law firm and there is only one mixed bathroom. The director has set some rules for its use:
-  Women and men cannot be in the bathroom at the same time.
-  At most there can be 3 women or 3 men at the same time.
-  They go to the bathroom twice during the working day.

The simulation has been programmed using Java and using only semaphores, both counters and binaries, as synchronization tools.
Women and men have been programmed as concurrent processes, they have been assigned an identifier which will be a string of characters. After leaving the bathroom the second time the processes end.
It must be guaranteed that there is no starvation (despite the completion of the processes). An incorrect simulation would be, for example, if all the women went there first and then all the men (or vice versa).

The problem has been solved in stages:
1. Launch the processes and have them end with two start and end messages
2. Access the bathroom with mutual exclusion, only one process at a time, empty bathroom control
3. Schedule the groups of 3 (men and women indifferent)
4. Separate men from women and maintenance of meters
5. Adjust simulation with waiting and message interleaving control








