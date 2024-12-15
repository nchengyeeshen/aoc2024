import sys


gridtext, moves = open(0 if sys.argv[1] == "-" else sys.argv[1]).read().split("\n\n")

expansion = {"#": ["#", "#"], "O": ["[", "]"], ".": [".", "."], "@": ["@", "."]}

grid = []
for row in gridtext.splitlines():
    buf = []
    for cell in row:
        buf.extend(expansion[cell])
    grid.append(buf)

moves = "".join(moves.splitlines())

for r, row in enumerate(grid):
    for c, cell in enumerate(row):
        if cell == "@":
            break
    else:
        continue
    break

for move in moves:
    targets = [(r, c)]
    dr = {"^": -1, "v": 1}.get(move, 0)
    dc = {"<": -1, ">": 1}.get(move, 0)
    unblocked = True
    i = 0
    while i < len(targets):
        cr, cc = targets[i]
        nr = cr + dr
        nc = cc + dc

        if (nr, nc) in targets:
            i += 1
            continue

        if grid[nr][nc] == "#":
            unblocked = False
            break

        if grid[nr][nc] == "[":
            targets.append((nr, nc))
            targets.append((nr, nc + 1))

        if grid[nr][nc] == "]":
            targets.append((nr, nc))
            targets.append((nr, nc - 1))

        i += 1

    if not unblocked:
        continue

    chars = {(tr, tc): grid[tr][tc] for tr, tc in targets}

    for tr, tc in targets:
        grid[tr][tc] = "."

    for tr, tc in targets:
        grid[tr + dr][tc + dc] = chars[(tr, tc)]

    r, c = r + dr, c + dc

print(
    sum(
        100 * r + c
        for r, row in enumerate(grid)
        for c, cell in enumerate(row)
        if cell == "["
    )
)
