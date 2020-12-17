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
	if len(os.Args) != 2 {
		return
	}

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
	if local == "/" || local == "." {
		local = "index.html"
	}
	f, err := os.Create("prac/fetched/" + local)
	if err != nil {
		return "", 0, err
	}
	defer func() {
		clsErr := f.Close()
		if clsErr != nil {
			err = clsErr
		}
	}()
	n, err = io.Copy(f, resp.Body)
	return local, n, err
}
