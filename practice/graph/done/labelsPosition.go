package main

import (
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

func main() {
	p := plot.New()
	p.Title.Text = "Title"
	p.X.Label.Text = "X [mm]"
	p.Y.Label.Text = "Y [A.U.]"
	p.X.Label.Position = draw.PosRight
	p.Y.Label.Position = draw.PosTop
	p.X.Min = -10
	p.X.Max = +10
	p.Y.Min = -10
	p.Y.Max = +10

	err := p.Save(10*vg.Centimeter, 10*vg.Centimeter, "../testdata/axis_labels.png")
	if err != nil {
		log.Fatalf("could not save plot: %+v", err)
	}
}
