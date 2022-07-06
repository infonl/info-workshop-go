package bubbleSort

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

func Swap[T Number](numbers []T, i int, j int) {
	temp := numbers[i]
	numbers[i] = numbers[j]
	numbers[j] = temp
}

func BubbleSort[T Number](numbers []T) {
	for i := 0; i < len(numbers)-1; i++ {
		for j := 0; j < len(numbers)-1; j++ {
			if numbers[j] > numbers[j+1] {
				Swap(numbers, j, j+1)
			}
		}
	}
}
