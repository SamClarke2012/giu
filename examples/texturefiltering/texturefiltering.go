package main

import (
	_ "image/jpeg"
	_ "image/png"

	g "github.com/AllenDang/giu"
)

var (
	spriteTexture *g.Texture
	largeTexture  *g.Texture
)

func loop() {
	g.SingleWindow("load image").Layout(
		g.Column(
			g.Label("15x20 pixel image"),
			g.Row(
				g.Column(
					g.Label("50%"),
					g.Image(spriteTexture).Size(8, 10),
				),
				g.Column(
					g.Label("100%"),
					g.Image(spriteTexture).Size(15, 20),
				),
				g.Column(
					g.Label("800%"),
					g.Image(spriteTexture).Size(120, 160),
				),
			),
		),
		g.Column(
			g.Label("215x140 image"),
			g.Row(
				g.Column(
					g.Label("50%"),
					g.Image(largeTexture).Size(215/2, 140/2),
				),
				g.Column(
					g.Label("100%"),
					g.Image(largeTexture).Size(215, 140),
				),
				g.Column(
					g.Label("200%"),
					g.Image(largeTexture).Size(215*2, 140*2),
				),
			),
		),
		g.Row(
			g.Button("Minify Filter Nearest").OnClick(func() {
				_ = g.Context.GetRenderer().SetTextureMinFilter(g.TextureFilterNearest)
			}),
			g.Button("Minify Filter Linear").OnClick(func() {
				_ = g.Context.GetRenderer().SetTextureMinFilter(g.TextureFilterLinear)
			}),
			/*g.Button("Nearest Mipmap Nearest", func() {
				g.Context.GetRenderer().SetTextureMinFilter(g.TextureFilterNearestMipmapNearest)
			}),
			g.Button("Linear Mipmap Nearest", func() {
				g.Context.GetRenderer().SetTextureMinFilter(g.TextureFilterLinearMipmapNearest)
			}),
			g.Button("Nearest Mipmap Linear", func() {
				g.Context.GetRenderer().SetTextureMinFilter(g.TextureFilterNearestMipmapLinear)
			}),
			g.Button("Linear Mipmap Linear", func() {
				g.Context.GetRenderer().SetTextureMinFilter(g.TextureFilterLinearMipmapLinear)
			}),*/
		),
		g.Row(
			g.Button("Magnify Filter Nearest").OnClick(func() {
				_ = g.Context.GetRenderer().SetTextureMagFilter(g.TextureFilterNearest)
			}),
			g.Button("Magnify Filter Linear").OnClick(func() {
				_ = g.Context.GetRenderer().SetTextureMagFilter(g.TextureFilterLinear)
			}),
		),
	)
}

func main() {
	wnd := g.NewMasterWindow("Texture Filtering", 800, 600, g.MasterWindowFlagsNotResizable)

	spriteImg, _ := g.LoadImage("gopher-sprite.png")
	largeImg, _ := g.LoadImage("gopher.png")
	go func() {
		spriteTexture, _ = g.NewTextureFromRgba(spriteImg)
		largeTexture, _ = g.NewTextureFromRgba(largeImg)
	}()

	wnd.Run(loop)
}
