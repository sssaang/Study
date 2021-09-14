# Ch.3 Role, Responsibility and Collaboration

* The core of Object-oriented paradigm is role, responsiblity and collaboration among objects.
* The purpose of oop is to create the community of collaborating objects.


## Collaboration
* The flow of an application logic that follows oo paradigm is rather controlled by different objects than by a single object
* Mutual interactions among objects to implement features is called collaboration
* A single object must be independent and autonomous in terms of its responsiblity
* An object send message (calling a method of other object in its context) to one another in order to perform jobs that is out of its boundary
* The object that receives the message (whose method is called by one another) responds by executing its implementation of the method
* In order to confirm an object's behavior, the understaning of the entire context (the illustration of collaboration) and interations between each entitiy (object) must precede


## Responsibility
* The responsibility of an object is decided by "what it knows" and "what it does"

### What it does
* creating objects or compute something
* initiating other objects' behavior
* controlling and manipulating other objects' behavior

### What it knows
* private informations
* informations of related objects
* what it induces or computes

* an object must know which and how other object that should help itself as how it behaves is constrained by its responsibility
* if an object has too small responsiblity, it cannot perform what it should do
* if an object has responsiblity that is too big for a single object, it is hard to maintain while increasing the risk of other objects' dependencies on itself

### delegating responsibility
* the most adequate way of creating an independent, autonomous object is to delegate a responsibility along with related information to it
* we call such object an information expert 

### Responsibility-Driven Design
* find responsility and adequate object to take the responsibility
1. know what system requires to implement required features (responsibility)
2. split the responsiblity into smaller pieces
3. find and create an object that could take the atomic responsibility
4. while an object needs the help from outside of its boundary, create another object to delegate the responsibility
5. following above makes object collaboration

### message decides an object
* recognizing message must precede the object creation

1. minimal interface 
* because what an object should know is decided by what other objects need it to do, recognizing their message helps to set minimal interface of the object
2. abstract interface
* recognizing the message helps to approach with "what" an object should do rather than "how" which helps to encapsulate the details of the creating object

### behavious decide the state
* the sole purpose of the existence of an object is the fact that other objects need it
* the state of an object must be decided by the behavior of the object 
* if behavior of an object is followed by the state (internal logic) it is likely to violates encapsulation (Data-Driven Design)
a. if how the object behaves is determined by what it has or know rather than what other objects require it to do, it is highly likely to lose context of collaboration


## Role
* role is a set of responsibility of objects
* a role can abstract similar responsibilities of objects, removing the possible room for logic duplication
* such effect could help to scale the responsibiliy (new features for example) within the realm of an role
