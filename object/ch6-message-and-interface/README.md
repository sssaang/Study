# Ch.6 Message and Interface

## Collaboration and message

- it is a common mistake to see an application as a set of classes instead of objects
- classes are interfaces that represent the relationships of objects
- what makes OOP great is that it captures collaboration among objects

### Client - Server model

- an object is required by other objects with specific purposes
- each object sends message to other objects when it needs something that is out of its boundary
- when an object receives messages from other objects, it performs what it is designed to do and returns what it is required for by other objects
- in terms of collaboration, an object is set of messages to be sent, and to be received
- such consideration gives us decoupled objects that could be reused throughout the system

### Mesasge and Sending Message

- an object (message sender) sends message to other objects to request something out of its boundary
- each message is comprised of receiver (the object being called), operation (method name), and parameter

### Message and Method

- when receiver is a concrete class that is derived from an interface, method being called within the message varies by the type of receiver's class
- message is an abstraction of concrete methods with specific instruction which enables us to design decoupled independent objects
- from caller(client)'s perspective, the caller does not need to know the detailed of how things are done within the method

### Public Interface and Operation

- an object should have clear distinction between it's inner area and external environment
- what an object discloses to outer environment is called public interface
- each interface is comprised of operations which represent what the object should do for other objects
- how operations are defined is called method

## Interface and Architecture Quality

- to minimize the interface, we need to focus on message to screen out unnecessary operations

### Law of Demeter

- don't talk to strangers, only talk to your immediate neighbors
- use only one dot
- class should send messages to subjects that qualifies the below

1. this object
2. method and its parameter
3. attributes of this object
4. local objects created within the method

- it is better to request the outcome (delegation) than to ask the details needed to perform task and execute the logic in a single place (tell, don't ask)

- however, forcing the law does not always bring favorable outcomes
- if we were to abstract all the operations in order to avoid chain method calls, we might end up low cohesion as objects might be comprised of irrelevant operations
- it is important to consider the trade-off of the law of demeter
