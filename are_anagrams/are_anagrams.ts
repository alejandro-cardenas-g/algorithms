function isAnagram(s: string, t: string): boolean {
  if (s.length !== t.length) return false;

  const map: Record<string, number> = {};

  for (let i = 0; i < s.length; i++) {
    const currentS = s[i];
    if (map[currentS]) {
      map[currentS]++;
    } else {
      map[currentS] = 1;
    }
  }

  for (let i = 0; i < s.length; i++) {
    const currentT = t[i];
    if (map[currentT]) {
      map[currentT]--;
    } else {
      return false;
    }
  }

  return Object.values(map).every((item) => item === 0);
}
