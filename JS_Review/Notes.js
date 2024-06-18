// function randomNumber(max, min) {
//   return Math.floor(Math.random() * (max - min) + min);
// }

// let randomNum = randomNumber(1, 5);
// let numGuesses = 0;

// function reset() {
//   num = randomNum(1, 5);
//   numGuesses = 0;
//   document.getElementById("num-guesses").innerHTML = numGuesses;
// }

// function makeGuess() {
//   const guessInput = document.getElementById("guess");
//   const guess = Number(guessInput.value);
//   guessInput.value = "";
//   const feedback = document.getElementById("feedback");

//   numGuesses++;
//   document.getElementById("num-guesses").innerHTML = 0;

//   if (guess === randomNum) {
//     feedback.innerHTML = "You Guessed It!!!";
//     reset();
//   } else if (guess > randomNum) {
//     feedback.innerHTML = "That's not it, Try Lower!";
//   } else {
//     feedback.innerHTML = "That's not it, Try Higher!";
//   }
// }

//////////////////////////////////////////////////////////////////
/////////////////////////// OOP NOTES: ///////////////////////////
//////////////////////////////////////////////////////////////////

// Prototype Chain
// Look in obj, then look in base prototype until null
// obj => __proto__ => Object.prototype => __proto__ => null

const Person = {
  name: "DEFAULT NAME!!!",
  sayHi: () => {
    console.log("Hey there!");
  },
  // Override toString()
  toString() {
    return this.name;
  },
};

// Will use obj method, if doesnt exist, then will look at prototype for it's method
console.log(Person.toString());

// Every object in JS has a Prototype
// This is another object, linked to the object that has base funcitonality that can be inherited on the object
// obj is just added onto the base prototype

console.log("The base Proto object:", Object.getPrototypeOf(Person));
// OR
console.log("The base Proto object:", Person.__proto__);

// Start by looking in base object
// If doesnt exist, then look in the prototype chain to look for method/property

// Init with a specified prototype
const Ben = Object.create(Person);
Ben.name = "Ben"
const Sarah = Object.create(Person);
Sarah.name = "Sarah"

Ben.sayHi();
//
console.log("Ben has a name but no functions!!!, using prototype", Ben);
console.log("Prototype of new obj", Ben.__proto__);


// function minSubArr(arr) {
//   // Dynamic Solution
//   let min = Infinity;
//   let lastSum = 0;

//   for (let i = 0; i < arr.length; i++) {
//     // If adding the last sum is less than the current elements value
//     if (arr[i] + lastSum < arr[i]) {
//       // Then it should add last sum for the current last sum
//       lastSum += arr[i];
//       if (min > lastSum) {
//         min = lastSum;
//       }
//     } else {
//       // If curr idx is better off alone, then set it as such
//       lastSum = arr[i];
//     }
//   }
//   return min;
// }

// const res = minSubArr([20, -7, -3, 9, -4, 6, -9, 10]);
// console.log(res);

// [20, -7, -3, 9, -4, 6, -9, 10]
// 20, -7, -10, -1, -5, 1, -9, 1


function minSubArr(arr) {
  let min = Infinity;
  let lastSum = 0;

  for (let i = 0; i < arr.length; i++) {
    // Always consider the smaller of the current element or the sum of the current element and the last subarray sum
    lastSum = Math.min(arr[i], arr[i] + lastSum);
    
    // Update the minimum subarray sum
    min = Math.min(min, lastSum);
  }
  
  return min;
}

const res = minSubArr([20, -7, -3, 9, -4, 6, -9, 10]);
console.log(res);