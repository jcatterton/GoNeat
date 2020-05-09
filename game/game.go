package Game

import (
	Network "NEAT/GoNeat"
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
	"image/png"
	"log"
	"math/rand"
	"os"
	"strconv"
)

var score = 0

func Start() {
	pixelgl.Run(run)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Flappy Bird",
		Bounds: pixel.R(0, 0, 500, 750),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	birdImage, err := loadPicture("./bird.png")
	if err != nil {
		panic(err)
	}

	pipeImage, err := loadPicture("./pipe.png")
	if err != nil {
		panic(err)
	}

	bird := Bird{
		375,
		0,
		*pixel.NewSprite(birdImage, birdImage.Bounds()),
	}

	pipes := make([]*Pipe, 6)
	for i := range pipes {
		if i%2 == 0 {
			pipes[i] = &Pipe{float64(rand.Intn(350) + 100), 600 + float64(200*i), true, *pixel.NewSprite(pipeImage, pipeImage.Bounds())}
		} else {
			pipes[i] = pipes[i-1].CreateSisterPipe()
		}
	}

	imd := imdraw.New(nil)

	pop := Network.InitPopulation(4, 1)

	for !win.Closed() {
		for x := range pop.GetAllGenomes() {
			dead := false
			for !dead {
				win.Clear(colornames.Black)
				if err := pop.GetAllGenomes()[x].TakeInput(
					[]float64{
						bird.GetYVel(),
						bird.GetInformationOnNextPipes(pipes)[0],
						bird.GetInformationOnNextPipes(pipes)[1],
						bird.GetInformationOnNextPipes(pipes)[2],
					}); err != nil {
					panic(err)
				}

				pop.GetAllGenomes()[x].FeedForward()

				output := pop.GetAllGenomes()[x].GetOutputs()[0]

				if win.Pressed(pixelgl.KeyUp) {
					bird.height++
				}
				if win.Pressed(pixelgl.KeyDown) {
					bird.height--
				}

				if output > 0.5 {
					bird.Jump()
				}

				bird.Fall()
				bird.Draw(win, pixel.IM.Moved(pixel.V(win.Bounds().W()/2, bird.height)))

				for i := range pipes {
					pipes[i].Draw(win, pixel.IM)
					pipes[i].MoveLeft()
					if pipes[i].xPos <= -200 && i%2 == 0 {
						pipes[i] = &Pipe{float64(rand.Intn(350) + 100), float64(1000), true, *pixel.NewSprite(pipeImage, pipeImage.Bounds())}
						pipes[i+1] = pipes[i].CreateSisterPipe()
					}
				}
				imd.Clear()
				imd.Color = colornames.White

				imd.Push(bird.Bounds().Center)
				imd.Push(pixel.V(bird.Bounds().Center.X+bird.GetInformationOnNextPipes(pipes)[0], bird.GetHeight()))
				imd.Line(2)
				imd.Draw(win)

				imd.Push(bird.Bounds().Center)
				imd.Push(pixel.V(250, bird.Bounds().Center.Y+bird.Bounds().Radius+bird.GetInformationOnNextPipes(pipes)[1]))
				imd.Line(2)
				imd.Draw(win)

				imd.Push(bird.Bounds().Center)
				imd.Push(pixel.V(250, bird.Bounds().Center.Y-bird.Bounds().Radius+bird.GetInformationOnNextPipes(pipes)[2]))
				imd.Line(2)
				imd.Draw(win)

				drawGenome(pop.GetAllGenomes()[x], win)

				if checkForCollisions(bird, pipes) {
					dead = true
					pop.GetAllGenomes()[x].SetFitness(float64(score))
					log.Println(score)
					score = 0
					bird = Bird{
						375,
						0,
						*pixel.NewSprite(birdImage, birdImage.Bounds()),
					}
					for i := range pipes {
						if i%2 == 0 {
							pipes[i] = &Pipe{float64(rand.Intn(350) + 100), 600.0 + float64(200*i), true, *pixel.NewSprite(pipeImage, pipeImage.Bounds())}
						} else {
							pipes[i] = pipes[i-1].CreateSisterPipe()
						}
					}
				}

				score++

				win.SetTitle("Flappy Bird - " + strconv.Itoa(score))
				win.Update()
			}
		}
		for x := range pop.GetSpecies() {
			pop.GetSpecies()[x].SetChampion()
			pop.GetSpecies()[x].CullTheWeak()
		}
		pop.SetGrandChampion()
		pop.ExtinctionEvent()
		pop.Mutate()
		log.Println(pop.GetGeneration(), " - ", pop.GetGrandChampion().GetFitness())
	}
}

