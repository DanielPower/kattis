(n, x1, y1) = [int(p) for p in input().split(' ')]
(x2, y2) = [n - p for p in [x1, y1]]
print(max(x1 * y1, x1 * y2, x2 * y1, x2 * y2) * 4)
