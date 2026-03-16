const test1 = "abaabbbcccc";

const strMax = (str: string): string => {
  let maxChar: string = "";
  let current: string = "";

  for (let i = 0; i < str.length; i++) {
    const currentChar = current[current.length - 1] || "";
    if (currentChar === str[i]) {
      current += str[i];
    } else {
      current = str[i];
    }

    if (current.length > maxChar.length) {
      maxChar = current;
    }
  }
  return maxChar;
};

console.log(strMax(test1));
