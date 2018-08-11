# Adapter

In go an adapter will allow us to use something that wasn't built for a specific task at the beginning.

The adapter pattern is useful when for example an interface gets outdated and it's not possible to replace it easily or fast.

Instead you create a new interface to deal with the current needs of your application which under the hood uses the implementation
of the old interface. Adapter also helps us to maintain the open or closed principle.

It helps you to fit the needs of 2 parts of the code that incompatible at first.

2 Interfaces that are incompatible but which must work together are good candidates for an adapter.
They could also use the facade pattern for example.


For this example we will have an old printer interface and a new one.
We need an adapter so that the user can still use an old implementation if necessary for the requirements.



* Create an adapter object that implements the ModernPrinter interface

* The new adapter object must contain an instance of the LegacyPrinter interface

* When using ModernPrinter, it must call the LegacyPrinter interface under the hood, prefixing it with the text Adapter.