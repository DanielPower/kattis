actions = {
    "S": 0,
    "R": 1,
    "P": 2,
}

rounds = int(input())
sven = [actions[i] for i in input()]
friend_count = int(input())
friend_plays = []
for i in range(friend_count):
    friend_plays.append([actions[i] for i in input()])

sven_score = 0
sven_possible_score = 0


def beats(a, b):
    return a - b == 1 or a - b == -2


def ties(a, b):
    return a == b


for round in range(rounds):
    totals = [0, 0, 0]
    for friend in friend_plays:
        totals[friend[round]] += 1
        if beats(sven[round], friend[round]):
            sven_score += 2
        elif ties(sven[round], friend[round]):
            sven_score += 1

    sven_possible_score += max(
        totals[0] * 2 + totals[1],
        totals[1] * 2 + totals[2],
        totals[2] * 2 + totals[0],
    )

print(sven_score)
print(sven_possible_score)
