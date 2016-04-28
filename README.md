# GoList is a datastructure implementation in Go

The purpose of GoList is to create a library packed with all of the datastructures and algoritms that may come of use when using Go.
This includes general datatypes such as Stack, Queue, Heap, HashSet etc and algorithms such as QuickSort, Insort, Mergesort etc.

For every datatype there is a set of functions that can be executed. For example: Insert, Get, Put, Push, Pop etc.
The most useful type is the dynamic List that has a set of features to it.

## HashMap

Go already contains a HashMap, you should use the already existing one over mine for your own project.
However, GoList contains an implementation to a self implemented HashMap.

Initialization:

```
hashmap := &HashMap{}
hashmap.Initialize()
```

Insertion: `hashmap.Insert(key string, value interface{})` Usually an O(1) operation, depends on potential key collisions.