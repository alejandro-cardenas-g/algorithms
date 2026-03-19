const threeSum = (nums: number[]): number[][] => {
  if (nums.length < 3) return [];

  nums.sort();

  const set = new Set<string>();

  const results: number[][] = [];

  for (let i = 0; i < nums.length - 2; i++) {
    let left = i + 1;
    let right = nums.length - 1;

    const expectedDifference = 0 - nums[i];

    while (left < right) {
      const result = nums[left] + nums[right];
      if (result === expectedDifference) {
        const key = `${nums[i]}${nums[left]}${nums[right]}`;
        if (!set.has(key)) {
          results.push([nums[i], nums[left], nums[right]]);
          set.add(key);
        }
        right--;
        continue;
      }

      if (result > expectedDifference) right--;
      if (result < expectedDifference) left++;
    }
  }

  return results;
};
