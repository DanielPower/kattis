from sys import stdin

input()
for line in stdin:
    value = int(line)
    word = "even" if value % 2 == 0 else "odd"
    print(f"{value} is {word}")

