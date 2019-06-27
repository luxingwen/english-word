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
		fmt.Printf("\n\tWord{Name:\"%s\", Ts:\"%s\", Desc:\"%s\"},", item.Name, covertdata(item.Ts), item.Desc)
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

func covertdata(s string) (r string) {
	for _, item := range s {
		r += convert(string(item))
	}
	return
}

func convert(s string) (r string) {
	switch s {
	case "i":
		r = "i"
	case ":":
		r = ":"
	case "e":
		r = "e"
	case "5":
		r = "æ"
	case "1":
		r = "ɑ"
	case "6":
		r = "ɔ"
	case "u":
		r = "u"
	case "2":
		r = "ʌ"
	case "3":
		r = "ə"
	case "a":
		r = "a"
	case "o":
		r = "o"
	case "4":
		r = "ɛ"
	case "j":
		r = "j"
	case "p":
		r = "p"
	case "b":
		r = "b"
	case "t":
		r = "t"
	case "d":
		r = "d"
	case "k":
		r = "k"
	case "g":
		r = "g"
	case "m":
		r = "m"
	case "n":
		r = "n"
	case "9":
		r = "ŋ"
	case "l":
		r = "l"
	case "f":
		r = "f"
	case "v":
		r = "v"
	case "0":
		r = "0"
	case "8":
		r = "ð"
	case "s":
		r = "s"
	case "z":
		r = "z"
	case "7":
		r = "ʃ"
	case "=":
		r = "ʒ"
	case "r":
		r = "r"
	case "h":
		r = "h"
	case "w":
		r = "w"
	case "[":
		r = "["
	case "'":
		r = "ˈ"
	case ";":
		r = ","
	case "]":
		r = "]"
	case "-":
		r = "-"
	case "(":
		r = "("
	case ")":
		r = ")"
	case ",":
		r = ","
	case " ":
		r = " "
	default:
		r = " "
	}
	return
}
