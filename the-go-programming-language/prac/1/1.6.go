// Lissajous: ランダムなリサージュ曲線のGIFアニメを作成(赤青緑)
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

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

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		rcycles = 2
		gcycles = 2
		bcycles = 2
		res     = 0.001 // 分解能
		size    = 100   // 画像キャンパスの範囲
		nframes = 64    // アニメーションフレーム数
		delay   = 8     // 10ms単位でフレーム間遅延
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		t := 0.0
		acc := 0.0
		for ; t < rcycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), rIndex)
		}
		acc += t
		for ; t < acc+gcycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), gIndex)
		}
		acc += t
		for ; t < acc+bcycles*2*math.Pi; t += res {
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
