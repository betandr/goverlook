package main

import (
	"flag"
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/betandr/goverlook/graphics"
	"github.com/betandr/goverlook/maze"
)

var width = flag.Int("width", 20, "width of the maze")
var height = flag.Int("height", 20, "height of the maze")

func main() {
	flag.Parse()
	generateMaze(os.Stdout)
}

func generateMaze(w io.Writer) {
	rand.Seed(time.Now().UnixNano())

	start := maze.Position{
		X: rand.Intn(*width - 1), // for zero-based index
		Y: rand.Intn(*height - 1),
	}

	m := maze.New(*width, *height, start)
	graphics.Render(w, &m, start)
}
