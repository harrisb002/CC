function fibIter(n) {
  if (n <= 2) return 1;

  let last = 1;
  let secondLast = 1;

  for (let i = 2; i < n; i++) {
    const temp = last;
    last = last + secondLast;
    secondLast = last;
  }
  return last;
}

const sum = fibIter(70);
console.log("Sum is: ", sum);

// const str = "Ben is awesome!";

// function recReverseStr(str, newStr = []) {
//   // Base Case
//   if (str.length == 0) return;

//   // Get the first letter in the string
//   const char = str[0];
//   // Recursively solve for the reverse of the rest of the string
//   recReverseStr(str.slice(1), newStr);
//   newStr.push(char);

//   console.log(str, newStr);

//   if (str.length === newStr.length) return newStr.join("");
// }

// const ans = recReverseStr(str);

// console.log(ans);
