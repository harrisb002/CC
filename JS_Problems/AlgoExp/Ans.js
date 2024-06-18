//////////////////////////////////////////////////////
// Question StaircaseTranversal
// You're given two positive integers representing the height of a staircase
// and the maximum number of steps that you can advance up the staircase at a time.
// Write a function that returns the number of ways in which you can climb the staircase.
// For example, if you were given a staircase of height = 3 and maxSteps = 2 you could climb the staircase in 3 ways.
// You could take 1 step, 1 step, then 1 step, you could also take 1 step, then 2 steps, and you could take 2 steps, then 1 step.
// Note that maxSteps <= height will always be true.
//////////////////////////////////////////////////////

//////////// SPACE & TIME COMP. ///////////////
// RECnumOfWaysToTop: O(k^n) time | O(n) space
// MEMnumWaysToTop: O(k*n) time | O(n) space // Cache the answers
// DYNnumWaysToTop: O(k*n) time | O(n) space
// SlidingWindow: O(n) time | O(n) space
///////////////////////////////////////////////
baseCases = { 0: 1, 1: 1 };

function staircaseTraversal(height, maxSteps, baseCases = undefined) {
  //   return RECnumOfWaysToTop(height, maxSteps);
  //   return MEMnumWaysToTop(height, maxSteps, baseCases); // Include the base Cases
  //   return DYNnumWaysToTop(height, maxSteps); // Dynamic Solution
  return SlidingWindow(height, maxSteps);
}

function staircaseTraversal(height, maxSteps) {
  let currentNumberOfWays = 0;
  let waysToTop = [1];

  for (let currentHeight = 1; currentHeight <= height; currentHeight++) {
    let startOfWindow = currentHeight - maxSteps - 1;
    let endOfWindow = currentHeight - 1;

    if (startOfWindow >= 0) {
      currentNumberOfWays -= waysToTop[startOfWindow];
    }
    currentNumberOfWays += waysToTop[endOfWindow];
    waysToTop.push(currentNumberOfWays);
  }

  return waysToTop[height];
}

console.log("SLIDING WINDOW ANS: ", staircaseTraversal(4, 3));

// function DYNnumWaysToTop(height, maxSteps) {
//   // Used to store the answers previously calc.
//   // The indicies in this DS represent: # of ways to get to the top for a staircase for that indicies height
//   let waysToTop = Array(height + 1).fill(0);

//   // Base cases
//   waysToTop[0] = 1;
//   waysToTop[1] = 1;

//   for (let currHeight = 2; currHeight < height + 1; currHeight++) {
//     let step = 1;

//     // Look at the previous steps, sum the previous ways to get to those steps
//     // Sum up the K work, K is the number of steps on could possibly take.
//     while (step <= maxSteps && step <= currHeight) {
//       // Look backwards in the Data Structure to add to current result
//       waysToTop[currHeight] =
//         waysToTop[currHeight] + waysToTop[currHeight - step];
//       step += 1;
//     }
//   }

//   return waysToTop[height];

//   console.log(waysToTop);
// }

// //Cache/Keep track of prior calls using Memoization
// function MEMnumWaysToTop(height, maxSteps, memoize) {
//   // Now find if height has been already calulated and if so use it
//   if (height in memoize) {
//     return memoize[height];
//   }
//   numOfWays = 0;
//   for (let step = 0; step < Math.min(maxSteps, height) + 1; step++) {
//     numOfWays += MEMnumWaysToTop(height - step, maxSteps, memoize);
//   }
//   // Store the solution in the memoize data Structure
//   memoize[height] = numberOfWays;
//   return numOfWays;
// }

// function RECnumOfWaysToTop(height, maxSteps) {
//   // Base Case
//   if (height <= 1) {
//     return 1;
//   }
//   numOfWays = 0;

//   // Try all the steps, (1, 2, ... ect) up to the maximum height and add to the numOfWays
//   for (let step = 1; step < Math.min(maxSteps, height) + 1; step++) {
//     // Subtract the step from the current height
//     numOfWays += numOfWaysToTop(height - step, maxSteps);
//   }

