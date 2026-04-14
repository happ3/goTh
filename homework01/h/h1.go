package main

import "fmt"

/**
只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。

回文数：判断一个整数是否是回文数

考察：数字操作、条件判断
*/

func main() {

	var arrInt = []int{1, 2, 5, 4, 4, 5, 7, 8, 9}

	var umap = make(map[int]int)
	for i := 0; i < len(arrInt); i++ {
		arr := arrInt[i]
		if val, ok := umap[arr]; ok {
			val++
			umap[arr] = val
		} else {
			umap[arr] = 1
		}
	}

	for k, v := range umap {
		if v < 2 {
			fmt.Println(k)
		}

	}
}
