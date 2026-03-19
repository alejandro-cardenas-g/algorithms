function isValid(s: string): boolean {
  const latests = [];

  const obj = {
    "}": "{",
    ")": "(",
    "]": "[",
  };

  for (const char of s) {
    const isOpenning = !obj[char as keyof typeof obj];
    if (isOpenning) {
      latests.push(char);
    } else {
      // closing before opening -> not valid
      if (!latests.length) return false;

      const lastOne = latests.pop();

      // is closing is not the same as the latest open -> not valid
      if (obj[char as keyof typeof obj] !== lastOne) {
        return false;
      }
    }
  }

  return true;
}
