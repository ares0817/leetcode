package problem_155

type node struct {
	value int
	min int
}

type MinStack struct {
	list []*node
}


/** initialize your data structure here. */
func Constructor() MinStack {
	stack := MinStack{
		list: make([]*node, 0),
	}
	return stack
}


func (this *MinStack) Push(x int)  {
	node := &node{
		value: x,
		min: x,
	}
	if len(this.list) != 0 && x > this.GetMin() {
		node.min = this.GetMin()
	}
	this.list = append(this.list, node)
}


func (this *MinStack) Pop()  {
	this.list = this.list[:len(this.list) - 1]
}


func (this *MinStack) Top() int {
	if len(this.list) == 0 {
		return 0
	}
	return this.list[len(this.list) - 1].value
}


func (this *MinStack) GetMin() int {
	if len(this.list) == 0 {
		return 0
	}
	return this.list[len(this.list) - 1].min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
