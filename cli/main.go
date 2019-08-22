package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
)

func loadWord(path string) (*words, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var data []string

	for _, v := range strings.Split(string(b), "\n") {
		if v == "" {
			continue
		}
		data = append(data, v)
	}

	return &words{
		Mutex: sync.RWMutex{},
		word:  data,
	}, nil

}

type words struct {
	Mutex sync.RWMutex
	word  []string
}

func main() {
	wordPath := os.Getenv("WORD_FILE")
	if wordPath == "" {
		fmt.Println("specify word directory to os environment WORD_DIR")
		os.Exit(1)
	}

	word, err := loadWord(wordPath)
	if err != nil {
		log.Fatal(err)
	}

	targetFile := os.Getenv("TARGET_FILE")
	if targetFile == "" {
		fmt.Println("specify target file path to os environment TARGET_FILE")
		os.Exit(1)
	}

	var fp *os.File
	fp, err = os.Open(targetFile)

	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = fp.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	r := bufio.NewReader(fp)
	cnt := 0
	for {
		cnt++
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		for _, w := range word.word {
			word.Mutex.RLock()
			if strings.Contains(line, w) {
				fmt.Printf("line %d contains word %s\n", cnt, w)
			}
		}
		word.Mutex.RUnlock()
	}
}