//   return numOfWays;
// }

//////////////////////////////////////////////////////
// Question Reverse Words In String
// Write a function that takes in a string of words separated by one or more whitespaces and
// returns a string that has these words in reverse order. For example, given the string "tim is great", your function should return "great is tim".
// For this problem, a word can contain special characters, punctuation, and numbers.
// The words in the string will be separated by one or more whitespaces,
// and the reversed string must contain the same whitespaces as the original string.
// For example, given the string "whitespaces 4" you would be expected to return "4 whitespaces".
// Note that you're not allowed to use any built-in split or reverse methods/functions.
// However, you are allowed to use a built-in join method/function.
// Also note that the input string isn't guaranteed to always contain words.
//////////////////////////////////////////////////////

// Sample Input
// string = "Chess is the best!" // "best! the is Chess"

//////////////////////////////////////////////////////
// Question
// You're given a two-dimensional array (a matrix) of potentially unequal height and width containing only 0s and 1s.
// The matrix represents a two-toned image, where each 1 represents black and each 0 represents white.
// An island is defined as any number of 1s that are horizontally or vertically adjacent (but not diagonally adjacent)
// and that don't touch the border of the image. In other words, a group of horizontally or vertically adjacent 1s isn't an island
// if any of those 1s are in the first row, last row, first column, or last column of the input matrix.
// Note that an island can twist. In other words, it doesn't have to be a straight vertical line or a straight horizontal line; it can be L-shaped, for example.
// You can think of islands as patches of black that don't touch the border of the two-toned image.
// Write a function that returns a modified version of the input matrix, where all of the islands are removed. You remove an island by replacing it with 0s.
// Naturally, you're allowed to mutate the input matrix.
//////////////////////////////////////////////////////

// Sample Input
// matrix =
// [
//   [1, 0, 0, 0, 0, 0],
//   [0, 1, 0, 1, 1, 1],
//   [0, 0, 1, 0, 1, 0],
//   [1, 1, 0, 0, 1, 0],
//   [1, 0, 1, 1, 0, 0],
//   [1, 0, 0, 0, 0, 1],
// ]

// Sample Output
// [
//     [1, 0, 0, 0, 0, 0],
//     [0, 0, 0, 1, 1, 1],
//     [0, 0, 0, 0, 1, 0],
//     [1, 1, 0, 0, 1, 0],
//     [1, 0, 0, 0, 0, 0],
//     [1, 0, 0, 0, 0, 1],
//   ]
// The islands that were removed can be clearly seen here:
// [
//   [   ,   ,   ,   ,   ,   ],
//   [   , 1,   ,   ,   ,   ],
//   [   ,   , 1,   ,   ,   ],
//   [   ,   ,   ,   ,   ,   ],
//   [   ,   , 1, 1,   ,   ],
//   [   ,   ,   ,   ,   ,   ],
// ]

//////////////////////////////////////////////////////
// Question Sum of Linked Lists
// You're given two Linked Lists of potentially unequal length.
// Each Linked List represents a non-negative integer, where each node in the Linked List
// is a digit of that integer, and the first node in each Linked List always represents
// the least significant digit of the integer. Write a function that returns the head of a
// new Linked List that represents the sum of the integers represented by the two input Linked Lists.
// Each LinkedList node has an integer value as well as a next node pointing to the next node in the list
// or to None / null if it's the tail of the list. The value of each LinkedList node is always in the range of 0 - 9.
// Note: your function must create and return a new Linked List, and you're not allowed to modify either of the input Linked Lists.
//////////////////////////////////////////////////////

// linkedListOne = 2 -> 4 -> 7 -> 1
// linkedListTwo = 9 -> 4 -> 5

// 1 -> 9 -> 2 -> 2
// linkedListOne represents 1742
// linkedListTwo represents 549
// 1742 + 549 = 2291
