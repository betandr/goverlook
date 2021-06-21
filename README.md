# Maze Generator
_Named after the Overlook hotel in The Shining, Goverlook generates a maze_

![Maze](/images/maze.png)

## Running
```
go run . -width=n -height=n > output.png
```
...where `n` are integers

## Building and Running Binary

### With Go installed [https://golang.org/doc/install]:
```
go build -o maze .
./maze -width=n -height=n > output.png
```

## Running Tests
```
go test .
```
