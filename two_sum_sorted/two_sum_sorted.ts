function twoSumSorted(numbers: number[], target: number): number[] {
  let left = 0;
  let right = numbers.length - 1;

  while (left < right) {
    const leftVal = numbers[left];
    const rightVal = numbers[right];

    const result = leftVal + rightVal;
    if (result === target) return [left + 1, right + 1];

    if (result > target) {
      right--;
    } else {
      left++;
    }
  }

  return [];
}
