
# Liskov

A program that uses an interface must not be confused by an implementation of that interface. <br />

Derived classes must be substitutable for their base classes.

What violate the Liskov principle?
* When a class does not have a proper implementation for all the methods of its parent
class (or its interface for that matter) <br />
Another violation is **Leaky abstraction**, I'll explain this in the following sentence.
* Different substitutes return things at different types
* A derived class is less permissive with regard to method arguments
* Secretly programming against a more specific type

What is **Leaky Abstraction**?

Imagine you have an interface called File
```go
type File interface{
	Rename(fileName string) error
	ChangeOwner(user, group string) error
}
```
Then you want to implement a local file concrete

```go
type LocalFile struct {}

func (l *LocalFile) Rename(fileName string) error {
	// Implementation of renaming a local file
    return nil
}

func (l *LocalFile) ChangeOwner(user, group string) error {
    // Implementation of changing owner of a file
    return nil
}
```
Now, we want to add Dropbox

```go
type Dropbox struct {}

func (d *Dropbox) Rename(fileName string) error {
	// Implementation of renaming a dropbox file
    return nil
}

func (d *Dropbox) ChangeOwner(user, group string) error {
    panic("can not change owner of a file")
}
```

As you can see Dropbox is not a good substitute for File interface because I does not implement the ```ChangeOwner``` method. <br />
So the File interface turned out to be an improper generalization of the "file" concept. such improper generalization is usually called a **Leaky Abstraction**


1. Dissecting the principle, we recognize two conceptual parts. First it’s about derived
   classes and base classes. Then it’s about being substitutable. 
   Bringing the two concepts together, the Liskov Substitution principle says that if we
   create a class that extends another class or implements an interface, it has to behave as
   expected.
2. If you have car your know that you have to change its motor oil at specific KMs, if your consider that every car that wants to change its motor oil has to go over lubritorium then this role does not work for Motorcycle.
3. 
4. Pointed out in the Leaky abstraction example
5. Seperated interfaces 