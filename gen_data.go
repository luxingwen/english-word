package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Word struct {
	Name string
	Ts   string
	Desc string
}

func getWordList(filename string) (rs []*Word, err error) {
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
		if index == 0 {
			index++
			continue
		}
		line = strings.TrimSpace(line)
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
			rs = append(rs, &Word{Name: w, Ts: t, Desc: m})
		}
	}
	return
}

func main() {
	genDataFile()
}

func genDataFile() {
	filename := "books/qqssbdc/xglyy/ck-n5.bok"
	words, err := getWordList(filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("package main\n\n")

	fmt.Printf("var wordlist []Word = []Word{")
	for _, item := range words {
		fmt.Printf("\n\tWord{Name:\"%s\", Desc:\"%s\"},", item.Name, item.Desc)
	}
	fmt.Printf("\n}\n")
}

func genData() {
	var words []*Word

	fs, err := ioutil.ReadDir("books/qqssbdc")
	if err != nil {
		log.Fatal(err)
	}
	var files []string
	for _, item := range fs {
		if item.IsDir() {
			fs1, err := ioutil.ReadDir("books/qqssbdc/" + item.Name())
			if err != nil {
				log.Fatal(err)
			}
			for _, item1 := range fs1 {
				if item1.Name() == "dirname" {
					continue
				}
				files = append(files, "books/qqssbdc/"+item.Name()+"/"+item1.Name())
			}
		}
	}

	for _, filename := range files {
		words1, err := getWordList(filename)
		if err != nil {
			log.Fatal(err)
		}
		words = append(words, words1...)
	}

	mw := make(map[string]string, 0)
	for _, item := range words {
		if strings.Contains(item.Name, "\"") {
			item.Name = strings.ReplaceAll(item.Name, "\"", "\\\"")
		}
		if strings.Contains(item.Desc, "\\") {
			item.Desc = strings.ReplaceAll(item.Desc, "\\", "")
		}
		if strings.Contains(item.Desc, "\"") {
			item.Desc = strings.ReplaceAll(item.Desc, "\"", "\\\"")
		}
		mw[item.Name] = item.Desc
	}

	fmt.Println("package main\n\n")

	fmt.Printf("var mWordsAll map[string]string = map[string]string{")
	for k, v := range mw {
		fmt.Printf("\n\t\"%s\":\"%s\",", k, v)
	}
	fmt.Printf("\n}\n")
}
