# Ch.5 Responsibility Delegation

- object architecture build on responsibility helps to bring about decoupling, high cohesiveness, and encapsulation

## General Responsibility Assignment Software Pattern (GRASP)

### Start from domain

- it is always recommended to draw out domain map of the system to be implemented
- domain concepts do not have to be precise or perfect in the initial state
- rather they should focus on providing relevant information and insight on types and relationship of objects
- domain concepts should convey what kinds of domain exists and how each domain collaborates

### Information Expert

- first, we should align responsibilities with the features of the system
- An information expert refers to an object which knows enough to fulfill the responsibility (feature)
- assigning responsiblity to an object creates an interface of the object which can be used throughout the system

### Creator

- an object that knows enough about specific object and has initimate relationship should be the creator of the very object

## Prevention from variations

- class should be splitted with following conditions

1. if class is subjected to changes with multiple reasons
2. if class contains member variables that are not related to each other (low cohesion)
3. if class has multiple type

- splitting class means splitting responsibility, decreasing the complexity of the code and places that needs to be updated along with the change
