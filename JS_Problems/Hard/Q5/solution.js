function customSort(arr) {
  // Object to store frequency of each element
  const frequency = {};
  arr.forEach((num) => {
    frequency[num] = (frequency[num] || 0) + 1;
  });

  // Manual sorting (could use any sorting algorithm - here uses bubble sort)
  for (let i = 0; i < arr.length - 1; i++) {
    for (let j = 0; j < arr.length - i - 1; j++) {
      const shouldSwap =
        frequency[arr[j]] < frequency[arr[j + 1]] ||
        (frequency[arr[j]] === frequency[arr[j + 1]] && arr[j] > arr[j + 1]);
      if (shouldSwap) {
        // Swap elements
        [arr[j], arr[j + 1]] = [arr[j + 1], arr[j]];
      }
    }
  }

  return arr;
}
