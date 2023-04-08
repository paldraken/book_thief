package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/paldraken/book_thief/pkg/parse"
)

type fb2Body struct {
	Title   string `xml:"title"`
	Content string `xml:",innerxml"`
}

type body struct {
	Chapters []*fb2Body `xml:"section"`
}

func main() {

	startAt := time.Now()

	url := "https://author.today/work/210338"

	p, err := parse.Newparse(url)

	if err != nil {
		log.Fatal(err)
	}

	pbi, err := p.Parse(url)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(pbi)

	// data, err := export.ToFormat(pbi, "FB2")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// f, err := openFile()
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// if _, err := f.Write(data); err != nil {
	// 	log.Fatalln(err)
	// }
	// defer f.Close()

	fmt.Printf("%.2fs elapsed\n", time.Since(startAt).Seconds())
}

func openFile() (*os.File, error) {
	filePath := "d:\\tmp\\test.fb2"

	_, statErr := os.Stat(filePath)
	if statErr != nil && !errors.Is(statErr, os.ErrNotExist) {
		return nil, statErr
	}

	err := os.Remove(filePath)
	if err != nil {
		return nil, err
	}

	if f, err := os.Create(filePath); err != nil {
		return nil, err
	} else {
		return f, nil
	}
}
