# Prototype design patter

* Useful when creating objects
* Can use an already created instance of some type to clone it and complete it
* Aims to have an object or set of objects that is already created at compilation time


- useful for example as a default template for a user who is registered on your webpage
- useful for a default pricing plan in some service

Difference between builder pattern: objects are cloned for the user instead of building them at runtime.
Can also build a cache like solution storing information using a prototype.

* Main objective for the prototype design pattern is to avoid repetitive object creation.
Imagine a default object composed of dozens of fields and embedded types.

* We don't with to write everything needed by this type every time that we use the object
especially if we can mess it up by creating instances with different foundations.
So we need to maintain a set of objects that will be cloned to create new instances & provide a default value.

* Free CPU of complex object initialization to take more memory resources.


