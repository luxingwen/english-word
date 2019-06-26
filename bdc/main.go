package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Word struct {
	Name string
	Ts   string
	Desc string
}

func main() {
	for {
		var op int
		mainPrint()
		fmt.Scanf("%d", &op)
		switch op {
		case 1:
			listWord()
		case 2:
			remerberWord()
		default:
		}
	}
}

var (
	num   int
	page  int
	start int
	end   int
)

func listWord() {
	fmt.Printf("每组多少个(%d):", len(wordlist))
	fmt.Scanf("%d", &num)
	fmt.Print("第几组:")
	fmt.Scanf("%d", &page)
	if page <= 0 {
		page = 1
	}

	if num <= 0 {
		num = 25
	}
	start = (page - 1) * num
	end = start + num
	if start >= len(wordlist) {
		start = len(wordlist) - 1 - num
	}
	if end > len(wordlist) {
		end = len(wordlist) - 1
	}
	m := make(map[string]string, 0)
	for _, item := range wordlist[start:end] {
		m[item.Name] = item.Desc
	}
	for k, v := range m {
		fmt.Printf("%20s: %s\n", k, v)
	}

	fmt.Println("开始浏览-->")
	buf := bufio.NewReader(os.Stdin)
	for len(m) > 0 {
		for k, v := range m {
			fmt.Println(v)
			line, _, err := buf.ReadLine()
			if err != nil {
				log.Fatal(err)
			}
			key := strings.TrimSpace(string(line))
			for key != k {
				fmt.Println(v)
				line, _, err := buf.ReadLine()
				if err != nil {
					log.Fatal(err)
				}
				key = strings.TrimSpace(string(line))
			}
			delete(m, key)
		}
	}
}

func remerberWord() {
	fmt.Printf("每组多少个(%d):", len(wordlist))
	fmt.Scanf("%d", &num)
	fmt.Print("第几组:")
	fmt.Scanf("%d", &page)
	if page <= 0 {
		page = 1
	}

	if num <= 0 {
		num = 25
	}

	start = (page - 1) * num
	end = start + num
	if start >= len(wordlist) {
		start = len(wordlist) - 1 - num
	}
	if end > len(wordlist) {
		end = len(wordlist) - 1
	}

	m := make(map[string]string, 0)
	for _, item := range wordlist[start:end] {
		m[strings.TrimSpace(item.Name)] = item.Desc
	}
	buf := bufio.NewReader(os.Stdin)
	for len(m) > 0 {
		for k, v := range m {
			fmt.Println(v)
			line, _, err := buf.ReadLine()
			if err != nil {
				log.Fatal(err)
			}
			key := strings.TrimSpace(string(line))
			if key == k {
				delete(m, key)
			}
		}
		for k, v := range m {
			fmt.Printf("%20s : %s\n", v, k)
		}
	}
}

func mainPrint() {
	fmt.Println("1.浏览")
	fmt.Println("2.默写记忆")
	fmt.Println("3.听写记忆")
}
