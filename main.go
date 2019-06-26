package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		if v, ok := mWordsAll[os.Args[1]]; ok {
			fmt.Println(v)
		} else {
			fmt.Println("没有找到单词：", os.Args[1])
		}
		return
	}
	count := 0
	for k, v := range mWordsAll {
		fmt.Println(k)
		fmt.Println(v)
		fmt.Println()
		count++
		if count >= 5 {
			break
		}
	}
}
