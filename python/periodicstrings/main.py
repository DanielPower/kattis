value = input()
for segmentLength in range(1, len(value) + 1):
    if len(value) % segmentLength != 0:
        continue
    segmentCount = len(value) // segmentLength
    segments = []
    for i in range(segmentCount):
        segment = []
        for j in range(segmentLength):
            segment.append(value[i * segmentLength + j])
        segments.append(segment)
    is_valid = True
    if segmentLength == 1:
        if all(v == value[0] for v in value):
            print(1)
            break
    for i in range(segmentCount - 1):
        if segments[i] != segments[i + 1][1:] + [segments[i + 1][0]]:
            is_valid = False
            break
    if is_valid:
        print(segmentLength)
        break
