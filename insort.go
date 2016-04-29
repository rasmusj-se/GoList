package GoList

import "errors"

//Insort creates an insort alogritm. It needs to be given a comparer function in the property Less.
type Insort struct {
	Less func(a interface{}, b interface{}) bool
}

//Sort sorts the given sortable with the Insort alogritm
func (insort *Insort) Sort(items Sortable) ([]interface{}, error) {
	if insort.Less == nil {
		return nil, errors.New("Less function is not implemented")
	}
	for i := 0; i < len(items)-1; i++ {
		for i := 0; i < len(items)-1; i++ {
			if !insort.Less(items[i], items[i+1]) {
				temp := items[i+1]
				items[i+1] = items[i]
				items[i] = temp
			}
		}
	}
	return items, nil
}
