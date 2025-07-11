# Open-Closed

A Module should be open for extension but closed for modification.

What violate the Open-Closed principle?
* It contains conditions to determine a strategy
* Conditions using the same variables or constants are recurring inside the
  class or related classes.
* The class contains hard-coded references to other classes or class names.
* Inside the class, objects are being created using the new operator.
* The class has protected properties or methods, to allow changing its behavior by overriding state or behavior.

1. You should be able to extend a class behaviour without modifying it. **A unit of code** can be considered **open for extension**. The fact the **no actual modification** is needed to change **the behaviour of a unit of code** makes it **closed** for modification

2. When you use your laptop and plug in your usd flash or your external hard drive, you don't have to make changes in your laptop to support these two kind of devices, you just use the USD port

3. https://github.com/iamamirsalehi/project-with-milad/blob/main/app/Application/Service/PaymentService/PaymentService.php

4. 

5. Extending a class behaviour by just passing to its constructor without changing the class itself, **Abstract factory** can also help to produce families of related objects without specifying their concrete 