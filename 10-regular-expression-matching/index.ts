function isMatch(s: string, p: string): boolean {
  const normalizedP = normalizePattern(p);
  const regex = new RegExp(`^${normalizedP}$`);
  return regex.test(s)
};

function normalizePattern(pattern) {
  return pattern.replace(/\*+/g, '*');
}
