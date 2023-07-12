# Observer pattern
The observer pattern is a software design pattern that defines a one-to-many dependency between objects. When one object (the subject or observable) changes its state, it notifies all the dependent objects (the observers) about the change, usually by calling a method on them. This way, the observers can update themselves automatically without being tightly coupled to the subject.

The observer pattern is often used for implementing event-driven systems, where the subject is a source of events and the observers are handlers of those events. 

In this example, **weather station** measures some parameters of the weather and then notifies some **farmers** and **roof worker** to start or stop working.