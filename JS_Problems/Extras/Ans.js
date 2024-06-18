//////////////////////////////////////////////////////
// Question 1 - Easy
// Write a simple program that uses both HTML and JavaScript (you can write your
// JavaScript code directly in your HTML file or by referencing a script).
// This program should prompt the user to enter a number and
// tell them if their number is odd or even.

// Example:
// ~Program displays a prompt~
// ~User enters 3~
// ~Program shows an alert saying: “Odd”~
// You can use the functions “prompt” and “alert” in your javascript code.
//////////////////////////////////////////////////////

// let resp = prompt("Please enter a number to determine if it's odd or even: ");
// alert(`${resp % 2 === 0 ? "Even" : "Odd"}`);

//////////////////////////////////////////////////////
// Question 2 - Easy
// Write a simple program that uses both HTML and
// JavaScript (you can write your JavaScript code directly in your HTML file or by referencing a script).
// This program should ask the user to enter 3 numbers, it should store
// each of the numbers and tell the user what the largest number they entered was.

// Example:
// ~Program asks for number 1~ ~User enters 1~
// ~Program asks for number 2~ ~User enters -3~
// ~Program asks for number 3~
// ~User enters -5~
// ~Program shows an alert saying “The largest number was: 1”~
//////////////////////////////////////////////////////

// let resp1 = prompt("Please enter first number: ");
// let resp2 = prompt("Please enter second number: ");
// let resp3 = prompt("Please enter third number: ");
// let max = Math.max(resp1, resp2);
// max = Math.max(max, resp3);
// alert(`Max is ${max}`);

//////////////////////////////////////////////////////
// Question 3 - Easy
// Write a simple program that only requires a JavaScript console.
// You can use node.js or the console from the chrome browser, it’s up to you.
// The program should declare two variables: min and max.
// These variables should store two positive numbers that you can change before running the code.
// Your code should print an array to the console that contains all of the values from min -> max (inclusive).
// You can define these variables as constants at the top of your program.
// The idea is that you can manually change them before you run the code to change what the output will be.
// //////////////////////////////////////////////////////
// const arr = [];
// let min = 2;
// const max = 7;
// while (min <= max) {
//   arr.push(min);
//   min++;
// }

// console.log(arr);

//////////////////////////////////////////////////////
// Question 4 - Medium
// Write a simple program that only requires a JavaScript console.
// You can use node.js or the console from the chrome browser, it’s up to you.

// For this program you define three variables: 1. Target (a string you will manipulate)
// 2. Delimiter (a single character)
// 3. Spacing (a positive integer greater than 0)
// You can define these variables as constants at the top of your program.
// The idea is that you can manually change them before you run the code to change what the output will be.
// Your goal will be to print a new version of the target string that has the
// delimiter character between the spacing number of characters. See the example below:
//////////////////////////////////////////////////////

// function q4(target, delimiter, spacing) {
//   let newStr = "";
//   for (let i = 0; i < target.length; i++) {
//     if (i % spacing === 0 && i !== 0) {
//       newStr += delimiter;
//     }
//     newStr += target[i];
//   }
//   return newStr;
// }

// const target = q4("benisgreat", "-", 3);
// const target2 = q4("coursecareers", "|", 1);

// console.log(target);
// console.log(target2);

// Example 1:
//
// ~Program prints: “tim-isg-rea-t”~ notice the - exists after 3 characters Example 2:
// Target = “coursecareers” Delimiter = “|”
// Spacing = “1”
// ~Program prints: “c|o|u|r|s|e|c|a|r|e|e|r|s”~

//////////////////////////////////////////////////////
// Question 5 - Medium
// Write a simple program that only requires a JavaScript console.
// You can use node.js or the console from the chrome browser, it’s up to you.

// For this question you will write a JavaScript function called: mathIsFun .
// This function will have one parameter called “numberString”.
// The string will contain a set of numbers, all separated by |’s.
// Your job will be to determine all of the numbers in the string and return the
// largest number that can be created by adding two of the numbers.
// You should return this number from the function.
//////////////////////////////////////////////////////

