from sys import stdin

input()
for line in stdin:
    a, b, c = [int(p) for p in line.split()]
    if a > b - c:
        print("do not advertise")
    elif a < b - c:
        print("advertise")
    else:
        print("does not matter")
