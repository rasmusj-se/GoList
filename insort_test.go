package GoList

import "testing"

func TestInsort(t *testing.T){
    sorter := &Insort{}
    //Comparer
    sorter.Less = func (a interface{}, b interface{}) bool {
        return a.(int) < b.(int)
    }
    //List to sort
    list := Sortable{9,8,7,6,5,4,3,2,1,0}
    //Commit the items to sort into the sorter
    list , _ = sorter.Sort(list)
    //Check if list was sorted
    for i, v := range list{
        if i != v{
            t.Error("Sorting was not correct")
        }
    }
}