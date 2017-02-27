package main

import (
	"fmt"
	"os"
	"path"

	"github.com/fogleman/ease"
	"github.com/fogleman/gg"
)

func render(function ease.Function, filename string, t float64) {
	fmt.Println(filename)
	os.MkdirAll(path.Dir(filename), 0755)
	const P = 32
	const Px = 256
	const W = 1024 + Px*2
	const H = 0 + P*2
	const N = 1024
	dc := gg.NewContext(W, H)
	dc.InvertY()
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.Translate(Px, P)
	dc.Scale(W-Px*2, 1)
	// draw minor grid
	for i := 1; i <= 10; i++ {
		x := float64(i) / 10
		dc.MoveTo(x, -P/2)
		dc.LineTo(x, P/2)
	}
	dc.SetRGBA(0, 0, 0, 0.25)
	dc.SetLineWidth(2)
	dc.Stroke()
	// draw axes
	dc.MoveTo(0, 0)
	dc.LineTo(1, 0)
	dc.MoveTo(0, -P/2)
	dc.LineTo(0, P/2)
	dc.MoveTo(1, -P/2)
	dc.LineTo(1, P/2)
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(3)
	dc.Stroke()
	// draw points
	x := function(t)
	dc.DrawPoint(x, 0, 16)
	dc.SetRGB(0, 0, 0)
	dc.Fill()
	dc.DrawPoint(x, 0, 12)
	dc.SetRGB(1, 0, 0)
	dc.Fill()
	dc.SavePNG(filename)
}

func main() {
	const n = 100
	for i := 0; i < n; i++ {
		t := float64(i) / (n - 1)
		render(ease.Linear, fmt.Sprintf("Linear/%02d.png", i), t)
		render(ease.InQuad, fmt.Sprintf("InQuad/%02d.png", i), t)
		render(ease.OutQuad, fmt.Sprintf("OutQuad/%02d.png", i), t)
		render(ease.InOutQuad, fmt.Sprintf("InOutQuad/%02d.png", i), t)
		render(ease.InCubic, fmt.Sprintf("InCubic/%02d.png", i), t)
		render(ease.OutCubic, fmt.Sprintf("OutCubic/%02d.png", i), t)
		render(ease.InOutCubic, fmt.Sprintf("InOutCubic/%02d.png", i), t)
		render(ease.InQuart, fmt.Sprintf("InQuart/%02d.png", i), t)
		render(ease.OutQuart, fmt.Sprintf("OutQuart/%02d.png", i), t)
		render(ease.InOutQuart, fmt.Sprintf("InOutQuart/%02d.png", i), t)
		render(ease.InQuint, fmt.Sprintf("InQuint/%02d.png", i), t)
		render(ease.OutQuint, fmt.Sprintf("OutQuint/%02d.png", i), t)
		render(ease.InOutQuint, fmt.Sprintf("InOutQuint/%02d.png", i), t)
		render(ease.InSine, fmt.Sprintf("InSine/%02d.png", i), t)
		render(ease.OutSine, fmt.Sprintf("OutSine/%02d.png", i), t)
		render(ease.InOutSine, fmt.Sprintf("InOutSine/%02d.png", i), t)
		render(ease.InExpo, fmt.Sprintf("InExpo/%02d.png", i), t)
		render(ease.OutExpo, fmt.Sprintf("OutExpo/%02d.png", i), t)
		render(ease.InOutExpo, fmt.Sprintf("InOutExpo/%02d.png", i), t)
		render(ease.InCirc, fmt.Sprintf("InCirc/%02d.png", i), t)
		render(ease.OutCirc, fmt.Sprintf("OutCirc/%02d.png", i), t)
		render(ease.InOutCirc, fmt.Sprintf("InOutCirc/%02d.png", i), t)
		render(ease.InElastic, fmt.Sprintf("InElastic/%02d.png", i), t)
		render(ease.OutElastic, fmt.Sprintf("OutElastic/%02d.png", i), t)
		render(ease.InOutElastic, fmt.Sprintf("InOutElastic/%02d.png", i), t)
		render(ease.InBack, fmt.Sprintf("InBack/%02d.png", i), t)
		render(ease.OutBack, fmt.Sprintf("OutBack/%02d.png", i), t)
		render(ease.InOutBack, fmt.Sprintf("InOutBack/%02d.png", i), t)
		render(ease.InBounce, fmt.Sprintf("InBounce/%02d.png", i), t)
		render(ease.OutBounce, fmt.Sprintf("OutBounce/%02d.png", i), t)
		render(ease.InOutBounce, fmt.Sprintf("InOutBounce/%02d.png", i), t)
	}
}
