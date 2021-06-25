package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
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
	var loadMaze bool

	if (stat.Mode() & os.ModeCharDevice) == 0 {
		loadMaze = true
	}

	args := make(map[string]string)
	args["width"] = strconv.Itoa(*width)
	args["height"] = strconv.Itoa(*width)
	args["out"] = *out

	if err := run(args, os.Stdout, os.Stdin, loadMaze); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

// run executes the logic to generate the maze
func run(args map[string]string, stdout io.Writer, stdin io.Reader, load bool) error {
	var mz maze.Maze
	var err error
	var start maze.Position

	if load {
		mz, err = loadMaze(stdin)
	} else {
		mz, start, err = generateMaze()
	}

	if err != nil {
		return fmt.Errorf("error: %s", err)
	}

	// todo
	switch args["out"] {
	case "json":
		s := mz.JSON()
		stdout.Write([]byte(s))
	default:
		graphics.Render(stdout, &mz, start)
	}

	return nil
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

func loadMaze(reader io.Reader) (maze.Maze, error) {
	text, err := ioutil.ReadAll(reader)
	if err != nil {
		return maze.Maze{}, errors.New("could not read from stdin")
	}

	m, err := maze.Load(string(text))
	if err != nil {
		return maze.Maze{}, fmt.Errorf("could not load maze from stdin: %s", err)
	}

	return m, nil
}
