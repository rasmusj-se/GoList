# GoList is a datastructure implementation in Go

A more detailed documentation is available at https://godoc.org/github.com/rasmusj-se/GoList

## Stack

**Initialization**

`stack := &Stack{}`

**Push**

`stack.Push("Value")`

Best: **O**(1), Worst: **O**(1)

**Pop**

`value := stack.Pop()`

Best: **O**(1), Worst: **O**(1)

**Peek**

`value := stack.Peek()`

Best: **O**(1), Worst: **O**(1)

**Count**

`stack.Count()`

Best: **O**(1), Worst: **O**(1)

## Queue

**Initialization**

`queue := &Queue{}`

**Enqueue**

`queue.Enqueue("Value")`

Best: **O**(n), Worst: **O**(n)

**Dequeue**

`value := queue.Dequeue()`

Best: **O**(1), Worst: **O**(1)

**Peek**

`value := queue.Peek()`

Best: **O**(1), Worst: **O**(1)

**Count**

`queue.Count()`

Best: **O**(1), Worst: **O**(1)