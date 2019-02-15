//lissajous draw gif
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

//var palette = []color.Color{color.White,color.Black}
//exec 1.5
//var palette = []color.Color{color.White,color.RGBA{0x0,0xff,0x0,0xff},color.RGBA{0xff,0x00,0x00,0xff},color.RGBA{0x00,0x00,0xff,0xff}}

const (
	//whiteIndex = 0 // first color in palette
	//blackIndex = 1 // next color in palette
	nColor = 10
)

func main() {
	seed := time.Now()
	rand.Seed(seed.Unix())

	var palette []color.Color
	palette = append(palette, color.White)
	for i := 1; i < nColor; i++ {
		r := uint8(rand.Uint32() % 256)
		g := uint8(rand.Uint32() % 256)
		b := uint8(rand.Uint32() % 256)
		palette = append(palette, color.RGBA{r, g, b, 0xff})
	}

	lissajous(os.Stdout, palette)
}

func lissajous(out io.Writer, palette []color.Color) {
	const (
		cycles  = 5     //number of complete x oscillator revolutions
		res     = 0.001 //angular resolution
		size    = 100   //image canvas convers [-size..+size]
		nframes = 64    //number of animation frames
		delay   = 8     //delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 //relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 //phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)

		img := image.NewPaletted(rect, palette)
		//exe 1.6
		ncolor := uint8(i%len(palette)-1) + 1 //0~9 -> 1->9

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				ncolor)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) //NOTE: ignoring encoding errors
}