func drawGenome(g *Network.Genome, win *pixelgl.Window) {
	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxt := text.New(pixel.V(0, 0), basicAtlas)

	imd := imdraw.New(nil)

	for i := 0; i < g.GetLayers(); i++ {
		for j := range g.GetNodesWithLayer(i + 1) {
			basicTxt.Color = colornames.White
			if g.GetNodesWithLayer(i + 1)[j].GetInnovationNumber() == 0 {
				fmt.Fprintf(basicTxt, "yVel "+strconv.FormatFloat(g.GetNodesWithLayer(i + 1)[j].GetWeight(), 'f', 2, 64))
			} else if g.GetNodesWithLayer(i + 1)[j].GetInnovationNumber() == 1 {
				fmt.Fprintf(basicTxt, "distance next pipe "+strconv.FormatFloat(g.GetNodesWithLayer(i + 1)[j].GetWeight(), 'f', 2, 64))
			} else if g.GetNodesWithLayer(i + 1)[j].GetInnovationNumber() == 2 {
				fmt.Fprintf(basicTxt, "distance above "+strconv.FormatFloat(g.GetNodesWithLayer(i + 1)[j].GetWeight(), 'f', 2, 64))
			} else if g.GetNodesWithLayer(i + 1)[j].GetInnovationNumber() == 3 {
				fmt.Fprintf(basicTxt, "distance below "+strconv.FormatFloat(g.GetNodesWithLayer(i + 1)[j].GetWeight(), 'f', 2, 64))
			}
			w := win.Bounds().W() / 2
			h := win.Bounds().H() / 3
			basicTxt.Draw(win, pixel.IM.Moved(pixel.V(
				(float64(i)+0.5)*(w/float64(g.GetLayers()))-1,
				(float64(j)+0.5)*(h/float64(len(g.GetNodesWithLayer(i+1))))+20)))
			basicTxt.Clear()
			if g.GetNodesWithLayer(i + 1)[j].IsActivated() {
				imd.Color = pixel.RGB(0, 1, 0)
			} else {
				imd.Color = pixel.RGB(1, 0, 0)
			}
			imd.Push(pixel.V(
				(float64(i)+0.5)*(w/float64(g.GetLayers())),
				(float64(j)+0.5)*(h/float64(len(g.GetNodesWithLayer(i+1))))))
			imd.Circle(5, 20)
			for k := range g.GetNodesWithLayer(i + 1)[j].GetOutwardConnections() {
				imd.Color = pixel.RGB(1, 1, 1)
				imd.Push(
					pixel.V(
						(float64(i)+0.5)*(w/float64(g.GetLayers()))+10,
						(float64(j)+0.5)*(h/float64(len(g.GetNodesWithLayer(i+1))))),
					pixel.V(
						(float64(g.GetNodesWithLayer(i + 1)[j].GetOutwardConnections()[k].GetNodeB().GetLayer())-0.5)*(w/float64(g.GetLayers()))-10,
						(float64(Network.NodeIndex(g.GetNodesWithLayer(g.GetNodesWithLayer(i + 1)[j].GetOutwardConnections()[k].GetNodeB().GetLayer()),
							g.GetNodesWithLayer(i + 1)[j].GetOutwardConnections()[k].GetNodeB()))+0.5)*(h/float64(len(g.GetNodesWithLayer(g.GetNodesWithLayer(i + 1)[j].GetOutwardConnections()[k].GetNodeB().GetLayer()))))))
				imd.Line(2)
			}
		}
	}
	imd.Draw(win)
}

func checkForCollisions(bird Bird, pipes []*Pipe) bool {
	if bird.height <= 0 {
		return true
	}
	for i := range pipes {
		if bird.Bounds().IntersectRect(pipes[i].Bounds()).X != 0 || bird.Bounds().IntersectRect(pipes[i].Bounds()).Y != 0 {
			return true
		}
	}
	return false
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, err := png.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func GetScore() int {
	return score
}
