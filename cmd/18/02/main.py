import sys
from collections import deque

M = 71

# x, y
byte_positions = [
    tuple(map(int, line.split(","))) for line in open(sys.argv[1]).read().splitlines()
]

sr, sc = 0, 0
er, ec = M - 1, M - 1


low, high = 0, len(byte_positions)
while high - low > 1:
    mid = low + (high - low) // 2

    blocked = set((y, x) for x, y in byte_positions[:mid])

    has_path = False

    # r, c, steps
    q = deque([(sr, sc, 0)])

    # r, c
    seen = set()

    while q:
        r, c, steps = q.popleft()
        if (r, c) == (er, ec):
            has_path = True
            break
        for nr, nc in [
            (r, c + 1),
            (r, c - 1),
            (r + 1, c),
            (r - 1, c),
        ]:
            if nr < 0 or nr >= M or nc < 0 or nc >= M:
                continue
            if (nr, nc) in blocked:
                continue
            if (nr, nc) in seen:
                continue
            seen.add((nr, nc))
            q.append((nr, nc, steps + 1))

    if has_path:
        low = mid + 1
    else:
        high = mid


print(*byte_positions[low + (high - low) // 2], sep=",")
