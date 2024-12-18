import sys
from collections import deque

M = 7

# y, x
blocked = set()

for line in open(sys.argv[1]).read().splitlines()[:12]:
    x, y = map(int, line.split(","))
    blocked.add((y, x))

sr, sc = 0, 0
er, ec = M - 1, M - 1

# r, c, steps
q = deque([(sr, sc, 0)])

# r, c
seen = set()

while q:
    r, c, steps = q.popleft()
    if (r, c) == (er, ec):
        print(steps)
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
