// dup - ２回以上出現する行を出現回数とともに返す
// stdinかargに指定されたファイル一覧から読み込む
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, filename := range files {
			data, err := ioutil.ReadFile(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			for _, line := range strings.Split(string(data), "\n") {
				counts[line]++
			}
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// 宣言よりも前に関数呼び出しがあってもOK
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++ // 値がまだないときは勝手に０値で初期化
	}
}
