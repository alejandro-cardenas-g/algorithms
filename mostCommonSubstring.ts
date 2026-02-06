const items = ["flowers", "flow", "floght", "fl"];

// Must be fl.

const fn = (r: string[]): string => {
  let max: number = 0;

  if (r.length === 1) return r[0];

  for (let i = 0; i < r.length - 1; i++) {
    const newMax = compareWithWhile(r[i], r[i + 1]);
    if (newMax === 0) {
      return "";
    }

    if (newMax < max) {
      max = newMax;
    }

    if (max === 0) {
      max = newMax;
    } else {
      max = Math.min(newMax, max);
    }
  }

  return r[0].substring(0, max);
};

const fn2 = (r: string[]): string => {
  if (r.length === 1) return r[0];

  let prefix = r[1];

  for (let i = 0; i < r.length; i++) {
    prefix = compareByChar(r[i], prefix);
  }

  return prefix;
};

function compare(r: string, r2: string): number {
  let current: number = 0;
  const isFirstLarger = r.length > r2.length;
  for (let i = 0; i < (isFirstLarger ? r2.length : r.length); i++) {
    if (r.substring(0, i + 1) === r2.substring(0, i + 1)) {
      current++;
    } else {
      break;
    }
  }
  return current;
}

function compareWithWhile(r: string, r2: string): number {
  let current: number = 0;

  while (current < r.length && current < r2.length) {
    if (r.substring(0, current + 1) === r2.substring(0, current + 1)) {
      current++;
    } else {
      break;
    }
  }

  return current;
}

function compareByChar(r: string, r2: string): string {
  let i = 0;
  let prefix: string = "";
  console.log(r, r2);
  while (i < r.length && i < r2.length) {
    if (r[i] === r2[i]) {
      prefix += r[i];
      i++;
    } else {
      break;
    }
  }
  return prefix;
}

console.log(fn2(items));
