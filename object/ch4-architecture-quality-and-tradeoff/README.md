# Ch.4 Architecture Quality and Trade-off

- the purpose of architecture lies on the needs of change and change always create cost
- Great architecture allows us to embrace changes in relatively lower cost
- such architecture is built upon high cohesion and lowly coupled designed
- high cohesion && decoupling requires focus on the behavior of an object rather than its state
- an attempt to perceive an object as a set of data causes us to disclose the internal logic of the object, violating encapsulation

## Data-Driven Design

- data driven design focuses on what an object should have (state of an object) before confirming the role of an object
- this brings about an emphasis on implementation (internal logic), causing the object vulnerable to change
- state focused object is likely to become intricate and causes highly coupled objects as state is, by its nature, subjected to frequent change

## Trade-Off

### Violation of encapsulation

- data-driven design does not care about what object does; rather it discloses all of its detail through its interface
- data-driven design causes an object to make its interface complicated by disclosing all of its data

### Highly coupled objects

- data-driven design induces us to spread responsibility (manipulation of data) of an object to other objects that use the object
- this causes many rooms for change in unexpected places when the very object requires a change
- such makes us to fear for change and destroying the flexibility of the entire design

### low cohesion

- if implementation of the internal logic of each object is spread out to other objects, they will have myriads of unrelated logic gathered in a single method, which decreases cohesion of the code
- such case also requires changes in widespread places for a small change in an object or module as responsibilities are spread throughout
