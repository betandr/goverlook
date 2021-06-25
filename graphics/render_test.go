package graphics

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/betandr/goverlook/maze"
)

func TestImageRenderedGreaterThanZeroBytes(t *testing.T) {
	start := maze.Position{X: 0, Y: 0}
	m, err := maze.New(10, 10, start)
	if err != nil {
		t.Errorf(err.Error())
	}

	var b bytes.Buffer
	buf := bufio.NewWriter(&b)

	Render(buf, &m, start)

	if buf.Size() <= 0 {
		t.Errorf("image not rendered: contains %d bytes", buf.Size())
	}
}
