import sys
from collections import deque

contents = open(sys.argv[1] if sys.argv[1] != "-" else 0).read()

start = None
end = None
grid = []
for r, row in enumerate(contents.splitlines()):
    buf = []
    for c, cell in enumerate(row):
        buf.append(cell)
        if cell == "S":
            start = (r, c)
        if cell == "E":
            end = (r, c)
    grid.append(buf)

path = None


def shortest_path(grid, start, end):
    visited = set()
    q = deque()
    q.append((start, []))
    paths = []

    while len(q) > 0:
        (r, c), path = q.popleft()
        visited.add((r, c))
        if (r, c) == end:
            paths.append(path[:])
            continue

        for dr, dc in [(-1, 0), (1, 0), (0, -1), (0, 1)]:
            nr, nc = r + dr, c + dc

            if grid[nr][nc] == "#" or (nr, nc) in path:
                continue

            q.append(((nr, nc), path[:] + [(r, c)]))

    return paths


def score(path):
    curr_dir = (0, 1)
    turns = 0
    for (pr, pc), (cr, cc) in zip(path, path[1:]):
        dr, dc = cr - pr, cc - pc
        turned = (dr, dc) != curr_dir
        curr_dir = (dr, dc)
        if turned:
            turns += 1
    return len(path) + turns * 1000


best_path = min(shortest_path(grid, start, end), key=score)

print(score(best_path))
