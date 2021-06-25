package maze

import (
	"testing"
)

func TestMazeStartAtMaximum(t *testing.T) {
	width := 20
	height := 20

	startX := width - 1
	startY := height - 1

	start := Position{
		X: startX,
		Y: startY,
	}

	m, err := New(width, height, start)
	if err != nil {
		t.Errorf(err.Error())
	}

	cell := m.Cells[startX][startX]

	if !cell.Visited {
		t.Errorf("cell not marked visited: cell [%d, %d] w/h: %d/%d", startX, startY, width, height)
	}
}

func TestGenerateMazeWithMaxStartPosition(t *testing.T) {

	width := 10
	height := 10
	start := Position{
		X: width - 1,
		Y: height - 1,
	}

	m, err := New(10, 10, start)
	if err != nil {
		t.Errorf(err.Error())
	}

	mX := len(m.Cells[0])
	mY := len(m.Cells)
	if len(m.Cells) != 10 && len(m.Cells[0]) != 10 {
		t.Errorf("%d by %d maze not generated correctly: is %d by %d ", width, height, mX, mY)
	}
}
