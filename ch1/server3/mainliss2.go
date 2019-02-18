//server2 is a minimal "echo" server, it can calc access count
//You could ask 'http://ip:8000?cycles=10&colorcount=10' to access the multi color gif.
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
	"time"
	//"os"
)

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
	//nColor = 10
)

type lconfig struct {
	cycles  int     //number of complete x oscillator revolutions
	res     float64 //angular resolution
	size    int     //image canvas convers [-size..+size]
	nframes int     //number of animation frames
	delay   int     //delay between frames in 10ms units
	freq    float64 //relative frequency of y oscillator
}

func main() {
	rand.Seed(time.Now().UnixNano())
	lconf := lconfig{
		cycles:  5,
		res:     0.001,
		size:    100,
		nframes: 64,
		delay:   8,
		freq:    rand.Float64() * 3.0,
	}

	var palette = []color.Color{color.White, color.Black}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		confs := r.URL.Query()
		for i, c := range confs {
			switch i {
			case "cycles":
				lconf.cycles, _ = strconv.Atoi(c[0])
			case "res":
				lconf.res, _ = strconv.ParseFloat(c[0], 64)
			case "size":
				lconf.size, _ = strconv.Atoi(c[0])
			case "nframes":
				lconf.nframes, _ = strconv.Atoi(c[0])
			case "delay":
				lconf.delay, _ = strconv.Atoi(c[0])
			case "freq":
				lconf.freq, _ = strconv.ParseFloat(c[0], 64)
			case "colorcount":
				_, palette = palette[len(palette)-1], palette[:len(palette)-1]
				nColor, _ := strconv.Atoi(c[0])
				for i := 1; i < nColor; i++ {
					r := uint8(rand.Uint32() % 256)
					g := uint8(rand.Uint32() % 256)
					b := uint8(rand.Uint32() % 256)
					palette = append(palette, color.RGBA{r, g, b, 0xff})
				}

			}
		}

		fmt.Println(len(palette))
		lissajous(w, lconf, palette)
	}) //each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, set lconfig, palette []color.Color) {
	anim := gif.GIF{LoopCount: set.nframes}
	phase := 0.0 //phase difference
	for i := 0; i < set.nframes; i++ {
		rect := image.Rect(0, 0, 2*set.size+1, 2*set.size+1)

		img := image.NewPaletted(rect, palette)

		var ncolor uint8
		if len(palette) > 2 {
			ncolor = uint8(i%len(palette)-1) + 1 //0~9 -> 1->9
		} else {
			ncolor = blackIndex
		}

		for t := 0.0; t < float64(set.cycles)*2*math.Pi; t += set.res {
			x := math.Sin(t)
			y := math.Sin(t*set.freq + phase)
			img.SetColorIndex(set.size+int(x*float64(set.size)+0.5), set.size+int(y*float64(set.size)+0.5),
				ncolor)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, set.delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) //NOTE: ignoring encoding errors
}
