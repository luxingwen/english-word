package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

type Word struct {
	Name string
	Ts   string
	Desc string
}

func main() {
	mWav = make(map[string]string, 0)
	if len(os.Args) == 2 {
		initmWav(os.Args[1])
	}

	for {
		var op int
		mainPrint()
		fmt.Scanf("%d", &op)
		switch op {
		case 1:
			listWord()
		case 2:
			remerberWord()
		case 3:
			listenWrite()
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
			speaker(k)
			fmt.Println(v)
			line, _, err := buf.ReadLine()
			if err != nil {
				log.Fatal(err)
			}
			key := strings.TrimSpace(string(line))
			for key != k {
				speaker(k)
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

func listenWrite() {
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
			if !speaker(k) {
				fmt.Println(v)
			}
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

var mWav map[string]string

func initmWav(folder string) {
	fs, err := ioutil.ReadDir(folder)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range fs {
		if item.IsDir() {

			fs1, err := ioutil.ReadDir(folder + "/" + item.Name())
			if err != nil {
				log.Fatal(err)
			}

			for _, item1 := range fs1 {
				ns := strings.Split(item1.Name(), ".")
				mWav[ns[0]] = folder + "/" + item.Name() + "/" + item1.Name()
			}
		}
	}
}

func speaker(words string) (ok bool) {
	for _, word := range strings.Split(words, " ") {
		if _, ok = mWav[word]; !ok {
			return false
		}
	}
	for _, word := range strings.Split(words, " ") {
		v, _ := mWav[word]
		err := exec.Command("./playerwav.exe", v).Run()
		if err != nil {
			log.Fatal(err)
		}
	}
	return true
}
