# Single Responsibility

Robert C. Martin (Uncle Bob) defines the Single Responsibility Principle as:
**Each software module should have one and only one reason to change.**

What violate the single responsibility?

* The class has many instance variable
* The class has many public methods
* Each method of the class uses different instance variables
* Specific tasks are delegated to private method

1. The Responsibility term is so confusing, what we define responsibility? What is exactly responsibility? maybe every developer has its own definition of responsibility. I think the better term to use for defining responsibility is cohesion. Responsibility are the reasons to change

2. Imagine you build a university and at the first you hire someone who does lots of things, but as more students gets into the university you want your employees work efficiently, so you may give them only one responsibility.

3. https://github.com/ory/dockertest/blob/207b20aded3b6876a3acb2a6cdc4447eb8f49bfc/docker/event.go#L121

4. Examples:
    - [Golang](example.go)

5. Collaborator classes, if a package or class uses collaborator classes and does not do many works it means it follows single Responsibility pattern
