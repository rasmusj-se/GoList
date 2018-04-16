# GoList is a datastructure implementation in Go

[![GitHub issues](https://img.shields.io/github/issues/rasmusj-se/GoList.svg)](https://github.com/rasmusj-se/GoList/issues)
[![Build Status](https://travis-ci.org/rasmusj-se/GoList.svg?branch=master)](https://travis-ci.org/rasmusj-se/GoList)
[![Go Report Card](https://goreportcard.com/badge/github.com/rasmusj-se/GoList)](https://goreportcard.com/report/github.com/rasmusj-se/GoList)
[![GoDoc](https://godoc.org/github.com/rasmusj-se/GoList?status.svg)](https://godoc.org/github.com/rasmusj-se/GoList)
[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/hyperium/hyper/master/LICENSE)

`This project is not maintained. It was created as a way to learn how to adopt datastructures and algorithms in Go.`

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

Best: **O**(1), Worst: **O**(1)

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

You will need to define your own comparer for your data types. Here is an example for an int list.

```
sorter.Less = func (a interface{}, b interface{}) bool {
    return a.(int) < b.(int) //For example
}
```

**Sort**

You must use the list type `Sortable` or `[]interface{}`.

*I'm looking for other solutions to this to allow sorting of already defined lists*

```
list := GoList.Sortable{9,8,7,6,5,4,3,2,1,0} //For example
list, err := sorter.Sort(list)
```

Best: **O**(n^2), Worst: **O**(n^2)
