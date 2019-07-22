//选择排序

package selectsort

func SelectSort(values []int) {
	if len(values) <= 1 {
		return
	}
	for i := 0; i < len(values); i++ {
		minIndex := 0
		for j := i + 1; j < len(values); j++ {
			if values[j] < values[minIndex] {
				minIndex = j
			}
		}
		values[i], values[minIndex] = values[minIndex], values[i]
	}
}
