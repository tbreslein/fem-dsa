// Package recursion foo
package recursion

import "github.com/tbreslein/fem-dsa/src/stack"

type Point struct {
	x, y int
}

func SolveMaze(maze []string, wall string, start Point, end Point) []Point {
  // NOTE: would be more ergonomic using a boolean mask for the coordinates, instead of this set
	var visitedTiles map[Point]struct{}

	path := stack.NewStack[Point]()
	step(start.x, start.y, &path, &maze, &visitedTiles, &wall, &end)

  // convert stack to array
  pathArray := make([]Point, 10)
  for path.List.PeekFront() != end {
    pathArray = append(pathArray, path.List.PopFront())
  }
  // since the for loop leaves the end point in the stack, append that outside of the loop
  pathArray = append(pathArray, path.List.PopFront())
	return pathArray
}

func step(j int, i int, path *stack.Stack[Point], maze *[]string, visitedTiles *map[Point]struct{}, wall *string, end *Point) bool {
	// base cases
  if i < 0 || j < 0 || i >= len((*maze)[0]) || j >= len(*maze) {
    return false
  }
  if string((*maze)[j][i]) == *wall {
    return false
  }
  if _, alreadyVisisted := (*visitedTiles)[Point{j,i}]; alreadyVisisted {
    return false
  }

  // new tile, so add it to the stack
  path.Push(Point{j,i})

	// check if we reached the end
	if i == end.x && j == end.y {
		return true
	}

  return step(j+1, i, path, maze, visitedTiles, wall, end) ||
    step(j, i+1, path, maze, visitedTiles, wall, end) ||
    step(j-1, i, path, maze, visitedTiles, wall, end) ||
    step(j, i-1, path, maze, visitedTiles, wall, end)
}
