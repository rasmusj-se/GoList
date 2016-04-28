package GoList

import "errors"

type Insort struct{
    Less func(a interface{}, b interface{}) bool
    items []interface{}
}

func (insort *Insort) Sort() ([]interface{}, error) {
    if insort.Less == nil{
        return nil, errors.New("Less function is not implemented")
    }
    for i := 0; i < len(insort.items)-1; i++{
        for i := 0; i < len(insort.items)-1; i++{
            if !insort.Less(insort.items[i], insort.items[i+1]){
                temp := insort.items[i+1]
                insort.items[i+1] = insort.items[i]
                insort.items[i] = temp
            }
        }
    }
    return insort.items, nil
}

func (insort *Insort) Put(item interface{}){
    insort.items = append(insort.items, item)
}