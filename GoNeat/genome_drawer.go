package GoNeat

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
	"strconv"
)

var g *Genome
var win *pixelgl.Window

func DrawGenome(gen *Genome, window *pixelgl.Window) {
	g = gen
	win = window
	pixelgl.Run(run)
}

func run() {
	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxt := text.New(pixel.V(0, 0), basicAtlas)

	imd := imdraw.New(nil)

	for i := 0; i < g.GetLayers(); i++ {
		for j := range g.GetNodesWithLayer(i + 1) {
			basicTxt.Color = colornames.Black
			//fmt.Fprintf(basicTxt, strconv.Itoa(g.GetNodesWithLayer(i + 1)[j].GetInnovationNumber()) + ", " +
			//strconv.Itoa(g.GetNodesWithLayer(i + 1)[j].GetLayer()))
			fmt.Fprintf(basicTxt, strconv.Itoa(g.GetNodesWithLayer(i + 1)[j].GetInnovationNumber())+", "+
				strconv.FormatFloat(g.GetNodesWithLayer(i + 1)[j].GetWeight(), 'f', 2, 64))
			//fmt.Fprintf(basicTxt, strconv.FormatFloat(g.GetNodesWithLayer(i + 1)[j].GetWeight(), 'f', 2, 64))
			basicTxt.Draw(win, pixel.IM.Moved(pixel.V(
				(float64(i)+0.5)*(win.Bounds().W()/float64(g.GetLayers()))-1,
				(float64(j)+0.5)*(win.Bounds().H()/float64(len(g.GetNodesWithLayer(i+1))))+50)))
			basicTxt.Clear()
			if g.GetNodesWithLayer(i + 1)[j].IsActivated() {
				imd.Color = pixel.RGB(0, 1, 0)
			} else {
				imd.Color = pixel.RGB(1, 0, 0)
			}
			imd.Push(pixel.V(
				(float64(i)+0.5)*(win.Bounds().W()/float64(g.GetLayers())),
				(float64(j)+0.5)*(win.Bounds().H()/float64(len(g.GetNodesWithLayer(i+1))))))
			imd.Circle(20, 40)
			for k := range g.GetNodesWithLayer(i + 1)[j].GetOutwardConnections() {
				imd.Color = pixel.RGB(0, 0, 0)
				imd.Push(
					pixel.V(
						(float64(i)+0.5)*(win.Bounds().W()/float64(g.GetLayers()))+40,
						(float64(j)+0.5)*(win.Bounds().H()/float64(len(g.GetNodesWithLayer(i+1))))),
					pixel.V(
						(float64(g.GetNodesWithLayer(i + 1)[j].GetOutwardConnections()[k].GetNodeB().GetLayer())-0.5)*(win.Bounds().W()/float64(g.GetLayers()))-40,
						(float64(NodeIndex(g.GetNodesWithLayer(g.GetNodesWithLayer(i + 1)[j].GetOutwardConnections()[k].GetNodeB().GetLayer()),
							g.GetNodesWithLayer(i + 1)[j].GetOutwardConnections()[k].GetNodeB()))+0.5)*(win.Bounds().H()/float64(len(g.GetNodesWithLayer(g.GetNodesWithLayer(i + 1)[j].GetOutwardConnections()[k].GetNodeB().GetLayer()))))))
				imd.Line(5)
			}
		}
	}
	imd.Draw(win)
}
