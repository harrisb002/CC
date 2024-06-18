function canFormPalindrome(str) {
  const charCounts = {};

  // Count occurrences of each character
  for (const char of str) {
    if (charCounts[char]) {
      charCounts[char]++;
    } else {
      charCounts[char] = 1;
    }
  }

  let oddCount = 0;

  // Determine the number of characters with odd counts
  for (const char in charCounts) {
    if (charCounts[char] % 2 !== 0) {
      oddCount++;
    }
  }

  // For a string to be able to form a palindrome, there must be at most one character with an odd count
  return oddCount <= 1;
}