package maze

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
)

type Position struct {
	X int
	Y int
}

type Route []Position

func (s *Route) Push(p Position) {
	*s = append(*s, p)
}

func (s *Route) Pop() Position {
	// TODO len check
	l := len(*s)
	p := (*s)[l-1]
	*s = (*s)[:l-1]

	return p
}

func (s *Route) IsEmpty() bool {
	if len(*s) > 0 {
		return false
	}
	return true
}

// cell is an individual part of a maze with a route north, west, and a visited flag
// a "route" north, representing lack of a wall, is used instead of a wall=true as
// the default type for a bool is false so by default there is no route from that cell.
// This saves constructing cells with true values.
type Cell struct {
	NorthRoute bool `json:"northRoute"`
	WestRoute  bool `json:"westRoute"`
	Visited    bool `json:"-"`
}

// maze represents the entire maze
type Maze struct {
	Cells [][]Cell `json:"cells"`
}

// New creates a new Maze problem
func New(width, height int, start Position) (Maze, error) {
	route := make(Route, 0)

	var m Maze
	m.Cells = make([][]Cell, width)

	for i := range m.Cells {
		m.Cells[i] = make([]Cell, height)
	}

	m.generate(start, &route)

	return m, nil
}

// Load creates a Maze from a JSON string
func Load(s string) (Maze, error) {
	var m Maze
	err := json.Unmarshal([]byte(s), &m)
	if err != nil {
		return Maze{}, fmt.Errorf("load error: %s", err)
	}

	return m, nil
}

// JSON returns a representation of the maze in JSON
func (m *Maze) JSON() (string, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return "", errors.New("could not marshal as json")
	}
	return string(b), nil
}

// generate a new maze problem if the maze is new
func (m *Maze) generate(p Position, r *Route) {
	r.Push(p)
	c := &m.Cells[p.X][p.Y]
	c.Visited = true

	neighbours := m.unvisitedFrom(p)

	if len(neighbours) > 0 {
		next := neighbours[rand.Intn(len(neighbours))]
		m.createRoute(p, next)
		m.generate(next, r)
	} else {
		if !r.IsEmpty() {
			prev, err := m.backtrackToRoute(r)
			if err == nil {
				m.generate(prev, r)
			}
		}
	}
}

// backtrackToRoute winds back down the stack to find another
// cell which has possible routes.
func (m *Maze) backtrackToRoute(r *Route) (Position, error) {
	if r.IsEmpty() {
		return Position{}, errors.New("no route")
	}

	var p Position

	for !r.IsEmpty() {
		p = r.Pop()
		u := m.unvisitedFrom(p)

		if len(u) > 0 {
			return p, nil
		}
	}

	return Position{}, errors.New("no unvisited cells")
}

// createRoute works out which cell should have its route
// set to true
func (m *Maze) createRoute(pos, next Position) {

	if pos.Y > next.Y {
		m.Cells[pos.X][pos.Y].NorthRoute = true
	}

	if pos.Y < next.Y {
		m.Cells[next.X][next.Y].NorthRoute = true
	}

	if pos.X < next.X {
		m.Cells[next.X][next.Y].WestRoute = true
	}

	if pos.X > next.X {
		m.Cells[pos.X][pos.Y].WestRoute = true
	}

}

// inBounds checks if an x,y coordinate is within the maze
func (m *Maze) inBounds(x, y int) bool {

	if x < 0 || x > len(m.Cells)-1 {
		return false
	}

	if y < 0 || y > len(m.Cells[0])-1 {
		return false
	}

	return true
}

// unvisitedFrom returns a list of pointers to unvisited
// neighbour cells from Position p
func (m *Maze) unvisitedFrom(p Position) []Position {
	var unvisited []Position

	// East
	x := p.X + 1
	y := p.Y
	if m.inBounds(x, y) {
		if !m.Cells[x][y].Visited {
			unvisited = append(unvisited, Position{X: x, Y: y})
		}
	}

	// West
	x = p.X - 1
	y = p.Y
	if m.inBounds(x, y) {
		if !m.Cells[x][y].Visited {
			unvisited = append(unvisited, Position{X: x, Y: y})
		}
	}

	// South
	x = p.X
	y = p.Y + 1
	if m.inBounds(x, y) {
		if !m.Cells[x][y].Visited {
			unvisited = append(unvisited, Position{X: x, Y: y})
		}
	}

	// North
	x = p.X
	y = p.Y - 1
	if m.inBounds(x, y) {
		if !m.Cells[x][y].Visited {
			unvisited = append(unvisited, Position{X: x, Y: y})
		}
	}

	return unvisited
}
