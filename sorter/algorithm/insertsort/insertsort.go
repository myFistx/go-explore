// 插入排序
// 从第二个数开始向右侧遍历，每次均把该位置的元素移动至左侧，放在一个正确的位置

package insertsort

func InsertSort(values []int) {
	if len(values) <= 1 {
		return
	}
	for i := 1; i < len(values); i++ {
		if values[i] < values[i-1] {
			j := i - 1
			temp := values[i]
			for ; j >= 0; j-- {
				if values[j] > temp {
					values[j+1] = values[j]
				} else {
					break
				}
			}
			values[j+1] = temp
		}
	}
}
