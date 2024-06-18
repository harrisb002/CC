function isPalindrome(str) {
    // Manually filter out non-alphanumeric characters and convert to lowercase
    let filteredStr = '';
    for (let char of str) {
        if ((char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9')) {
            filteredStr += char.toLowerCase();
        }
    }

    // Initialize pointers for the start and end of the filtered string
    let start = 0;
    let end = filteredStr.length - 1;

    // Iterate towards the middle of the string
    while (start < end) {
        if (filteredStr[start] !== filteredStr[end]) {
            return false; // If characters at start and end don't match, it's not a palindrome
        }
        start++;
        end--;
    }

    return true; // The string is a palindrome
}

console.log(isPalindrome("A man, a plan, a canal, Panama"))