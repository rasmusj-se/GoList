package GoList

import (
    "errors"
    "fmt"
)

//HashMap creates a HashMap datastructure. This uses the internal Go map since it is already a HashMap
type HashMap struct{
    items []interface{}
    keys []string
    count int
    size int
    blocks uint
}

func hash(key string, size int) int {
    v := 0
    for _,r := range key{
        v += int(r)
    }
    return v % size;
}

func step(pos, size int) int {
    return (((pos/size) % (size/2))*2)+1
}

func (hashmap *HashMap) Initialize() error {
    hashmap.blocks = 1
    hashmap.items = make([]interface{}, 1<<hashmap.blocks)
    hashmap.keys = make([]string, 1<<hashmap.blocks)
    hashmap.size = 1<<hashmap.blocks
    return nil
}

//Insert puts an element into the HashMap. O(1) operation.
func (hashmap *HashMap) Insert(key string, value interface{}) error {
    
    //Grow hashmap if needed
    if hashmap.Count() == hashmap.size {
        items := make([]interface{}, 1<<(hashmap.blocks+1))
        keys := make([]string, 1<<(hashmap.blocks+1))
        for i := 0; i < hashmap.size; i++{
            items[i] = hashmap.items[i]
            keys[i] = hashmap.keys[i] 
        }
        hashmap.items = items
        hashmap.keys = keys
        hashmap.blocks++
        hashmap.size = 1<<hashmap.blocks
    }
    
    location := hash(key, hashmap.size)
    for i := location; i < location + hashmap.size; i += step(i, hashmap.size) {
        if hashmap.keys[i % hashmap.size] == "" {
            hashmap.keys[i % hashmap.size] = key
            hashmap.items[i % hashmap.size] = value
            hashmap.count++;
            return nil
        }
    }
    return errors.New("Map is full")
}

//Remove removes an item from the HashMap. O(1) operation.
func (hashmap *HashMap) Remove(key string) error {
    
    //Shrink hashmap if needed  
    if hashmap.Count() < 1<<(hashmap.blocks-1) {
        hashmap.blocks--
        hashmap.size = 1<<hashmap.blocks
        items := hashmap.items
        keys := hashmap.keys
        hashmap.items = make([]interface{}, 1<<hashmap.blocks)
        hashmap.keys = make([]string, 1<<hashmap.blocks)
        copy(hashmap.items[0:hashmap.size], items[:])
        copy(hashmap.keys[0:hashmap.size], keys[:])
    }
    
    location := hash(key, hashmap.size)
    for i := location; i < location + hashmap.size; i += step(i, hashmap.size) {
        if hashmap.keys[i % hashmap.size] == key {
            hashmap.keys[i % hashmap.size] = ""
            hashmap.items[i % hashmap.size] = nil
            hashmap.count--;
        }
    }
    return nil
}

//Get retreives an item from the HashMap
func (hashmap *HashMap) Get(key string) (interface{}, error) {
    location := hash(key, hashmap.size)
    for i := location; i < location + hashmap.size; i += step(i, hashmap.size) {
        if hashmap.keys[i % hashmap.size] == key{
            return hashmap.items[i % hashmap.size], nil
        }
    }
    return nil, errors.New("The specified key was not found in the map")
}

//Count returns the number of values in the hashmap
func (hashmap *HashMap) Count() int {
    return hashmap.count
}

//Count returns the number of values in the hashmap
func (hashmap *HashMap) Print() {
    fmt.Printf("HashMap. Size %v. Count %v.", hashmap.size, hashmap.count)
    fmt.Println("")
    for k,v := range hashmap.keys{
        fmt.Printf("%v %v = %v", k, v, hashmap.items[k])
        fmt.Println("")
    }
    fmt.Println("")
}