import sys
import heapq
from dataclasses import dataclass, field

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


@dataclass(order=True)
class Item:
    cost: int
    r: int = field(compare=False)
    c: int = field(compare=False)
    direction: tuple[int, int] = field(compare=False)
    path: list[tuple[int, int]] = field(compare=False)


items = {start: Item(0, *start, (0, 1), [])}
q = [items[start]]

while len(q) > 0:
    item = heapq.heappop(q)
    if (item.r, item.c) == end:
        break

    for dr, dc in [(-1, 0), (1, 0), (0, -1), (0, 1)]:
        new_cost = item.cost + 1
        if (dr, dc) != item.direction:
            new_cost += 1000

        nr, nc = item.r + dr, item.c + dc
        if grid[nr][nc] == "#" or (item.path and (nr, nc) == item.path[-1]):
            continue

        if (nr, nc) in items:
            next_item = items[(nr, nc)]
            if new_cost < next_item.cost:
                next_item.cost = new_cost
                heapq.heapify(q)
                next_item.path = item.path[:] + [(item.r, item.c)]
        else:
            new_item = Item(
                new_cost, nr, nc, (dr, dc), item.path[:] + [(item.r, item.c)]
            )
            items[(nr, nc)] = new_item
            heapq.heappush(q, new_item)

print(items[*end].cost)
