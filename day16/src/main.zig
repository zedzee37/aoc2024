const std = @import("std");
const Allocator = std.mem.Allocator;

const Vec2 = struct {
    x: i32,
    y: i32,

    const Zero: Vec2 = .{ .x = 0, .y = 0 };

    fn init(x: i32, y: i32) Vec2 {
        return .{
            .x = x,
            .y = y,
        };
    }

    inline fn add(self: Vec2, other: Vec2) Vec2 {
        return .{
            .x = self.x + other.x,
            .y = self.y + other.y,
        };
    }

    inline fn right(self: Vec2) Vec2 {
        return .{
            .x = -self.y,
            .y = self.x,
        };
    }

    inline fn left(self: Vec2) Vec2 {
        return .{
            .x = self.y,
            .y = -self.x,
        };
    }
};

fn PriorityQueue(comptime T: type, comptime compareFn: fn(a: T, b: T) bool) type {
    return struct {
        heap: std.ArrayList(T),

        const Self = @This();

        fn init(allocator: Allocator) Self {
            return .{
                .heap = std.ArrayList(T).init(allocator),
            };
        }

        fn deinit(self: Self) void {
            self.heap.deinit();
        }

        inline fn right(idx: usize) usize {
            return idx * 2 + 2;
        }

        inline fn left(idx: usize) usize {
            return idx * 2 + 1;
        }

        inline fn parent(idx: usize) usize {
            return (idx - 1) / 2;
        }
        
        inline fn get(self: Self, idx: usize) T {
            return self.heap.items[idx];
        }
        
        fn swap(self: *Self, a: usize, b: usize) void {
            if (a >= self.size() or b >= self.size()) {
                @panic("Index out of bounds in swap");
            }
            std.mem.swap(T, &self.heap.items[a], &self.heap.items[b]);
        }

        fn checkHeapUp(self: *Self) void {
            var idx = self.size() - 1; 
            while (idx != 0) {
                const element = self.get(idx);
                const parent_idx = Self.parent(idx);
                const parent_element = self.get(parent_idx);

                if (compareFn(element, parent_element)) {
                    self.swap(idx, parent_idx);
                    idx = parent_idx;
                } else {
                    break;
                }
            }
        }

        fn insert(self: *Self, element: T) !void {
            try self.heap.append(element); 
            self.checkHeapUp();
        }

        fn checkHeapDown(self: *Self, idx: usize) void {
            var smallest = idx;
            const l = Self.left(idx);
            const r = Self.right(idx);
            
            if (l < self.size() and compareFn(self.get(l), self.get(smallest))) {
                smallest = l;
            }
            if (r < self.size() and compareFn(self.get(r), self.get(smallest))) {
                smallest = r;
            }
            if (smallest != idx) {
                self.swap(idx, smallest);
                self.checkHeapDown(smallest);
            }
        }

        fn pop(self: *Self) T {
            self.swap(0, self.size() - 1);
            const ret = self.heap.orderedRemove(self.size() - 1);
            self.checkHeapDown(0);
            return ret;
        }

        inline fn size(self: Self) usize {
            return self.heap.items.len;
        }
    };
}

const Node = struct {
    pos: Vec2,
    dir: Vec2,
    cost: u32,

    const Self = @This();

    fn init(pos: Vec2, dir: Vec2, cost: u32) Self {
        return .{
            .pos = pos,
            .dir = dir,
            .cost = cost,
        };
    }

    fn initEmpty(cost: u32) Self {
        return .{
            .pos = Vec2.Zero,
            .dir = Vec2.Zero,
            .cost = cost,
        };
    }

    fn compareCost(n1: Node, n2: Node) bool {
        return n1.cost < n2.cost;
    }
};

pub fn main() !void {
}

fn testCompareUsize(a: usize, b: usize) bool {
    return a < b;
}

test "Priority Queue" {
    const allocator = std.testing.allocator;
    var queue = PriorityQueue(usize, testCompareUsize).init(allocator);
    defer queue.deinit();

    try queue.insert(10);
    try std.testing.expect(queue.pop() == 10);

    try queue.insert(13);
    try queue.insert(53);
    try queue.insert(23);

    try std.testing.expect(queue.pop() == 13);
    try std.testing.expect(queue.pop() == 23);
    try std.testing.expect(queue.pop() == 53);
}

test "Node Priority Queue" {
    const allocator = std.testing.allocator;
    var queue = PriorityQueue(Node, Node.compareCost).init(allocator);
    defer queue.deinit();

    try queue.insert(Node.initEmpty(10));
    try queue.insert(Node.initEmpty(30));
    try queue.insert(Node.initEmpty(20));

    try std.testing.expect(queue.pop().cost == 10);
    try std.testing.expect(queue.pop().cost == 20);
    try std.testing.expect(queue.pop().cost == 30);
}
