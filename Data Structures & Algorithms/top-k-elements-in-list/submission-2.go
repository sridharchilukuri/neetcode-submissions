type Item struct {
	number    int
	frequency int
}

type MinHeap struct {
	Heap []Item
}

func NewMinHeap(k int) *MinHeap {
	// Index 0 is unused, so the root starts at index 1.
	return &MinHeap{
		Heap: make([]Item, 1, k+1),
	}
}

func (mh *MinHeap) Length() int {
	return len(mh.Heap) - 1
}

func (mh *MinHeap) push(value Item) {
	mh.Heap = append(mh.Heap, value)
	i := mh.Length()

	for i > 1 {
		parent := i / 2

		if mh.Heap[parent].frequency <= mh.Heap[i].frequency {
			break
		}

		mh.Heap[i], mh.Heap[parent] = mh.Heap[parent], mh.Heap[i]
		i = parent
	}
}

func (mh *MinHeap) pop() Item {
	root := mh.Heap[1]
	last := mh.Length()

	mh.Heap[1] = mh.Heap[last]
	mh.Heap = mh.Heap[:last]

	i := 1

	for {
		left, right := 2*i, 2*i+1
		smallest := i

		if left <= mh.Length() &&
			mh.Heap[left].frequency < mh.Heap[smallest].frequency {
			smallest = left
		}

		if right <= mh.Length() &&
			mh.Heap[right].frequency < mh.Heap[smallest].frequency {
			smallest = right
		}

		if smallest == i {
			break
		}

		mh.Heap[i], mh.Heap[smallest] = mh.Heap[smallest], mh.Heap[i]
		i = smallest
	}

	return root
}

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)

	for _, number := range nums {
		freq[number]++
	}

	minHeap := NewMinHeap(k)

	for value, count := range freq {
		// Once we have k elements, only replace the root if the
		// incoming frequency is greater than the smallest frequency.
		if minHeap.Length() >= k {
			if count <= minHeap.Heap[1].frequency {
				continue
			}

			minHeap.pop()
		}

		minHeap.push(Item{
			number:    value,
			frequency: count,
		})
	}

	// All remaining heap items belong to the top k.
	// We don't need to pop because the output order doesn't matter.
	result := make([]int, 0, k)

	for i := 1; i <= minHeap.Length(); i++ {
		result = append(result, minHeap.Heap[i].number)
	}

	return result
}