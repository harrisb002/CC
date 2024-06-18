// Helper function to reverse a portion of the array
function reverse(arr, start, end) {
  while (start < end) {
    let temp = arr[start];
    arr[start] = arr[end];
    arr[end] = temp;
    start++;
    end--;
  }
}

function rotateArray(arr, k) {
  // Ensure k is within the bounds of the array length
  k %= arr.length;

  // Step 1: Reverse the entire array
  reverse(arr, 0, arr.length - 1);

  // Step 2: Reverse the first k elements
  reverse(arr, 0, k - 1);

  // Step 3: Reverse the remaining elements
  reverse(arr, k, arr.length - 1);

  return arr
}