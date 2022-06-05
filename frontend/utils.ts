export function classes(...classNames: string[]): string {
  return classNames.reduce<string[]>((acc: string[], cur: string) => {
    const cleaned = cur.trim();
    if (cleaned) {
      acc.push(cleaned);
    }
    return acc;
  }, []).join(' ');
}
