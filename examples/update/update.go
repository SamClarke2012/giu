package main

import (
	"os"
	"fmt"
	"time"
	"math"
	"image"
	"image/color"
	"image/draw"
	"github.com/AllenDang/giu"
	"github.com/disintegration/imaging"
	//"code.google.com/p/graphics-go/graphics"
)

var (
	d float64
	counter  float64
	buffer   []float64
	ticks    []giu.PlotTicker
	font     *giu.FontInfo
	font1    *giu.FontInfo
	gaugeTex *giu.Texture
	face     image.Image
	needle   image.Image
	combo   *image.RGBA
)

func refresh() {
	// ~60 fps
	ticker := time.NewTicker(time.Second / 60)

	for {
		counter += d
		if counter > 0.98 {
			d *= -1
		} else if counter < 0.01 {
			d *= -1
		}
		giu.Update()

		<-ticker.C
	}
}

func loop() {
	giu.SetMouseCursor(giu.MouseCursorNone)
	giu.SingleWindow("Update").Layout(
		giu.Label("Below number is updated by a goroutine").Font(font1),
		giu.Style().SetFont(font).To(
			giu.Label(fmt.Sprintf("%.3f", counter)),
		),
		giu.ProgressBar(float32(counter)).Size(-1,0),
		giu.Row(
			giu.Custom(func() {
				canvas := giu.GetCanvas()
				pos := giu.GetCursorScreenPos()
				col := color.RGBA{63,191,191,255}
				p2 :=pos
				p1 := pos.Add(image.Pt(500+int(50*counter),int(50*counter)))
				canvas.AddCircleFilled(p1, float32(35*counter), col)
				if gaugeTex != nil {
					canvas.AddImage(gaugeTex, p2.Add(image.Pt(520,50)), p2.Add(image.Pt(701,231)))
				}
				go func() {
					combo = comboRotate(face, needle, -240*counter)
					gaugeTex, _ = giu.NewTextureFromRgba(combo)
					buffer = buffer[1:]
					buffer = append(buffer, math.Cos(counter*1.5))
				}()
			}),
			giu.Plot("Stuff Over Time").AxisLimits(0,1000,0,1.1, giu.ConditionOnce).Size(490, 350).Plots(
				giu.PlotLine("Is this here?", buffer),
				giu.PlotScatter("Weeee", buffer),
			),
		),
	)
}

func main() {
	d = 0.01
	counter = 0.0
	font = giu.AddFont("Orbitron-Medium.ttf",64)
	font1 = giu.AddFont("/usr/share/fonts/truetype/liberation/LiberationMono-Regular.ttf",32)
	wnd := giu.NewMasterWindow("Update", 800, 480, giu.MasterWindowFlagsFrameless)
	face = openImage("1x8_face.png")
	needle = openImage("1x8_needle.png")
	d =0.01
	for i := 0.0; i<10; i+=d {
		buffer = append(buffer, math.Sin(i))
	}
	go func() {
		combo = comboRotate(face, needle, 0)
		gaugeTex, _ = giu.NewTextureFromRgba(combo)
	}()
	go refresh()
	wnd.Run(loop)
}

func comboRotate(face image.Image, needle image.Image, angle float64) *image.RGBA {
	// Background face is 194x184
	// Needle is 181x181
	// 7x1 and 188x182 dest rectangle
	nrotated := imaging.Rotate(needle, angle, color.Transparent)
	centered := imaging.PasteCenter(face, nrotated)
	rgba := image.NewRGBA(face.Bounds())
	draw.Draw(rgba, face.Bounds(), face, image.ZP, draw.Src)
	draw.Draw(rgba, face.Bounds(), centered, image.ZP, draw.Over)
	return rgba
}

func openImage(file string) image.Image {
	imgFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	img, _, err := image.Decode(imgFile)
	if err != nil {
		fmt.Println(err)
	}
	return img
}
