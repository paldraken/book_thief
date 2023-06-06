package download

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/paldraken/book_thief/internal/export"
	"github.com/paldraken/book_thief/internal/parse"
	"github.com/spf13/viper"
)

func Console(url string) {
	startAt := time.Now()

	p, err := parse.Newparse(url)

	if err != nil {
		log.Fatal(err)
	}

	pbi, err := p.Parse(url)

	if err != nil {
		log.Fatalln(err)
	}

	data, err := export.ToFormat(pbi, "FB2")
	if err != nil {
		log.Fatalln(err)
	}

	f, err := openFile(viper.GetString("output"), pbi.Title+".fb2")
	if err != nil {
		log.Fatalln(err)
	}

	if _, err := f.Write(data); err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	fmt.Printf("%.2fs elapsed\n", time.Since(startAt).Seconds())

}

func openFile(dir, fileName string) (*os.File, error) {

	path := filepath.Join(dir, fileName)

	fmt.Println("save to", path)

	_, err := os.Stat(path)
	if err == nil {
		err = os.Remove(path)
		if err != nil {
			return nil, err
		}
	}
	f, err := os.Create(path)
	if err != nil {

		return nil, err
	}
	return f, nil
}
