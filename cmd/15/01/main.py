grid, moves = open(0).read().split("\n\n")

grid = [list(row) for row in grid.splitlines()]

moves = "".join(moves.splitlines())

robot_pos = None
for r, row in enumerate(grid):
    for c, cell in enumerate(row):
        if cell == "@":
            robot_pos = (r, c)
            break
    if robot_pos is not None:
        break

lookup = {"<": (0, -1), ">": (0, 1), "^": (-1, 0), "v": (1, 0)}


def helper(r, c, move):
    global robot_pos

    if r < 0 or r >= len(grid) or c < 0 or c >= len(grid[0]):
        return False

    if grid[r][c] == "#":
        return False

    if grid[r][c] == ".":
        return True

    nr, nc = r + lookup[move][0], c + lookup[move][1]
    can_move = helper(nr, nc, move)
    if not can_move:
        return False

    if grid[r][c] == "@":
        robot_pos = (nr, nc)

    grid[r][c], grid[nr][nc] = grid[nr][nc], grid[r][c]
    return True


for move in moves:
    helper(*robot_pos, move)


print(
    sum(
        100 * r + c
        for r, row in enumerate(grid)
        for c, cell in enumerate(row)
        if cell == "O"
    )
)
