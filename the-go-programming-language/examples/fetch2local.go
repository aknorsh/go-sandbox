package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

func main() {

	sn, n, err := fetch(os.Args[1])
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	fmt.Printf("%s,%d\n", sn, n)

}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create("examples/fetch/" + local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return local, n, err
}
