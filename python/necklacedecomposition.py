def make_necklace(scenario):
    max_zeros = 0
    zeros = 0
    max_index = None
    for index, char in enumerate(scenario):
        if char == "0":
            zeros += 1
            if zeros > max_zeros:
                max_zeros = zeros
                max_index = index
        if char == "1":
            zeros = 0


for _ in range(int(input())):
    scenario = [char for char in input()[::-1]]
    print(scenario)
    necklace_number = 0
    necklaces = [""]
    has_one = False
    while len(scenario):
        char = scenario.pop()
        print(has_one, char)
        if char == "1":
            necklaces[0] += "1"
            has_one = True
        else:
            if has_one:
                scenario.append("0")
                break
            else:
                necklaces[0] += "0"
    print("(" + ")(".join(necklaces) + ")")
