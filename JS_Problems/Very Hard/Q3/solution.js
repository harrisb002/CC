function mergeIntervals(intervals) {
  // Sort intervals by their start time
  intervals.sort((a, b) => a[0] - b[0]);

  const merged = [];
  for (let i = 0; i < intervals.length; i++) {
    const [start, end] = intervals[i];

    // If the list of merged intervals is empty or if the current interval
    // does not overlap with the previous, simply append it.
    if (!merged.length || merged[merged.length - 1][1] < start) {
      merged.push([start, end]);
    } else {
      // Otherwise, there is overlap, so we merge the current and previous intervals.
      merged[merged.length - 1][1] = Math.max(
        merged[merged.length - 1][1],
        end
      );
    }
  }

  return merged;
}

