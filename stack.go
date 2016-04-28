package GoList

//Stack is a basic stack implementetation in go using a linked list
type Stack struct{
    count int
    top *node
}

type node struct{
    Child *node
    Value interface{}
}

//Push pushes a new item to the stack
func (stack *Stack) Push(value interface{}){
    stack.top = &node{
        Value: value,
        Child: stack.top,
    }
    stack.count++
}

//Pop returns and removes the top item of the stack
func (stack *Stack) Pop() interface{} {
    if stack.count == 0{
        return nil
    }
    value := stack.top.Value
    stack.top = stack.top.Child
    stack.count--
    return value
}
//Peek returns the top item of the stack but does not remove it
func (stack *Stack) Peek() interface{}{
    if stack.count == 0{
        return nil
    }
    return stack.top.Value
}

//Count returnes the number of items in the stack
func (stack *Stack) Count() interface{}{
    return stack.count
}