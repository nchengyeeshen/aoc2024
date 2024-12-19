import sys
from functools import cache

parts = open(sys.argv[1]).read().split("\n\n")

towels = set(parts[0].split(", "))

max_towel_len = max(len(towel) for towel in towels)

desired = parts[1].splitlines()


@cache
def ways(design):
    if design == "":
        return 1

    result = 0
    for i in range(
        1,
        min(len(design), max_towel_len) + 1,
    ):
        if design[:i] not in towels:
            continue
        result += ways(design[i:])

    return result


print(sum(ways(design) for design in desired))
