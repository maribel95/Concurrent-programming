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

<pre>
  <code>
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
  </code>
</pre>


## P2: The problem of traveling baboons with protected objects.

A gorge is a hollow or narrow and not very long passage between two mountains, or a narrow valley between steep mountains. There is a gorge with a rope crossing it from north to south. On each side live 5 baboons (a type of ape), the northern baboons and the southern baboons. Throughout the day, the baboons do the same thing 3 times, cross to the other side, go down and go back up for a walk, the ones from the north cross in a southern direction and those from the south cross in a northern direction. When a baboon starts to cross, it will surely reach the other side, that is, it will not turn back. There are two issues to consider, one, that the rope cannot support the weight of more than 3 baboons at once and two, that when one baboon crosses the rope, it cannot find another in the opposite direction. There are baboons that are faster than others and therefore can overtake them, even on the rope, this is no problem. There is also no problem if the northern or southern baboons find others in the opposite direction outside the rope. In Fig.1 you can see the walk of the northern baboons, one baboon waits, 3 are on the crossing rope and another turns around to start again.


<img width="378" alt="Captura de pantalla 2023-11-20 a las 11 19 34" src="https://github.com/maribel95/Concurrent-programming/assets/61268027/3189c0c7-3a36-4d9d-af2c-4bff319a60d1">

The simulation has been programmed using the Ada language, using protected objects as synchronization tools.
The northern and southern baboons have been programmed as concurrent tasks, they have been assigned an identifier that is sufficient to be a number and their origin (North or South). After the third round the tasks end.

## P3: The bear and bees problem with Rabbitmq and the Go client.

The simulation is that of the well-known problem described by Andrews (2000). There is a single bear who basically eats honey out of a jar and sleeps. There are N bees that carry 1 portion of honey and take it to the jar. There are H portions of honey in the jar. The bees take honey to the jar until it is full. When the jar is full, the last bee wakes up the bear. The bear eats the entire jar and while doing so the bees do not bother him. When he finishes eating he goes to sleep and the bees start all over again.
In these cases we will define the size of the jar of 10 units, the number of bees indeterminate and for the issue of finishing, we will set that the bone will eat 3 jars of honey and then end up finishing off the bees as well.

The simulation must be programmed using the Go language and using the Rabbitmq message server for communication between processes.
Both the bear and each of the bees must be programmed as Go processes to be launched from the command line, in the case of the bees they will be identified by a string on the same command line as will be the name of the bee.
There are different solutions to this problem, in this exercise it is explicitly asked not to implement the version with a Pot process that would act as a server between the bees and the bear, that is, the synchronization must be carried out by sending messages between bear and bees using Rabbitmq queues.
This problem has been studied as an example of the producer/consumer case, where bee processes are producers and the bear is the only consumer. To carry out this exercise, it may be interesting to consider the solution where the bear is the producer of the permits that allow the bees to make honey and the bees as consumers of these permits. It is not mandatory to use this idea but it can simplify your code.
To solve the termination it is important to do it in the most elegant way, that is to say, when the bear decides to end the simulation it must send the relevant messages to all the bees. As said, the number of bees is indeterminate and your code must be correct both for a simulation with 3 bees and for one with 300. That is why it is convenient to use the Fanout Exchange Routing provided by Rabbitmq.

The processes will show a message when starting and one when the simulation ends, if the bees are launched first they will wait for the bear to arrive, on the other hand if the bear arrives first it will go to sleep until the jar is full up.
When the bees are filling the jar by units they will print a message indicating what their portion of honey they have added is, the bees will execute a wait to simulate the process of making honey. The bee that puts the last portion of honey in the jar (portion 10) will print a message that wakes up the bear.
For its part, the bear when it wakes up prints the name of the bee that woke it up, then it simulates the time it eats by emptying the jar, all this time the bees are blocked, finally the bear communicates that it is leaving he goes to sleep allowing the bees to fill the jar again. When this process is repeated 3 times, a message appears that the jar is broken by terminating the bee processes and clearing the queues before finishing.
Clearing the queues is convenient while programming to not have messages in them from previous executions, in any case they can also be cleared with the interactive tool or with the Rabbitmq queue management commands.








