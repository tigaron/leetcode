function twoSum(nums: number[], target: number): number[] {
  const numMap = new Map<number, number>();
  for (let i = 0; i < nums.length; i++) {
    const diff = target - nums[i];
    if (numMap.has(diff)) {
      return [numMap.get(diff)!, i];
    }
    numMap.set(nums[i], i);
  }

  return [];
}
