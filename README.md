# Maze Generator
_Named after the Overlook hotel in The Shining, Goverlook generates a maze_

![Maze](/images/maze.png)

## Building

### With Go installed [https://golang.org/doc/install]:
```
go build -o goverlook .
```

## Running
_(with `goverlook` in executable path)_

Default will create a new 20x20 maze and output as a PNG on stdout:
```
goverlook > out.png
```

Maze with specific size:
```
goverlook -width=60 -height=40 > out.png
```

Maze rendering as PNG (default):
```
goverlook -out=png > out.png
```

Maze rendering as JSON:
```
goverlook -out=json > out.json
```

Rendering a JSON maze from stdin:
```
cat out.json | goverlook > out.png
```

## Testing
```
go test .
```
