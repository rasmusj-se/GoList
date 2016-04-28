# GoList is a datastructure implementation in Go

The purpose of GoList is to create a library packed with all of the datastructures and algoritms that may come of use when using Go.
This includes general datatypes such as Stack, Queue, Heap, HashSet etc and algorithms such as QuickSort, Insort, Mergesort etc.

For every datatype there is a set of functions that can be executed. For example: Insert, Get, Put, Push, Pop etc.
The most useful type is the dynamic List that has a set of features to it.

A more detailed documentation is available at https://godoc.org/github.com/rasmusj-se/GoList

## HashMap

The HashMap automatically grows and shrinks as data is inserted or removed.
It grows in the order of 2^m. (Note: For big data this could mean a large HashMap)

Initialization:

```
hashmap := &HashMap{}
hashmap.Initialize()
```

Insert: `.Insert(key string, value interface{})` 

Best: O(1). Worst: O(n)

Get: `.Get(key string) (interface{})` 

Best: O(1). Worst: O(n)

Insertion: `hashmap.Insert(key string, value interface{})` 

Best: O(1). Worst: O(n)

*Note: Go already contains a standard implementation of a hashmap known as only map. Use that over GoList!*