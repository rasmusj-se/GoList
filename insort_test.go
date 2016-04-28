package GoList

import "testing"

func TestInsort(t *testing.T){
    sorter := &Insort{}
    //Comparer
    sorter.Less = func (a interface{}, b interface{}) bool {
        return a.(int) < b.(int)
    }
    //List to sort
    list := [10]int{9,8,7,6,5,4,3,2,1,0}
    //Commit the items to sort into the sorter
    for _, v := range list{
        sorter.Put(v)
    }
    sorted, _ := sorter.Sort()
    for i, v := range sorted{
        if i != v{
            t.Error("Sorting was not correct")
        }
    }
}