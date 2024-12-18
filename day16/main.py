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

    start_pos, end_pos = res
    visited = set()
    current = []
    current.append({
        "pos": start_pos,
        "cost": 0,
        "dir": 1 + 0j
    })

    i = 0
    while len(current) > 0:
        min_cost_index = min(range(len(current)), key=lambda i: current[i]["cost"])
        node = current.pop(min_cost_index)

        if node["pos"] in visited:
            continue

        visited.add(node["pos"])
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
                "cost": cost + node["cost"]
            }
            
            if neighbor == end_pos:
                return neighbor_node
            
            current.append(neighbor_node)
        i += 1


with open("input.txt", 'r') as file:
    grid = file.read().split('\n')
    grid = grid[:len(grid) - 1]

shortest_path = get_shortest_path(grid)
print(shortest_path["cost"])
