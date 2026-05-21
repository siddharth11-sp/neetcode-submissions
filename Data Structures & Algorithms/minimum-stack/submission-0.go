type MinStack struct {
	mainStack []int
	minStack  []int
}

func Constructor() MinStack {
	return MinStack{
		mainStack: []int{},
		minStack:  []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.mainStack = append(this.mainStack, val)

	if len(this.minStack) == 0 || val <= this.minStack[len(this.minStack)-1] {
		this.minStack = append(this.minStack, val)
	}
}

func (this *MinStack) Pop() {

	top := this.mainStack[len(this.mainStack)-1]

	this.mainStack = this.mainStack[:len(this.mainStack)-1]

	if top == this.minStack[len(this.minStack)-1] {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}

}

func (this *MinStack) Top() int {
	return this.mainStack[len(this.mainStack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}
