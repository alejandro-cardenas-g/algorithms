const input = [
  [1, 2],
  [3, 4],
];

const fn = (input: number[][], x: number): number => {
  const targetList = [];

  for (let i = 0; i < input.length; i++) {
    for (let j = 0; j < input[0].length; j++) {
      targetList.push(input[i][j]);
    }
  }

  //mxn

  let bestResult = Infinity;
  for (const target of targetList) {
    let currentResult = 0;

    for (let i = 0; i < input.length; i++) {
      for (let j = 0; j < input[0].length; j++) {
        const current = input[i][j];

        const operationsAtSumming = (target - current) / x;
        const operationsAtSubstracting = (current - target) / x;

        if (operationsAtSumming === 0) continue;

        if (
          operationsAtSumming > 0 &&
          [0, 1].includes(operationsAtSumming % 2)
        ) {
          currentResult += operationsAtSumming;
          continue;
        }

        if (
          operationsAtSubstracting > 0 &&
          [0, 1].includes(operationsAtSubstracting % 2)
        ) {
          currentResult += operationsAtSubstracting;
          continue;
        }

        return -1;
      }
    }
    bestResult = Math.min(bestResult, currentResult);
  }

  return bestResult;
};

const fn2 = (input: number[][], x: number): number => {
  let targetList = [];

  // O(M x N)
  for (let i = 0; i < input.length; i++) {
    for (let j = 0; j < input[0].length; j++) {
      targetList.push(input[i][j]);
    }
  }

  const checkValue = targetList[0] % x;

  // O(MxN)
  for (const val of targetList) {
    if (val % x != checkValue) return -1;
  }

  // MxN O(log MxN)
  targetList = targetList.sort((a, b) => a - b);

  const medium = targetList[Math.floor(targetList.length / 2)];

  let result = 0;

  // O (MxN)

  for (const val of targetList) {
    result += Math.abs(val - medium) / x;
  }

  return result;
};

console.log(fn(input, 1));
console.log(fn(input, 1));

// current + x*n = target; => n = (target-current)/x
// current - x*n = target; => n = (current-target)/x
