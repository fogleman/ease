package main

import (
	"fmt"

	"github.com/fogleman/ease"
	"github.com/fogleman/gg"
)

func render(function ease.Function, path, title string) {
	fmt.Println(path)
	const P = 256
	const S = 1024 + P*2
	const N = 1024
	dc := gg.NewContext(S, S)
	dc.InvertY()
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.Translate(P, P)
	dc.Scale(S-P*2, S-P*2)
	// draw minor grid
	for i := 1; i <= 10; i++ {
		x := float64(i) / 10
		dc.MoveTo(x, 0)
		dc.LineTo(x, 1)
		dc.MoveTo(0, x)
		dc.LineTo(1, x)
	}
	dc.SetRGBA(0, 0, 0, 0.25)
	dc.SetLineWidth(2)
	dc.Stroke()
	// draw axes
	dc.MoveTo(0, 0)
	dc.LineTo(1, 0)
	dc.LineTo(1, 1)
	dc.LineTo(0, 1)
	dc.ClosePath()
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(3)
	dc.Stroke()
	// draw points
	for i := 0; i < N; i++ {
		x := float64(i) / (N - 1)
		y := function(x)
		dc.LineTo(x, y)
	}
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(12)
	dc.StrokePreserve()
	dc.SetRGB(1, 0, 0)
	dc.SetLineWidth(6)
	dc.Stroke()
	// title
	dc.Identity()
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace("/Users/fogleman/Library/Fonts/DejaVuSansMono.ttf", 48); err != nil {
		panic(err)
	}
	dc.DrawStringAnchored(title, S/2, P/4, 0.5, 0.5)
	dc.SavePNG(path)
}

func main() {
	render(ease.Linear, "Linear.png", "ease.Linear(t)")
	render(ease.InQuad, "InQuad.png", "ease.InQuad(t)")
	render(ease.OutQuad, "OutQuad.png", "ease.OutQuad(t)")
	render(ease.InOutQuad, "InOutQuad.png", "ease.InOutQuad(t)")
	render(ease.InCubic, "InCubic.png", "ease.InCubic(t)")
	render(ease.OutCubic, "OutCubic.png", "ease.OutCubic(t)")
	render(ease.InOutCubic, "InOutCubic.png", "ease.InOutCubic(t)")
	render(ease.InQuart, "InQuart.png", "ease.InQuart(t)")
	render(ease.OutQuart, "OutQuart.png", "ease.OutQuart(t)")
	render(ease.InOutQuart, "InOutQuart.png", "ease.InOutQuart(t)")
	render(ease.InQuint, "InQuint.png", "ease.InQuint(t)")
	render(ease.OutQuint, "OutQuint.png", "ease.OutQuint(t)")
	render(ease.InOutQuint, "InOutQuint.png", "ease.InOutQuint(t)")
	render(ease.InSine, "InSine.png", "ease.InSine(t)")
	render(ease.OutSine, "OutSine.png", "ease.OutSine(t)")
	render(ease.InOutSine, "InOutSine.png", "ease.InOutSine(t)")
	render(ease.InExpo, "InExpo.png", "ease.InExpo(t)")
	render(ease.OutExpo, "OutExpo.png", "ease.OutExpo(t)")
	render(ease.InOutExpo, "InOutExpo.png", "ease.InOutExpo(t)")
	render(ease.InCirc, "InCirc.png", "ease.InCirc(t)")
	render(ease.OutCirc, "OutCirc.png", "ease.OutCirc(t)")
	render(ease.InOutCirc, "InOutCirc.png", "ease.InOutCirc(t)")
	render(ease.InElastic, "InElastic.png", "ease.InElastic(t)")
	render(ease.OutElastic, "OutElastic.png", "ease.OutElastic(t)")
	render(ease.InOutElastic, "InOutElastic.png", "ease.InOutElastic(t)")
	render(ease.InBack, "InBack.png", "ease.InBack(t)")
	render(ease.OutBack, "OutBack.png", "ease.OutBack(t)")
	render(ease.InOutBack, "InOutBack.png", "ease.InOutBack(t)")
	render(ease.InBounce, "InBounce.png", "ease.InBounce(t)")
	render(ease.OutBounce, "OutBounce.png", "ease.OutBounce(t)")
	render(ease.InOutBounce, "InOutBounce.png", "ease.InOutBounce(t)")
}
