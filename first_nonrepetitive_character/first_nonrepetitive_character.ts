const fn = (s: string): number => {
  const list: number[] = Array.from({ length: 26 }).fill(0) as number[];

  for (const l of s) {
    const pos = l.charCodeAt(0) - "a".charCodeAt(0);
    list[pos]++;
  }

  for (let i = 0; i < s.length; i++) {
    const current = s[i];

    const posToLookup = current.charCodeAt(0) - "a".charCodeAt(0);

    if (list[posToLookup] === 1) {
      return i;
    }
  }

  return -1;
};

const fn2 = (s: string): number => {
  const map: Map<string, number> = new Map();

  for (const l of s) {
    if (!map.has(l)) {
      map.set(l, 1);
      continue;
    }
    const current = map.get(l)!;
    map.set(l, current + 1);
  }

  for (let i = 0; i < s.length; i++) {
    const current = s[i];
    if (map.get(current)! === 1) {
      return i;
    }
  }

  return -1;
};
