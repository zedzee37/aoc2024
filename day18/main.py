from math import sqrt


def parse_input(file_path):
    with open(file_path) as file:
        lines = file.readlines()

    cells_dropped = []
    for line in lines:
        nums = line.split(",")

        x = int(nums[0])
        y = int(nums[1])

        cells_dropped.append((x, y))

    return cells_dropped


def print_grid(grid):
    for row in grid:
        for cell in row:
            print("#" if cell["filled"] else " ", end="")

        print()


def distance(x1, y1, x2, y2):
    return sqrt((float(x2 - x1) ** 2) + (float(y2 - y1) ** 2))


def is_in_grid(x, y):
    return x >= 0 and x < GRID_SIZE and y >= 0 and y < GRID_SIZE


def add(p1, p2):
    return (p1[0] + p2[0], p1[1] + p2[1])


def neighbors(x, y):
    neighboring = []

    directions = [(1, 0), (-1, 0), (0, 1), (0, -1)]
    for dir in directions:
        neighbor = add((x, y), dir)

        if is_in_grid(*neighbor):
            neighboring.append(neighbor)

    return neighboring


def f_cost(cell):
    return cell["g_cost"] + cell["h_cost"]


def lowest_cost_cell(cells):
    cell = (-1, -1)
    i = -1
    lowest_cost = 10e10

    for j, cell_info in enumerate(cells):
        c = cell_info[0]
        cost = f_cost(c)

        if cost < lowest_cost:
            lowest_cost = cost
            cell = cell_info[1]
            i = j

    return (cell, i)


def find_path(grid):
    t_x, t_y = (GRID_SIZE - 1, GRID_SIZE - 1)
    c_x, c_y = 0, 0

    grid[0][0]["g_cost"] = 0
    grid[0][0]["h_cost"] = distance(0, 0, t_x, t_y)

    visited = set()
    cells = [(grid[0][0], (0, 0))]
    while (c_x, c_y) != (t_x, t_y):
        lowest_cost = lowest_cost_cell(cells)

        if lowest_cost[1] == -1:
            break

        if lowest_cost[0] in visited:
            cells.pop(lowest_cost[1])
            continue

        cells.pop(lowest_cost[1])
        c_x, c_y = lowest_cost[0]
        current = grid[c_x][c_y]
        neighboring = neighbors(c_x, c_y)

        visited.add((c_x, c_y))

        for neighbor in neighboring:
            n_x, n_y = neighbor
            neighbor_cell = grid[n_x][n_y]

            if neighbor_cell["filled"] or (n_x, n_y) in visited:
                continue

            neighbor_cell["g_cost"] = current["g_cost"] + 1
            neighbor_cell["h_cost"] = distance(n_x, n_y, t_x, t_y)
            neighbor_cell["came_from"] = current

            cells.append((neighbor_cell, (n_x, n_y)))

    if (c_x, c_y) != (t_x, t_y):
        return None

    return grid[c_x][c_y]


GRID_SIZE = 71


grid = [[]] * GRID_SIZE
for x in range(GRID_SIZE):
    row = [{}] * GRID_SIZE

    for y in range(GRID_SIZE):
        row[y] = {"g_cost": -1, "h_cost": -1, "filled": False, "came_from": None}

    grid[x] = row


cells_dropped = parse_input("input.txt")

for i in range(1024):
    x, y = cells_dropped[i]

    grid[x][y]["filled"] = True

i = 1024

print_grid(grid)

path = find_path(grid)
while path is not None:
    x, y = cells_dropped[i]

    grid[x][y]["filled"] = True

    print_grid(grid)

    path = find_path(grid)

    if path == None:
        print(x, y)
        break
    else:
        print(path["g_cost"])

        i += 1
