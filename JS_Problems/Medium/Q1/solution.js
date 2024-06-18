function findUniqueCharacters(str) {
    const charCount = {}; // Object to hold character counts
    let uniqueChars = ''; // String to hold unique characters

    // First pass: Count occurrences of each character
    for (const char of str) {
        if (charCount[char] === undefined) {
            charCount[char] = 1; // Initialize with 1 on first occurrence
        } else {
            charCount[char] += 1; // Increment count on subsequent occurrences
        }
    }

    // Second pass: Build string of unique characters
    for (const char of str) {
        if (charCount[char] === 1) {
            uniqueChars += char; // Append character if it occurred exactly once
        }
    }

    return uniqueChars;
}
