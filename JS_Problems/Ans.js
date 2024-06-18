//////////////////////////////////////////////////////////////////////////////////////////
// QS: Implement a JavaScript function named findUniqueIndexPairs that, given an array of
// integer numbers and a target sum targetSum, identifies all pairs of indices whose values
// sum up to the target sum. Each pair of indices should be unique and returned as an
// array of pairs, with each pair represented as a two-element array.

// The function must ensure that:
// ● The sum of the values at the two indices in each pair equals the target sum.
// ● Each index from the array is used at most once across all pairs.
// ● The pairs are unique; there are no duplicate pairs in the output.
// ● The order of pairs or indices in the output does not matter.

// Constraints:
// ● The array may contain both positive and negative integers.
// ● The function should correctly handle cases with duplicate values in the input
// array.
// ● If no pairs match the criteria, the function should return an empty array.
// ● Indices in a pair should be listed in ascending order (i.e., [i, j] where i < j).
//////////////////////////////////////////////////////////////////////////////////////////

// function findUniqueIndexPairs(numbers, targetSum) {
//   let pairs = [];
//   let usedIndices = new Set();

//   // Iterate through the array to consider all possible pairs
//   for (let i = 0; i < numbers.length; i++) {
//     for (let j = i + 1; j < numbers.length; j++) {
//       // Check if the current pair sums up to the target and neither index has been used
//       if (
//         numbers[i] + numbers[j] === targetSum &&
//         !usedIndices.has(i) &&
//         !usedIndices.has(j)
//       ) {
//         pairs.push([i, j]); // Add the pair of indices
//         usedIndices.add(i).add(j); // Mark indices as used
//         break; // Ensure we don't reuse the first index in another pair
//       }
//     }
//   }

//   return pairs;
// }

// const numbers = [2, 2, 3, 3, 4, 4];
// const targetSum = 6;

// console.log(findUniqueIndexPairs(numbers, targetSum));

// Sample Input 2
// const numbers = [1, 2, 3, 4, 5];
// const targetSum = 9;
// findUniqueIndexPairs(numbers, targetSum)
// Sample Output 2
// [
// [3, 4] // 4 + 5 = 9

// only one pair meets the criteria without reusing indices.

// ////////////////////////////////////////////////////////////////////////////////////////
// QS: You are given a 2D grid representing a maze, where:
// ● 'S' represents the starting point.
// ● 'F' represents the finish point.
// ● '.' represents an open path that can be traveled.
// ● '#' represents a wall that cannot be passed through.
// Write a JavaScript function named canTraverseMaze that takes two arguments: a 2D
// grid (array of strings) representing the maze, and a string of directions to attempt to
// traverse the maze. The directions are given as a string consisting of the letters 'N'
// (north), 'S' (south), 'E' (east), and 'W' (west). The function should return true if the
// directions lead from the start to the finish point without running into walls or going out of
// bounds, and false otherwise.
// ////////////////////////////////////////////////////////////////////////////////////////

// function canTraverseMaze(maze, directions) {
//   let currPos = maze[0][0];
//   let currRow = 0;
//   let currCol = 0;

//   for (let char of directions) {
//     console.log("Char is: ", char);
//     if (char === "N") {
//       currRow--;
//     } else if (char === "S") {
//       currRow++;
//     } else if (char === "E") {
//       currCol++;
//     } else {
//       // WEST
//       currCol--;
//     }

//     // bounds checking
//     if (
//       currCol >= maze[0].length ||
//       currRow >= maze.length ||
//       currCol < 0 ||
//       currRow < 0
//     ) {
//       return false;
//     }

//     // Update current position
//     currPos = maze[currRow][currCol];

//     if (maze[currRow][currCol] === "#") {
//       return false;
//     }

//     if (maze[currRow][currCol] === "F") {
//       return true;
//     }

//     currPos = maze[currRow][currCol]; // Update pos.
//     console.log("Current Position:", currPos);
//   }
// }
// const maze = ["S.#", "...", "..F"];

// const directions = "SEES"; // true
// console.log(canTraverseMaze(maze, directions));

// const maze2 = ["S.#", "...", "..F"];

// const directions2 = "SSS"; // false
// console.log(canTraverseMaze(maze2, directions2));

//////////////////////////////////////////////////////////////////////////////////////////
// QS: Write a JavaScript function named mergeIntervals that merges all overlapping intervals
// from a given collection of intervals. Each interval is represented as an array of two
// numbers [start, end], where start is the beginning of the interval and end is the end of
// the interval. The function should return an array of the merged intervals in no specific
// order. Intervals that do not overlap should be included in the output as they are.

// Constraints:
// ● The input is an array of intervals, where each interval is represented as an array
// of two integers [start, end].
// ● The start of an interval will always be less than or equal to the end.
// ● The output should be an array of the merged intervals.
// ● Intervals [start, end] and [end, nextStart] are considered contiguous and should
// be merged.
//////////////////////////////////////////////////////////////////////////////////////////
function mergeIntervals(intervals) {
  if (intervals.length === 0) return [];

  // Sort by starting indices
  intervals.sort((a, b) => a[0] - b[0]);

  let merged = [];
  let currInterval = intervals[0];

  for (let i = 1; i < intervals.length; i++) { // Start loop from index 1
    let [currStart, currEnd] = currInterval;
    let [nextStart, nextEnd] = intervals[i];
    if (currEnd >= nextStart) {
      // Combine Intervals
      currInterval = [currStart, Math.max(currEnd, nextEnd)];
    } else {
      // Add the current interval to merged and move to the next interval
      merged.push(currInterval);
      currInterval = intervals[i];
    }
  }
  
  // Add the last interval
  merged.push(currInterval);

  return merged;
}

const intervals = [
  [1, 3],
  [2, 6],
  [8, 10],
  [15, 18],
];
console.log(mergeIntervals(intervals)); // [[1, 6], [8, 10], [15, 18]]

const intervals2 = [
  [1, 4],
  [4, 5],
];
console.log(mergeIntervals(intervals2)); //[[1, 5]]

