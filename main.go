package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/paldraken/book_thief/pkg/export"
	"github.com/paldraken/book_thief/pkg/parse"
	"github.com/paldraken/book_thief/pkg/parse/types"
)

func main() {
	downloadAndSave()
}

func downloadAndSave() {
	startAt := time.Now()

	//url := "https://author.today/work/210338"
	// url := "https://author.today/work/210575"
	url := "https://author.today/work/26323" // -- с картинками

	p, err := parse.Newparse(url)

	if err != nil {
		log.Fatal(err)
	}

	pbi, err := p.Parse(url, &types.Config{Username: "", Password: ""})

	if err != nil {
		log.Fatalln(err)
	}

	data, err := export.ToFormat(pbi, "FB2")
	if err != nil {
		log.Fatalln(err)
	}

	f, err := openFile()
	if err != nil {
		log.Fatalln(err)
	}

	if _, err := f.Write(data); err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	fmt.Printf("%.2fs elapsed\n", time.Since(startAt).Seconds())
}

func openFile() (*os.File, error) {
	filePath := "d:\\tmp\\test.fb2"

	_, err := os.Stat(filePath)
	if err == nil {
		err = os.Remove(filePath)
		if err != nil {
			return nil, err
		}
	}
	f, err := os.Create(filePath)
	if err != nil {

		return nil, err
	}
	return f, nil
}
