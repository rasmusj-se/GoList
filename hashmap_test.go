package GoList

import (
    "testing"
    "math/rand"
    "time"
)

func TestHashMap(t *testing.T) {
        
       rand.Seed(time.Now().UnixNano())
    
       hashmap := &HashMap{}
       hashmap.Initialize()
       hashmap.Insert("test", "one")
       if val, err := hashmap.Get("test"); err != nil || val != "one"{
           t.Error("Map did not return correct value")
       }
       
       hashmap.Insert("testing", "two")
       if val, err := hashmap.Get("testing"); err != nil || val != "two"{
           t.Error("Map did not return correct value")
       }
      
       
       for i := 0; i < 5; i++ {
        if err :=hashmap.Insert(RandString(5), "value"); err != nil {
           t.Error(err)
        }
       }
       
       
       hashmap.Remove("test")
       
       if hashmap.Count() != 6 {
            t.Errorf("Map removal faliure. Count does not match. %v", hashmap.Count())
       }
       
       hashmap = &HashMap{}
       hashmap.Initialize()
       
        if hashmap.Count() != 0 {
            t.Errorf("Map reinitialization failure")
       }
       
       //Large insertion
        for i := 0; i < 32; i++ {
        if err := hashmap.Insert(RandString(31), "value"); err != nil {
           t.Error(err)
        }
       }
       hashmap.Print()

}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandString(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}