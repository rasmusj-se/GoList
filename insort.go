package GoList

import "errors"

type Insort struct{
    Less func(a interface{}, b interface{}) bool
}

func (insort *Insort) Sort(items Sortable) ([]interface{}, error) {
    if insort.Less == nil{
        return nil, errors.New("Less function is not implemented")
    }
    for i := 0; i < len(items)-1; i++{
        for i := 0; i < len(items)-1; i++{
            if !insort.Less(items[i], items[i+1]){
                temp := items[i+1]
                items[i+1] = items[i]
                items[i] = temp
            }
        }
    }
    return items, nil
}