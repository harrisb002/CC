function compressString(str) {
  // Return the original string if it's empty or has a length of 1
  if (str.length <= 1) {
    return str;
  }

  let compressed = "";
  let currentChar = str[0];
  let charCount = 1;

  for (let i = 1; i < str.length; i++) {
    if (str[i] === currentChar) {
      // Increment the count if the current character matches the previous one
      charCount++;
    } else {
      // Append the current character and its count (if more than 1) to the result string
      compressed += currentChar + (charCount > 1 ? charCount : "");
      currentChar = str[i];
      charCount = 1;
    }
  }

  // Append the last character and its count (if more than 1) to the result string
  compressed += currentChar + (charCount > 1 ? charCount : "");

  return compressed;
}
