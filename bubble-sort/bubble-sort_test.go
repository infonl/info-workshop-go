package bubbleSort

import "testing"

var numbers = []int{1, 2, 3}

func TestSwap(t *testing.T) {
	Swap(numbers, 0, 1)
	if numbers[0] != 2 && numbers[1] != 1 {
		t.Error("Swap did not work")
	}
}

func TestBubbleSort(t *testing.T) {
	BubbleSort(numbers)
	if !(numbers[2] > numbers[1] && numbers[1] > numbers[0]) {
		t.Error("Sort did not work")
	}
}
