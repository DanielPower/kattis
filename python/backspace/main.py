value = input()
output = []
for char in value:
    output.pop if char == "<" else output.append(char)
print("".join(output))
