import sys
from functools import cache

parts = open(sys.argv[1]).read().split("\n\n")

towels = set(parts[0].split(", "))

max_towel_len = max(len(towel) for towel in towels)

desired = parts[1].splitlines()


@cache
def possible(design):
    if design == "":
        return True

    result = False
    for i in range(
        1,
        min(len(design), max_towel_len) + 1,
    ):
        if design[:i] not in towels:
            continue
        result = result or possible(design[i:])

    return result


print(sum(1 for design in desired if possible(design)))
