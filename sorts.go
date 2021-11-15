package ds

import (
	dsheap "cf/heap"
	"cf/interfaces"
	"container/heap"
	"math"
)

func BubbleSort(sortable []interfaces.Comparable) []interfaces.Comparable {

	for i, size := 0, len(sortable); i < size-1; i++ {
		for j := 0; j < size-i-1; j++ {
			if sortable[j].Compare(sortable[j+1]) > 0 {
				sortable[j+1], sortable[j] = sortable[j], sortable[j+1]
			}
		}
	}

	return sortable
}

func SelectionSort(sortable []interfaces.Comparable) []interfaces.Comparable {

	for i, size := 0, len(sortable); i < size-2; i++ {
		min := i
		for j := i + 1; j < size-1; j++ {
			if sortable[j].Compare(sortable[min]) > 0 {
				min = j
			}
		}
		sortable[min], sortable[i] = sortable[i], sortable[min]
	}

	return sortable
}

func InsertionSort(sortable []interfaces.Comparable) []interfaces.Comparable {

	for i, size := 0, len(sortable); i < size-1; i++ {

		for j := i - 1; j >= 0 && sortable[j].Compare(sortable[j+1]) > 0; j-- {
			sortable[j], sortable[j+1] = sortable[j+1], sortable[i]
		}
	}

	return sortable
}

func MergeSort(sortable []interfaces.Comparable) []interfaces.Comparable {

	count := len(sortable)
	switch {
	case count > 2:

		lb, rb := MergeSort(sortable[:count/2]), MergeSort(sortable[count/2:])
		sortable = append(lb, rb...)
		lastIndex := len(sortable) - 1

		for i := 0; i < lastIndex; i++ {
			move_index := i
			for j := i + 1; j < lastIndex+1; j++ {
				if sortable[move_index].Compare(sortable[j]) > 0 {
					move_index = j
				}
			}

			if move_index != i {
				sortable[i], sortable[move_index] = sortable[move_index], sortable[i]
			}
		}

	case len(sortable) > 1 && sortable[0].Compare(sortable[1]) > 0:
		sortable[0], sortable[1] = sortable[1], sortable[0]
	}

	return sortable
}

func MergeSortStack(sortable []interfaces.Comparable) []interfaces.Comparable {

	//stack := make([]int, len(sortable)/2)

	count := len(sortable)
	switch {
	case count > 2:

		left, right := MergeSort(sortable[:count/2]), MergeSort(sortable[count/2:])
		sortable = append(left, right...)
		lastIndex := len(sortable) - 1

		for i := 0; i < lastIndex; i++ {
			move_index := i
			for j := i + 1; j < lastIndex+1; j++ {
				if sortable[move_index].Compare(sortable[j]) > 0 {
					move_index = j
				}
			}

			if move_index != i {
				sortable[i], sortable[move_index] = sortable[move_index], sortable[i]
			}
		}

	case len(sortable) > 1 && sortable[0].Compare(sortable[1]) > 0:
		sortable[0], sortable[1] = sortable[1], sortable[0]
	}

	return sortable
}

func hoarpartion(arr []interfaces.Comparable, low, high int) ([]interfaces.Comparable, int) {

	i, j, pivot := low, high, arr[low+high/2]
	for {
		for arr[i].Compare(pivot) < 0 {
			i++
		}
		for arr[j].Compare(pivot) > 0 {
			j--
		}
		if i >= j {
			return arr, j
		}
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

func QuickSortHoare(arr []interfaces.Comparable, low, high int) []interfaces.Comparable {
	if low < high {
		var p int
		arr, p = hoarpartion(arr, low, high)
		arr = QuickSortHoare(arr, low, p-1)
		arr = QuickSortHoare(arr, p+1, high)
	}
	return arr
}

func lomutoPartitionWithMedian(arr []interfaces.Comparable, low, high int) ([]interfaces.Comparable, int) {

	mid := (low + high) / 2

	if arr[mid].Compare(arr[low]) < 0 {
		arr[low], arr[mid] = arr[mid], arr[low]
	}

	if arr[high].Compare(arr[low]) < 0 {
		arr[high], arr[low] = arr[low], arr[high]
	}

	if arr[mid].Compare(arr[high]) < 0 {
		arr[mid], arr[high] = arr[high], arr[mid]
	}

	pivot, i := arr[high], low
	for j := low; j < high; j++ {
		if arr[j].Compare(pivot) < 0 {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func QuickSortLomuto(arr []interfaces.Comparable, low, high int) []interfaces.Comparable {
	if low < high {
		var p int
		arr, p = lomutoPartitionWithMedian(arr, low, high)
		arr = QuickSortLomuto(arr, low, p-1)
		arr = QuickSortLomuto(arr, p+1, high)
	}
	return arr
}

func ShellSort(arr []interfaces.Comparable) []interfaces.Comparable {
	for size, s := len(arr), len(arr); s > 0; s /= 2 {
		for i := 0; i < size; i++ {
			for j := i + s; j < size; j += s {
				if arr[i].Compare(arr[j]) > 0 {
					arr[i], arr[j] = arr[j], arr[i]
				}
			}
		}
	}

	return arr
}

func CountSort(arr []uint16) []uint16 {
	n, counter := len(arr), make([]uint64, math.MaxUint16)
	for i := 0; i < n; i++ {
		counter[arr[i]]++
	}
	for arrIndex, i := 0, uint16(0); i < math.MaxUint16; i++ {
		for j := counter[i]; j > 0; j-- {
			arr[arrIndex] = i
			arrIndex++
		}

	}

	return arr
}

func HeapSort(arr []interfaces.Comparable) []interfaces.Comparable {

	pq := make(dsheap.Heap, len(arr))

	for index, value := range arr {
		pq[index] = value
	}

	heap.Init(&pq)

	for item, i := heap.Pop(&pq), 0; item != nil; item = heap.Pop(&pq) {
		arr[i] = item.(interfaces.Comparable)
		i++
	}

	return arr
}
