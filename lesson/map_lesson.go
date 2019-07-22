package main

import (
	"fmt"
	"sort"
)

func main() {
	//map声明方式1
	var mapList map[int]string
	mapList = map[int]string{0: "hello", 1: "world"}
	fmt.Println(mapList)

	//map声明方式2
	map2 := make(map[int]string)
	map2[3] = "hello"
	map2[4] = "golang"
	fmt.Println(map2)

	//判断map中的key是否存在
	value1, isPresent1 := map2[3]
	fmt.Println(value1)
	fmt.Println(isPresent1)
	value2, isPresent2 := map2[5]
	fmt.Println(value2)
	fmt.Println(isPresent2)

	capitals := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}
	for key := range capitals {
		fmt.Println("Map item: Capital of", key, "is", capitals[key])
	}

	//map与切片结合
	items := make([]map[int]int, 5)
	fmt.Println(items)

	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][1] = 2
	}
	fmt.Println(items)

	s1 := make([]map[int]string, 0)
	map11 := map[int]string{0: "hello"}
	s1 = append(s1, map11)
	fmt.Println(s1)

	//map排序
	maps := map[string]string{
		"a": "a1",
		"f": "f1",
		"d": "d1",
		"e": "e1",
	}
	fmt.Println(maps)
	keys := make([]string, 0)
	i := 0
	for k, _ := range maps {
		keys = append(keys, k)
		i++
	}
	sort.Strings(keys)
	fmt.Println(keys)
	for _, k := range keys {
		fmt.Println("Key: %v, Value: %v / ", k, maps[k])
	}

	//键值对调
	invMap := make(map[string]string, len(maps))
	for k, v := range maps {
		invMap[v] = k
	}
	fmt.Println(invMap)
}
