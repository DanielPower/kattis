def char_range(c1, c2):
    """Generates the characters from `c1` to `c2`, inclusive."""
    for c in range(ord(c1), ord(c2) + 1):
        yield chr(c)


VOWELS = {"A", "E", "I", "O", "U"}
LETTERS = {letter for letter in char_range("A", "Z")}


def validate_word(word):
    cons = 0
    vow = 0
    has_l = False
    for char in word:
        if char == "L":
            has_l = True
        if char in VOWELS:
            cons = 0
            vow += 1
        else:
            vow = 0
            cons += 1
        if cons == 3 or vow == 3:
            return False
    return has_l


def around(word, index):
    return {
        letter
        for letter in word[max(0, index - 2) : index]
        + word[index : min(len(word), index + 3)]
    }


def make_words(word):
    underscore = word.find("_")
    if underscore == -1:
        return 1 if validate_word(word) else 0
    sum = 0
    for char, value in [("L", 1), ("A", 5), ("Q", 20)]:
        new_word = word[:underscore] + char + word[underscore + 1 :]
        sum += value * make_words(new_word)
    return sum


word = input()
print(make_words(word))
