type entry struct {
	val int
	min int
}

type MinStack struct {
	items []entry
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	currentMin := val

	if len(this.items) > 0 {
		currentMin = min(val, this.items[len(this.items)-1].min)
	}

	this.items = append(this.items, entry{val: val, min: currentMin})
}

func (this *MinStack) Pop() {
	this.items = this.items[:len(this.items)-1]
}

func (this *MinStack) Top() int {
	return this.items[len(this.items)-1].val
}

func (this *MinStack) GetMin() int {
	return this.items[len(this.items)-1].min
}