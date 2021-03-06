# How to run 
The system is hardcoded to run with 5 nodes but needs an input so run the following comands in each terminal.
All comands schould be run from root folder.

***IMPORTANT***  
Once node 0 has been activated it will wait 15 seconds before entering the critical section and starting the cirkel connection.

```golang
go run node/node.go 0

go run node/node.go 1

go run node/node.go 2

go run node/node.go 3

go run node/node.go 4
```




# DISYS 2021 Mandatory excercise 2

You have to implement distributed mutual exclusion between nodes in your distributed system. 

You can choose to implement any of the algorithms, that were discussed in lecture 7.

System Requirements:

R1: Any node can at any time decide it wants access to the Critical Section

R2: Only one node at the same time is allowed to enter the Critical Section 

R2: Every node that requests access to the Critical Section, will get access to the Critical Section (at some point in time)

Technical Requirements:

    1. Use Golang to implement the service's nodes
    2. Use gRPC for message passing between nodes
    3. Your nodes need to find each other.  For service discovery, you can choose one of the following options
        1. supply a file with  ip addresses/ports of other nodes
        2. enter ip adress/ports trough command line
        3. use the Serf package for service discovery
    4. Demonstrate that the system can be started with at least 3 nodes
    5. Demonstrate using logs,  that a node gets access to the Critical Section

Hand-in requirements:

    1. Hand in a single report in a pdf file
    2. Provide a link to a Git repo with your source code in the report
    3. Include system logs, that document the requirements are met, in the appendix of your report
