function longestCommonPrefix(strs: string[]): string {
  let longestCommonPrefix = strs[0];

  if (!longestCommonPrefix) return "";

  for (const str of strs) {
    longestCommonPrefix = getCommonPrefix(longestCommonPrefix, str);
    if (longestCommonPrefix === "") return "";
  }
  return longestCommonPrefix;
}

function getCommonPrefix(a: string, b: string) {
  let i = 0;
  const minLength = Math.min(a.length, b.length);

  while (i < minLength && a[i] === b[i]) {
    i++;
  }

  return a.substring(0, i);
}
