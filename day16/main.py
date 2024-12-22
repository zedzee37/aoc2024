class PriorityQueue:
    def __init__(self, cmp):
        self.cmp = cmp
        self.heap = []

    @staticmethod
    def parent(idx):
        return (idx - 1) // 2
    
    @staticmethod
    def left(idx):
        return idx * 2 + 1

    @staticmethod
    def right(idx):
        return idx * 2 + 2

    def swap(self, i1, i2):
        self.heap[i1], self.heap[i2] = self.heap[i2], self.heap[i1]

    def check_heap_up(self):
        idx = len(self.heap) - 1
        while idx != 0:
            element = self.heap[idx]
            parent_idx = PriorityQueue.parent(idx)
            parent_element = self.heap[parent_idx]

            if self.cmp(element, parent_element):
                self.swap(idx, parent_idx)
                idx = parent_idx
            else:
                break

    def insert(self, element):
        self.heap.append(element)
        self.check_heap_up()

    def check_heap_down(self, idx):
        smallest = idx
        l = PriorityQueue.left(idx)
        r = PriorityQueue.right(idx)

        if l < len(self.heap) and self.cmp(self.heap[l], self.heap[smallest]):
            smallest = l

        if r < len(self.heap) and self.cmp(self.heap[r], self.heap[smallest]):
            smallest = r

        if smallest != idx:
            self.swap(idx, smallest)
            self.check_heap_down(smallest)

    def pop(self):
        self.swap(0, -1)
        ret = self.heap.pop()
        self.check_heap_down(0)
        return ret
    
    def __len__(self):
        return len(self.heap)


def right(dir):
    return dir * -1j


def left(dir):
    return dir * 1j


def is_off_grid(grid, pos):
    x = int(pos.real)
    y = int(pos.imag)

    return x < 0 or y < 0 or x >= len(grid) or y >= len(grid) 


def get(grid, pos):
    x = int(pos.real)
    y = int(pos.imag)

    return grid[y][x]


class Node:
    def __init__(self, pos, dir, previous):
        self.pos = pos
        self.dir = dir
        self.previous = previous


def get_shortest_paths(grid):
    paths = set()
    one_off = len(grid) - 2
    start_pos = 1 + (1j * one_off)
    end_pos = one_off + 1j

    visited = {}
    current = PriorityQueue(lambda n1, n2: n1[1] < n2[1])

    current.insert((Node(start_pos, 1 + 0j, set()), 0))
    end_path_cost = None
    while len(current) > 0:
        node, cost = current.pop()

        if end_path_cost and cost > end_path_cost:
            break
        elif node.pos == end_pos:
            end_path_cost = cost
            paths = paths.union(node.previous)
        elif (node.pos, node.dir) not in visited or visited[(node.pos, node.dir)] >= cost:
            visited[(node.pos, node.dir)] = cost

            next_pos = node.pos + node.dir
            if not is_off_grid(grid, next_pos) and get(grid, next_pos) != '#':
                previous = set(node.previous)
                previous.add(node.pos)

                new_node = Node(next_pos, node.dir, previous)
                current.insert((new_node, cost + 1))

            current.insert((Node(node.pos, right(node.dir), set(node.previous)), cost + 1000))
            current.insert((Node(node.pos, left(node.dir), set(node.previous)), cost + 1000))

    return len(paths) + 1


with open("input.txt", 'r') as file:
    grid = file.read().split('\n')[:-1]


path_count = get_shortest_paths(grid)
print(path_count)
