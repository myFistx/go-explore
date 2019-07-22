// 冒泡排序
// 冒泡排序只会操作相邻的两个数据。
// 每次冒泡操作都会对相邻的两个元素进行比较，看是否满足大小关系要求。
// 如果不满足就让它俩互换。一次冒泡会让至少一个元素移动到它应该在的位置，重复n 次，就完成了 n 个数据的排序工作。

package bubblesort

func BubbleSort(values []int) {
	if len(values) <= 1 {
		return
	}
	flag := true
	for i := 0; i < len(values)-1; i++ {
		flag = true
		for j := 0; j < len(values)-i-1; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				flag = false
			}
		}
		if flag == true {
			break
		}
	}
}
