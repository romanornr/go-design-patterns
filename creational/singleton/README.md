The singleton pattern will give you the power to havbe a unique instance of some struct in your application

and that no other package can create any clone of the struct.

With a singleton, you're also hiding the complexity of creating the object in case it requires some computation

and the pitfall of creating it every time you need an instance of it.

All this code writing, checking if the variable already exists and storage are encapsulated in the singleton and you don't
have to repeat it everywhere if you use a global varibale.

Note: This implementation is not thread safe.