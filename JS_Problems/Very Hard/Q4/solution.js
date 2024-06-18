function findIslands(map) {
  let islandCount = 0;

  // Helper function for DFS traversal
  function dfs(x, y) {
    // Check for out-of-bounds or water or visited land
    if (
      x < 0 ||
      x >= map.length ||
      y < 0 ||
      y >= map[x].length ||
      map[x][y] === 0
    ) {
      return;
    }

    // Mark the land as visited by setting it to '0'
    map[x][y] = 0;

    // Explore all adjacent lands (up, down, left, right)
    dfs(x - 1, y); // up
    dfs(x + 1, y); // down
    dfs(x, y - 1); // left
    dfs(x, y + 1); // right
  }

  // Iterate over each cell in the map
  for (let i = 0; i < map.length; i++) {
    for (let j = 0; j < map[i].length; j++) {
      if (map[i][j] === 1) {
        // Found an unvisited land, start DFS from here
        dfs(i, j);
        // After returning from DFS, all lands connected to [i,j] have been visited
        // Increment the island count
        islandCount++;
      }
    }
  }

  return islandCount;
}
