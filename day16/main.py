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
    def __init__(self, c):
        self.c = c


with open("input.txt", 'r') as file:
    grid = file.read().split('\n')[:-1]

