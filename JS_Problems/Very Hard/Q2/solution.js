function canTraverseMaze(maze, directions) {
  if (!directions) return false;

  let x = 0;
  let y = 0;

  // Start by locating the starting point
  for (let row = 0; row < maze.length; row++) {
    for (let col = 0; col < maze[row].length; col++) {
      if (maze[row][col] === "S") {
        x = col;
        y = row;
      }
    }
  }

  const isValidMove = (x, y) => {
    return (
      x >= 0 &&
      x < maze[0].length &&
      y >= 0 &&
      y < maze.length &&
      maze[y][x] !== "#"
    );
  };

  for (const direction of directions) {
    switch (direction) {
      case "N":
        y--;
        break;
      case "S":
        y++;
        break;
      case "E":
        x++;
        break;
      case "W":
        x--;
        break;
    }

    // Check if out of bounds or hit a wall
    if (!isValidMove(x, y)) {
      return false;
    }

    // Check if reached the finish
    if (maze[y][x] === "F") {
      return true;
    }
  }

  // Return false if finish wasn't reached by the end of directions
  return false;
}
