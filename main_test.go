package main

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/betandr/goverlook/graphics"
	"github.com/betandr/goverlook/maze"
)

func TestMazeStartAtMaximum(t *testing.T) {
	width := 20
	height := 20

	startX := width - 1
	startY := height - 1

	start := maze.Position{
		X: startX,
		Y: startY,
	}

	m := maze.New(width, height, start)

	cell := m.Cells[startX][startX]

	if !cell.Visited {
		t.Errorf("cell not marked visited: cell [%d, %d] w/h: %d/%d", startX, startY, width, height)
	}
}

func TestImageRenderedGreaterThanZeroBytes(t *testing.T) {
	start := maze.Position{X: 0, Y: 0}
	m := maze.New(10, 10, start)

	var b bytes.Buffer
	buf := bufio.NewWriter(&b)

	graphics.Render(buf, &m, start)

	if buf.Size() <= 0 {
		t.Errorf("image not rendered: contains %d bytes", buf.Size())
	}
}

func TestGenerateMaze(t *testing.T) {
	var b bytes.Buffer
	buf := bufio.NewWriter(&b)

	generateMaze(buf)

	if buf.Size() <= 0 {
		t.Errorf("maze not rendered: contains %d bytes", buf.Size())
	}
}

func TestGenerateMazeWithMaxStartPosition(t *testing.T) {

	width := 10
	height := 10
	start := maze.Position{
		X: width - 1,
		Y: height - 1,
	}

	m := maze.New(10, 10, start)

	mX := len(m.Cells[0])
	mY := len(m.Cells)
	if len(m.Cells) != 10 && len(m.Cells[0]) != 10 {
		t.Errorf("%d by %d maze not generated correctly: is %d by %d ", width, height, mX, mY)
	}
}
