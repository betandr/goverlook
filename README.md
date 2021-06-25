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

Chaining `goverlook` into generation / render stages:
```
goverlook -out=json | goverlook > out.png
```

## JSON Format
Mazes are represented as a collection of collections of "cells". Each cell has a `northRoute` and `westRoute` boolean.

```
{
	"cells": [
		[{
			"northRoute": false,
			"westRoute": false
		}, {
			"northRoute": false,
			"westRoute": false
		}],
		[{
			"northRoute": false,
			"westRoute": true
		}, {
			"northRoute": true,
			"westRoute": true
		}]
	]
}
```

## Testing
```
go test ./...
```

### Testing With Coverage
```
go test -v -coverpkg=./... -coverprofile=profile.cov ./...
```

### Test Coverage
```
go tool cover -func profile.cov
```
