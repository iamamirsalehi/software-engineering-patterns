# Chatty I/O

## 1.Understood the aim of the pattern

Chatty I/O happens when an application performs too many small, sequential I/O operations, such as:

- Hitting the database inside a loop
- Making separate HTTP calls for each item in a list
- Reading multiple files one by one in a blocking way

Instead of optimizing access to external systems, the code ends up chatting (talking a lot, doing very little per request).

## 2.Understood the real example of pattern in the day-to-day life

Imagine you have something to tell your friend but you tell him/her just one word everytime you call him/her

## 3.Found some examples in the product code or open-source projects

- [An example in github PR](https://github.com/skkuding/codedang/pull/2682)

## 4.Implemented an example

Implemented examples are:
- [Golang](/patterns/anti-patterns/chatty-io/example.go)
- [PHP](/patterns/anti-patterns/chatty-io/example.php)
- [Python](/patterns/anti-patterns/chatty-io/example.go) 

## 5.Sign of the pattern

- Slow batch operations (e.g. syncing thousands of records)
- Increased database or API response time
- Tight loops with calls to external systems
- Profiling shows lots of I/O wait time
- High number of open connections (especially in MySQL, Redis, or MongoDB)