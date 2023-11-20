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

A simulation example could be this:

´´´ java
MILERDA arrive at the office.
MILERDA is working.
NEREPINGU arrive at the office.
NEREPINGU is working.
ONOFRE arrive at the office.
ONOFRE is working.
PUYOLA arrive at the office.
PUYOLA is working.
JORDI FERNANDEZ arrive at the office.
JORDI FERNANDEZ is working.
GORDO MONSTRUOSO arrive at the office.
GORDO MONSTRUOSO is working.
MILERDA enters 1/2. Women at the bathroom: 1
FABADA arrive at the office.
FABADA is working.
NEREPINGU enters 1/2. Women at the bathroom: 2
MILERDA gets out of bathroom.
CALAMARDO arrive at the office.
CALAMARDO is working.
MARGABRA arrive at the office.
MARGABRA is working.
EL MATI arrive at the office.
EL MATI is working.
MARGA arrive at the office.
MARGA is working.
FRANCIS arrive at the office.
FRANCIS is working.
MILERDA is working.
NEREPINGU gets out of bathroom.
*** Bathroom is empty ***
ONOFRE enters 1/2. Men at bathroom: 1
ONOFRE gets out of bathroom.
*** Bathroom is empty ***
PUYOLA enters 1/2. Women at the bathroom: 1
NEREPINGU is working.
ONOFRE is working.
PUYOLA gets out of bathroom.
*** Bathroom is empty ***
JORDI FERNANDEZ enters 1/2. Men at bathroom: 1
GORDO MONSTRUOSO enters 1/2. Men at bathroom: 2
GORDO MONSTRUOSO gets out of bathroom.
JORDI FERNANDEZ gets out of bathroom.
*** Bathroom is empty ***
FABADA enters 1/2. Women at the bathroom: 1
PUYOLA is working.
JORDI FERNANDEZ is working.
GORDO MONSTRUOSO is working.
FABADA gets out of bathroom.
*** Bathroom is empty ***
CALAMARDO enters 1/2. Men at bathroom: 1
CALAMARDO gets out of bathroom.
*** Bathroom is empty ***
MARGABRA enters 1/2. Women at the bathroom: 1
CALAMARDO is working.
FABADA is working.
MARGABRA gets out of bathroom.
*** Bathroom is empty ***
EL MATI enters 1/2. Men at bathroom: 1
EL MATI gets out of bathroom.
*** Bathroom is empty ***
MARGA enters 1/2. Women at the bathroom: 1
MARGA gets out of bathroom.
*** Bathroom is empty ***
FRANCIS enters 1/2. Men at bathroom: 1
MARGABRA is working.
FRANCIS gets out of bathroom.
*** Bathroom is empty ***
MILERDA enters 2/2. Women at the bathroom: 1
NEREPINGU enters 2/2. Women at the bathroom: 2
MARGA is working.
EL MATI is working.
FRANCIS is working.
NEREPINGU gets out of bathroom.
MILERDA gets out of bathroom.
*** Bathroom is empty ***
ONOFRE enters 2/2. Men at bathroom: 1
ONOFRE gets out of bathroom.
*** Bathroom is empty ***
PUYOLA enters 2/2. Women at the bathroom: 1
PUYOLA gets out of bathroom.
*** Bathroom is empty ***
JORDI FERNANDEZ enters 2/2. Men at bathroom: 1
GORDO MONSTRUOSO enters 2/2. Men at bathroom: 2
CALAMARDO enters 2/2. Men at bathroom: 3
JORDI FERNANDEZ gets out of bathroom.
GORDO MONSTRUOSO gets out of bathroom.
CALAMARDO gets out of bathroom.
*** Bathroom is empty ***
FABADA enters 2/2. Women at the bathroom: 1
MARGABRA enters 2/2. Women at the bathroom: 2
MARGA enters 2/2. Women at the bathroom: 3
FABADA gets out of bathroom.
MARGABRA gets out of bathroom.
MARGA gets out of bathroom.
*** Bathroom is empty ***
EL MATI enters 2/2. Men at bathroom: 1
FRANCIS enters 2/2. Men at bathroom: 2
EL MATI gets out of bathroom.
FRANCIS gets out of bathroom.
*** Bathroom is empty ***
BUILD SUCCESSFUL (total time: 0 seconds)


´´´









