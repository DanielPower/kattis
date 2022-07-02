import sys
from collections import deque

n, k = [int(i) for i in input().strip().split(" ")]
weights = deque([int(i) for i in input().strip().split(" ")])

total_weight = sum(weights)
if total_weight % n != 0:
    print("NO")
    sys.exit()
weight_per_person = total_weight // n

for _ in range(len(weights)):
    current_weight = 0
    possible = True
    for weight in weights:
        current_weight += weight
        if current_weight > weight_per_person:
            possible = False
            break
        if current_weight == weight_per_person:
            current_weight = 0
    if possible and current_weight == 0:
        print("YES")
        sys.exit()
    weights.rotate()
print("NO")
