def right(dir):
   return dir * -1j

def left(dir):
    return dir * 1j

def is_off_grid(grid, pos):
    return pos.real < 0 or pos.imag < 0 or pos.real >= len(grid) or pos.imag >= len(grid)

def find_end_points(grid):
    start_pos = None
    end_pos = None

    for y in range(len(grid)):
        for x in range(len(grid)):
            pos = x + (1j * y)
            ch = grid[y][x]

            match ch:
                case 'S':
                    start_pos = pos
                case 'E':
                    end_pos = pos

            if start_pos and end_pos:
                return start_pos, end_pos

    return None

def get(grid, pos):
    return grid[int(pos.imag)][int(pos.real)]

def get_shortest_path(grid):
    res = find_end_points(grid)

    if not res:
        return None

    paths = []
    start_pos, end_pos = res
    visited = {}
    current = []
    current.append({
        "pos": start_pos,
        "cost": 0,
        "dir": 1 + 0j,
        "length": 0,
        "prev": None
    })
    lowest_cost = 10e6

    while len(current) > 0:
        min_cost_index = min(range(len(current)), key=lambda i: current[i]["cost"])
        node = current.pop(min_cost_index)

        if node["pos"] in visited:
            continue

        visited[node["pos"]] = node
        neighbor_directions = [
            node["dir"],
            right(node["dir"]),
            left(node["dir"])
        ]

        for dir in neighbor_directions:
            neighbor = node["pos"] + dir

            if is_off_grid(grid, neighbor) or neighbor in visited or get(grid, neighbor) == '#':
                continue

            cost = 1
            if dir != node["dir"]:
                cost += 1000

            neighbor_node = {
                "pos": neighbor,
                "dir": dir,
                "cost": cost + node["cost"],
                "prev": node,
            }
            
            if neighbor == end_pos:
                if lowest_cost == None:
                    lowest_cost = neighbor_node["cost"]
                    paths.append(neighbor_node)
                lowest_cost = min(lowest_cost, neighbor_node["cost"])
                paths.append(neighbor_node)

                new_paths = []
                for path in paths:
                    if path["cost"] <= lowest_cost:
                        new_paths.append(path)
                paths = new_paths
            else:
                current.append(neighbor_node)
    return paths

def get_path(end):
    current = end
    path = []

    while current != None:
        path.append(current)
        current = current["prev"]

        if current == None:
            return path

    return path

def print_path(grid, path):
    for y in range(len(grid)):
        for x in range(len(grid)):
            pos = x + (y * 1j)

            if pos in path:
                print("X", end="")
            else:
                print(grid[y][x], end="")
        print()
    print()

with open("input.txt", 'r') as file:
    grid = file.read().split('\n')
    grid = grid[:len(grid) - 1]

visited = set()
shortest_paths = get_shortest_path(grid)
if shortest_paths:
    for end in shortest_paths:
        print(end)
        print(end["cost"])
        path = get_path(end)
        print_path(grid, path)
        
        for pos in path:
            visited.add(pos)
    print(len(visited))