// function mathIsFun(numberString) {
//   const nums = numberString.split("|").map(num => parseInt(num, 10)).filter(num => !isNaN(num));
//   if (nums.length < 2) {
//     return null; // Not enough numbers to form a sum
//   }

//   let max = Number.MIN_SAFE_INTEGER;
//   let secMax = Number.MIN_SAFE_INTEGER;

//   for (const num of nums) {
//     if (num > max) {
//       secMax = max;
//       max = num;
//     } else if (num > secMax) {
//       secMax = num;
//     }
//   }

//   return max + secMax;
// }

// console.log(mathIsFun("12|13|-4|5|")); // Output: 25
// console.log(mathIsFun("-2|-4|-1|-1")); // Output: -2

// Example 1:
// mathIsFun(“12|13|-4|5|”) -> 25
// // the numbers are 12, 13, -4 and 5. The largest sum of two numbers is 12 + 13 = 25
// Example 2:
// mathIsFun(“-2|-4|-1|-1”) -> -2
// // the numbers are -2, -4, -1 and -1. The largest sum of two numbers is -1 + -1 = -2
// Hint: You will need to use a nested for loop to try every possible sum of numbers.

//////////////////////////////////////////////////////
// Question 6 - Hard
// Write a simple program that only requires a JavaScript console.
// You can use node.js or the console from the chrome browser, it’s up to you.
// For this question you will write a JavaScript function called: friends .
// This function will accept a parameter called “people” that contains a list of objects.
// These objects will have two properties, the first is “name” and the second is a list of strings called “friends”.

// Your job will be to determine which person has the most loyal friends.
// A loyal friend is one that a person has in their friends list that also has that person in their friends list.
// For example, if John has Susan in his friends list and Susan has John in her friends list then they are “loyal friends”.
// On the contrary, if John has Mike in his friends list but Mike does not have John in his friends list they are not loyal friends.
// Note: This is a difficult question!
//////////////////////////////////////////////////////

// function friends(people) {
//   let mvp = "";
//   let currMax = 0;

//   // Init the loyalty scores
//   people.forEach((person) => (person.loyaltyScore = 0));

//   for (const person of people) {
//     console.log("Current person: ", person);
//     // Loop through each friend in the persons list

//     for (let i = 0; i < person.friends.length; i++) {
//       console.log("Current friend of person: ", person.friends[i]);

//       // Get the friend from the list of people objects
//       const friendObj = people.find((p) => p.name === person.friends[i]);

//       console.log(friendObj);

//       if (friendObj) {
//         // Check if the person is listed as a friend is their list of friends
//         const isFriend = friendObj.friends.includes(person.name);

//         if (isFriend) {
//           // Increment the persons loyalty score that has them as a friend
//           person.loyaltyScore++;
//           console.log(
//             "Incremeneted Loyalty score of with:",
//             person.name,
//             person.loyaltyScore
//           );
//         }
//       }

//       // Update MVP if their loyaltyScore is now the greatest
//       if (person.loyaltyScore > currMax) {
//         console.log("Person name of new MVP:", person.name);
//         mvp = person.name;
//         currMax = person.loyaltyScore;
//         console.log("Current max: ", currMax);
//       }
//     }
//   }

//   return mvp;
// }

// people = [
//   { name: "Tim", friends: ["John", "Sally"] },
//   { name: "John", friends: ["Tim", "Mike"] },
//   { name: "Mike", friends: [] },
//   { name: "Sally", friends: ["Tim"] },
// ];

// console.log(friends(people)); // -> “Tim”

// Example:

// // In this example Tim has two loyal friends, John and Sally, because both of them also list Tim as a friend
// // John has one loyal friend, Tim, as Mike doesn’t list John
// // Mike has no loyal friends as he lists no one as a friend
// // Sally has one loyal friend, “Tim”