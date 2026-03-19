function groupAnagrams(strs: string[]): string[][] {
  const map = new Map<string, string[]>();

  for (const word of strs) {
    const sortedWord = word
      .split("")
      .sort((a, b) => a.charCodeAt(0) - b.charCodeAt(0))
      .join();
    if (map.has(sortedWord)) {
      map.get(sortedWord)!.push(word);
    } else {
      map.set("sortedWord", [word]);
    }
  }

  return [...map.values()];
}
