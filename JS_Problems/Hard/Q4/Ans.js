//////////////////////////////////////////////////////////////////////////////////////////
// QS: 4 Write a JavaScript function named canFormPalindrome that checks if it is possible to
// form a palindrome from the given input string by rearranging its letters. The function
// should return true if forming a palindrome is possible, and false otherwise. A palindrome
// is a word, phrase, number, or other sequences of characters that reads the same
// forward and backward (ignoring spaces, punctuation, and capitalization).
//////////////////////////////////////////////////////////////////////////////////////////

// function filter(str) {
//   let filteredStr = "";
//   for (let char of str) {
//     if (
//       str >= "A" ||
//       str <= "Z" ||
//       str >= "a" ||
//       str <= "z" ||
//       str >= "0" ||
//       str <= "9"
//     ) {
//       filteredStr += str.toLowerCase();
//     }
//   }
//   return filteredStr;
// }

// function canFormPalindrome(str) {
//     const str = str.replace(/[^a-zA-Z0-9]/g,''.toLowerCase());

//     let charCount = {}

//     for(let char of str) {
//         if(charCount[char]) {
//             charCount[char]++;
//         } else {
//             charCount[char] = 1;
//         }
//     }

//     let oddCount = 0;

//     for(let char of charCount) {
//         if(charCount[char] % 2 !== 0) {
//             oddCount++;
//         }
//     }
//     if(oddCount > 1) {
//         return false;
//     }
//     return true;
// }
// Sample Input and Output:
// Input: "carrace"
// Output: true (Because "carrace" can be rearranged to "racecar", which is a palindrome.)
// Input: "hello"
// Output: false (Because there's no way to rearrange "hello" to form a palindrome.)

//////////////////////////////////////////////////////////////////////////////////////////
// QS: Write a JavaScript function named customSort that sorts an array of numbers based on
// the frequency of the numbers, from highest frequency to lowest. If two numbers have
// the same frequency, the number which comes first in ascending order should come first.
// The function should not use the built-in .sort() method for its primary sorting logic.
//////////////////////////////////////////////////////////////////////////////////////////

function customSort(arr) {
  //store frequencys
  const frequency = {};

  arr.forEach((num) => {
    frequency[num] = (frequency[num] || 0) + 1;
  });

  // bubble sort)
  for (let i = 0; i < arr.length - 1; i++) {
    for (let j = 0; j < arr.length - i - 1; j++) {
      const shouldSwap =
        frequency[arr[j]] < frequency[arr[j + 1]] ||
        (frequency[arr[j]] === frequency[arr[j + 1]] && arr[j] > arr[j + 1]);
      if (shouldSwap) {
        // Swap 
        [arr[j], arr[j + 1]] = [arr[j + 1], arr[j]];
      }
    }
  }

  return arr;
}

console.log(customSort([4, 3, 1, 6, 4, 3]));

