// server: echo機能とcounter機能とHeader/FormData表示機能
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var mu sync.Mutex
var count int

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/lissajous", func(w http.ResponseWriter, r *http.Request) {
		lissajous(w, r)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

var palette = []color.Color{
	color.White,
	color.RGBA{0xff, 0, 0, 0xff},
	color.RGBA{0, 0xff, 0, 0xff},
	color.RGBA{0, 0, 0xff, 0xff},
}

const (
	bgIndex = 0
	rIndex  = 1
	gIndex  = 2
	bIndex  = 3
)

func lissajous(out io.Writer, r *http.Request) {
	const (
		res     = 0.001 // 分解能
		size    = 100   // 画像キャンパスの範囲
		nframes = 64    // アニメーションフレーム数
		delay   = 8     // 10ms単位でフレーム間遅延
	)
	rcycles := 2.0
	gcycles := 2.0
	bcycles := 2.0

	queries := r.URL.Query()
	if q, ok := queries["rcycles"]; ok {
		intval, err := strconv.Atoi(q[0])
		if err == nil {
			rcycles = float64(intval)
		}
	}
	if q, ok := queries["gcycles"]; ok {
		intval, err := strconv.Atoi(q[0])
		if err == nil {
			gcycles = float64(intval)
		}
	}
	if q, ok := queries["bcycles"]; ok {
		intval, err := strconv.Atoi(q[0])
		if err == nil {
			bcycles = float64(intval)
		}
	}

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		t := 0.0
		acc := 0.0
		for ; t < float64(rcycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), rIndex)
		}
		acc += t
		for ; t < acc+float64(gcycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), gIndex)
		}
		acc += t
		for ; t < acc+float64(bcycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), bIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
