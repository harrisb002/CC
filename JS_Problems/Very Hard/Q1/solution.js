function findUniqueIndexPairs(numbers, targetSum) {
  let pairs = [];
  let usedIndices = new Set();

  // Iterate through the array to consider all possible pairs
  for (let i = 0; i < numbers.length; i++) {
    for (let j = i + 1; j < numbers.length; j++) {
      // Check if the current pair sums up to the target and neither index has been used
      if (
        numbers[i] + numbers[j] === targetSum &&
        !usedIndices.has(i) &&
        !usedIndices.has(j)
      ) {
        pairs.push([i, j]); // Add the pair of indices
        usedIndices.add(i).add(j); // Mark indices as used
        break; // Ensure we don't reuse the first index in another pair
      }
    }
  }

  return pairs;
}
