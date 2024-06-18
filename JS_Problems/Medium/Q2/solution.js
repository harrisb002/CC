function mergeSortedArrays(arr1, arr2) {
    let mergedArray = [];
    let i = 0, j = 0; // Initialize pointers for both arrays

    // Continue looping until we reach the end of one of the arrays
    while (i < arr1.length && j < arr2.length) {
        if (arr1[i] < arr2[j]) {
            mergedArray.push(arr1[i]);
            i++; // Move pointer in arr1 forward
        } else {
            mergedArray.push(arr2[j]);
            j++; // Move pointer in arr2 forward
        }
    }

    // At this point, one of the arrays has been fully iterated through.
    // We now append the remaining elements from the other array.
    // Note: Only one of these while loops will execute, depending on which array has elements left.
    while (i < arr1.length) {
        mergedArray.push(arr1[i]);
        i++;
    }
    while (j < arr2.length) {
        mergedArray.push(arr2[j]);
        j++;
    }

    return mergedArray;
}