package main

import (
	"bufio"
	"errors"
	"flag"
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

type BDC struct {
	*flag.FlagSet
	ttsforlder string
	wordfile   string
	h          bool
	winCmd     bool
	mWav       map[string]string
	book       *Book
	buf        *bufio.Reader
}

type Book struct {
	Name    string
	PageNum int
	Page    int
	Num     int
	Words   []Word
}

func main() {
	bdc := &BDC{mWav: make(map[string]string, 0)}
	bdc.FlagSet = flag.NewFlagSet("BDC Global Params", flag.ContinueOnError)
	bdc.StringVar(&bdc.ttsforlder, "tts", "", "语言文件目录")
	bdc.StringVar(&bdc.wordfile, "data", "", "词库文件")
	bdc.BoolVar(&bdc.h, "h", false, "帮助")
	bdc.BoolVar(&bdc.winCmd, "winCmd", false, "windows cmd命令行")
	if len(os.Args) > 1 {
		err := bdc.Parse(os.Args[1:])
		if err != nil {
			log.Fatal(err)
		}
	}

	if bdc.h {
		bdc.Usage()
		os.Exit(0)
	}

	bdc.init()
	fmt.Println("词库：", bdc.book.Name)
	for {
		var op int
		mainPrint()
		fmt.Scanf("%d", &op)
		if bdc.winCmd {
			bdc.buf.ReadBytes('\n')
		}

		switch op {
		case 1:
			bdc.listWord()
		case 2:
			bdc.remerberWord()
		case 3:
			bdc.listenWrite()
		default:
		}
	}
}

func (b *BDC) init() {
	if b.ttsforlder != "" {
		b.initmWav(b.ttsforlder)
	} else {
		fmt.Println("未加载语音文件")
	}
	book := &Book{Name: "默认词库", Num: len(wordlist), Page: 1, PageNum: 20, Words: wordlist}

	if b.wordfile != "" {
		name, datas, err := getWordList(b.wordfile)
		if err != nil {
			b.Usage()
			log.Fatal(err)
		}
		book.Name = name
		book.Words = datas
		book.Num = len(datas)
	}
	b.buf = bufio.NewReader(os.Stdin)
	b.book = book
}

func (b *BDC) inputPageNumAndPage() {
	fmt.Printf("每组多少个？（当前是%d个,单词总数:%d）:", b.book.PageNum, b.book.Num)
	fmt.Scanf("%d", &num)
	if b.winCmd {
		b.buf.ReadBytes('\n')
	}

	fmt.Printf("第几组(当前是第%d组):", b.book.Page)
	fmt.Scanf("%d", &page)
	if b.winCmd {
		b.buf.ReadBytes('\n')
	}
	if page <= 0 {
		page = b.book.Page
	}

	if num <= 0 {
		num = b.book.PageNum
	}
	start = (page - 1) * num
	end = start + num
	if start >= b.book.Num {
		start = b.book.Num - 1 - num
	}
	if end > b.book.Num {
		end = b.book.Num - 1
	}
	b.book.Page = page
	b.book.PageNum = num
}

var (
	num   int
	page  int
	start int
	end   int
)

func (b *BDC) listWord() {
	b.inputPageNumAndPage()
	m := make(map[string]string, 0)
	for _, item := range b.book.Words[start:end] {
		m[item.Name] = item.Desc
	}
	for k, v := range m {
		fmt.Printf("%20s: %s\n", k, v)
	}
	fmt.Println("开始浏览-->")
	for len(m) > 0 {
		for k, v := range m {
			fmt.Println(v)
			b.speaker(k)
			line, _, err := b.buf.ReadLine()
			if err != nil {
				log.Fatal(err)
			}
			key := strings.TrimSpace(string(line))
			for key != k {
				b.speaker(k)
				fmt.Println(v)
				line, _, err := b.buf.ReadLine()
				if err != nil {
					log.Fatal(err)
				}
				key = strings.TrimSpace(string(line))
			}
			delete(m, key)
		}
	}
}

func (b *BDC) remerberWord() {
	b.inputPageNumAndPage()
	m := make(map[string]string, 0)
	for _, item := range b.book.Words[start:end] {
		m[strings.TrimSpace(item.Name)] = item.Desc
	}
	for len(m) > 0 {
		for k, v := range m {
			fmt.Println(v)
			line, _, err := b.buf.ReadLine()
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

func (b *BDC) listenWrite() {
	b.inputPageNumAndPage()
	m := make(map[string]string, 0)
	for _, item := range b.book.Words[start:end] {
		m[strings.TrimSpace(item.Name)] = item.Desc
	}
	for len(m) > 0 {
		for k, v := range m {
			if !b.speaker(k) {
				fmt.Println(v)
			}
			line, _, err := b.buf.ReadLine()
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

func (b *BDC) initmWav(folder string) {
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
				b.mWav[ns[0]] = folder + "/" + item.Name() + "/" + item1.Name()
			}
		}
	}
}

func (b *BDC) speaker(words string) (ok bool) {
	for _, word := range strings.Split(words, " ") {
		if _, ok = b.mWav[word]; !ok {
			return false
		}
	}
	for _, word := range strings.Split(words, " ") {
		v, _ := b.mWav[word]
		err := exec.Command("./playerwav.exe", v).Run()
		if err != nil {
			log.Fatal(err)
		}
	}
	return true
}

func getWordList(filename string) (bookName string, rs []Word, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()
	buf := bufio.NewReader(file)

	wordFunc := func(s string, indexStr string) (r string, err error) {
		index := strings.Index(s, indexStr)
		if index < 0 {
			return "", errors.New("not found > " + indexStr)
		}
		index += len(indexStr)
		end := index + strings.Index(s[index:], "[")
		if end < index {
			return s[index:], nil
		}
		return s[index:end], nil
	}

	index := 0
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)

		if index == 0 {
			bookName, err = wordFunc(line, "[N]")
			if err != nil {
				return "", nil, err
			}
			index++
			continue
		}
		w, err := wordFunc(line, "[W]")
		if err != nil {
			continue
		}
		t, err := wordFunc(line, "[T]")
		if err != nil {
			continue
		}
		m, err := wordFunc(line, "[M]")
		if err != nil {
			continue
		}
		if line != "" {
			rs = append(rs, Word{Name: w, Ts: t, Desc: m})
		}
	}
	return
}
