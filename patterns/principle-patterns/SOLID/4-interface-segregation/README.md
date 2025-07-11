# Interface Segregation

Make fine-grained interfaces that are client specific. <br />
“Fine-grained interfaces” stands for interfaces with a small amount of methods.
“Client specific” means that interfaces should define methods that make sense from the
point of view of the client that uses the interface.

What violate the Interface Segregation?
1. Multiple usecases (Just like service containers)
2. No interface, Just a class