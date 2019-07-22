/*
主程序
1.获取命令行输入并解析
2.从对应文件中读取数据
3.调用对应的排序算法
4.将排序结果写入文件中
5.计算排序算法的时间复杂度
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"go-explore/sorter/algorithm/bubblesort"
	"go-explore/sorter/algorithm/insertsort"
	"go-explore/sorter/algorithm/qsort"
	"go-explore/sorter/algorithm/selectsort"
	"io"
	"math/rand"
	"os"
	"strconv"
	"time"
)

//声明指针变量
//命令行输入的文件名，默认为“infile”
var infile *string = flag.String("i", "infile", "File contains values of sorting")

//命令行输出的文件名，默认为“outfile”
var outfile *string = flag.String("o", "outfile", "File to receive values")

//命令行输入要排序等算法类型，默认“qsort”
var algorithm *string = flag.String("a", "qSort", "sort algorithm")

//生成的文件里面有多少个元素，默认10个
var num *int = flag.Int("n", 10, "values number")

//数据的排序类型，desc 倒序，asc 正序，diSort 无序
var sortType *string = flag.String("t", "diSort", "sorting type")

//读取文件内容的方法
func readValues(infile string) (values []int, err error) {
	//打开输入的文件名内容，不存在则打印错误信息
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Failed to open the input file ", infile)
		return
	}
	//关闭文件，关键字defer，用于释放资源，在函数返回之前被执行
	defer file.Close()

	//创建一个文件操作句柄
	br := bufio.NewReader(file)

	//make函数创建一个数组，并返回引用该数组的切片。
	values = make([]int, 0)
	//无限循环，读取文件中的每一行
	for {
		//读取文件中的一行
		line, isPerFix, lineErr := br.ReadLine()
		//EOF是文件结束符（End Of File）的缩写
		if lineErr != nil {
			if lineErr != io.EOF {
				err = lineErr
			}
			break
		}

		//当文件中某一行长度过长的时候，isPerFix就为true
		if isPerFix {
			fmt.Println("a too long line! ", infile)
			return
		}
		//转成字符串
		str := string(line)
		//字符串转成数字
		value, err2 := strconv.Atoi(str)
		if err2 != nil {
			err = err2
			return
		}
		//往数组中添加元素
		values = append(values, value)
	}
	return
}

//将排序后的结果写入到文件中
func writeValues(values []int, outfile string, fileType string) {
	var newOutFile = fileType + "_" + outfile
	//创建一个输出文件
	file, err := os.Create(newOutFile)
	if err != nil {
		fmt.Println("Failed to create the output file ", newOutFile)
		return
	}
	//释放资源，函数退出时执行
	defer file.Close()

	for _, v := range values {
		str := strconv.Itoa(v)
		file.WriteString(str + "\n")
	}
	return
}

/**
 * 生成未排序的文件
 */
func generateFile(count int, sortType string, infile string) (inFileName string, errs error) {
	var newInFile = sortType + "_" + infile
	_, err := os.Stat(newInFile)
	if err == nil {
		fmt.Println("infile is exist!")
		return newInFile, nil
	} else {
		switch sortType {
		case "desc":
			values := descType(count)
			writeValues(values, infile, sortType)
		case "asc":
			values := ascType(count)
			writeValues(values, infile, sortType)
		default:
			values := diSortType(1, 1000, count)
			writeValues(values, infile, sortType)
		}
	}
	return newInFile, nil
}

//正序
func ascType(num int) (values []int) {
	values = make([]int, 0)
	for i := 1; i <= num; i++ {
		values = append(values, i)
	}
	return values
}

//倒序
func descType(count int) (values []int) {
	values = make([]int, 0)
	for i := count; i > 0; i-- {
		values = append(values, i)
	}
	return values
}

//无序
func diSortType(start int, end int, count int) (values []int) {
	if end < start || (end-start) < count {
		return nil
	}
	values = make([]int, 0)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(values) < count {
		num := r.Intn((end - start)) + start

		exist := false
		for _, v := range values {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			values = append(values, num)
		}
	}

	return values
}

//主程序
//命令行执行
//go build src/sorter/sorter.go
//go install sorter
//bin/sorter -n 100 -t diSort -i unsorted.dat -o sorted.dat -a qSort

func main() {
	//获取命令行输入
	flag.Parse()
	if infile != nil {
		fmt.Println("values number is", *num, ",sorting type is", *sortType, ",inPush file is", *infile, ",outPush file is", *outfile, ",select algorithm:", *algorithm)
	} else {
		fmt.Println("inPush err!!")
		return
	}

	//生成随机数文件
	// *num, *sortType, *infile 指针的解引用
	inFileName, err := generateFile(*num, *sortType, *infile)
	if err != nil {
		fmt.Println("generateFile err!!")
		return
	}

	//获取文件中的内容
	values, err := readValues(inFileName)
	if err == nil {
		t1 := time.Now()

		//排序
		switch *algorithm {
		case "qSort":
			qsort.QuickSort(values)
		case "bubbleSort":
			bubblesort.BubbleSort(values)
		case "insertSort":
			insertsort.InsertSort(values)
		case "selectSort":
			selectsort.SelectSort(values)

		default:
			fmt.Println("sorting algorithm", *algorithm, "is unknown")
		}

		t2 := time.Now()
		fmt.Println(*algorithm, "sorting algorithm time consuming：", t2.Sub(t1))

		//输出排序后的结果文件
		writeValues(values, *outfile, *algorithm)

	} else {
		fmt.Println(err)
	}
}
