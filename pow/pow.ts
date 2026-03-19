function myPow(x: number, n: number): number {
  let acc: number = x;

  let newPow = Math.abs(n) - 1;
  const isNegative = n < 0;

  while (newPow !== 0) {
    acc *= x;
    newPow--;
  }

  if (isNegative) return 1 / acc;
  return acc;
}
