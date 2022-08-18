from sys import stdin

best = 0
best_id = 0

for index, line in enumerate(stdin):
    value = sum([int(p) for p in line.split(' ')])
    if value > best:
        best = value
        best_id = index + 1

print(best_id, best)
