package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/betandr/goverlook/graphics"
	"github.com/betandr/goverlook/maze"
)

var width = flag.Int("width", 20, "width of the maze")
var height = flag.Int("height", 20, "height of the maze")
var out = flag.String("out", "png", "The output format (png/json)")

func main() {
	flag.Parse()
	stat, _ := os.Stdin.Stat()
	var mz maze.Maze
	var err error
	var start maze.Position

	if (stat.Mode() & os.ModeCharDevice) == 0 {
		mz, err = loadMaze(os.Stdin)
	} else {
		mz, start, err = generateMaze()
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	switch *out {
	case "png":
		graphics.Render(os.Stdout, &mz, start)
	case "json":
		s, err := mz.JSON()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(s)
	}
}

func generateMaze() (maze.Maze, maze.Position, error) {
	rand.Seed(time.Now().UnixNano())

	start := maze.Position{
		X: rand.Intn(*width - 1), // for zero-based index
		Y: rand.Intn(*height - 1),
	}

	m, err := maze.New(*width, *height, start)
	if err != nil {
		msg := fmt.Sprintf("error: creating %dx%d maze with start at [%d, %d]", *width, *height, start.X, start.Y)
		return maze.Maze{}, maze.Position{}, errors.New(msg)
	}

	return m, start, nil
}

func loadMaze(r io.Reader) (maze.Maze, error) {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return maze.Maze{}, errors.New("could not read from stdin")
	}

	m, err := maze.Load(text)
	if err != nil {
		return maze.Maze{}, fmt.Errorf("could not load maze from stdin: %s", err)
	}

	return m, nil
}
