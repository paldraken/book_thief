package main

import (
	"fmt"
	"log"
	"time"

	"github.com/paldraken/book_thief/pkg/parse"
)

func main() {
	startAt := time.Now()

	url := "https://author.today/work/210338"

	p, err := parse.Newparse(url)

	if err != nil {
		log.Fatal(err)
	}

	p.Parse(url)

	fmt.Printf("%.2fs elapsed\n", time.Since(startAt).Seconds())
}
