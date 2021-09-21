# Ch.7 Object Dicomposition

## Procedure Dicomposition

- breaks down a operation into smaller operations (top-down)
- define topmost operation then get into the smaller operations with more detail and less responsibility
- child operation must be less abstract and simpler than its parent
- problems of procedure dicomposition

1. an application is rarely comprised of a single topmost operation
2. a smaller change or addition of a feature will make a frequent change in the topmost operation
3. business logic is likely to be confined with user interface
4. the order of the logic is settled too early with the top-down composition approach
5. when data used throughout the logic is changed, it is not easy to find out side-effects

- procedure composition is best fitted for smaller, well-defined operation which is not subjected to change

## Module

- module groups implementations and discloses public interface while hiding detailed implementation
- module should hide the below

1. complexity: it is easier for clients to use the module when simple interface lies on the complex implementation
2. possible change: when implementation is leaked to the outside, the side-effect caused by a small change in the logic increases

- changes happens only within the module
- prevent namespace pollution as variable or method name is kept wihtin the module so that they can be used somewhere else
