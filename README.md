# GoList is a datastructure implementation in Go

A more detailed documentation is available at https://godoc.org/github.com/rasmusj-se/GoList

## Stack

**Initialization**

`stack := &GoList.Stack{}`

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

`queue := &GoList.Queue{}`

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

## Insort

**Initialization**

` sorter := &GoList.Insort{}`

**Comparer**

You will need to define your own comparer for your data types.
Here is an example for an int list.

```
sorter.Less = func (a interface{}, b interface{}) bool {
    return a.(int) < b.(int)
}
```

**Put**

Due to the design of Go you will have to commit yor list to the sorter to have them sorted

*I'm looking for a better solution to this as it adds O(n) to the execution time*

`sorter.Put(value)`

Best: **O**(n), Worst: **O**(n)

**Sort**

`sorted, err := sorter.Sort()`

Best: **O**(n^2), Worst: **O**(n^2)
