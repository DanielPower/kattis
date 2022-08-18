from sys import stdin

values = set()
for line in stdin:
    values.add(int(line) % 42)
print(len(values))
