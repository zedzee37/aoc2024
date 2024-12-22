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

    inline fn eq(self: Vec2, other: Vec2) bool {
        return self.x == other.x and self.y == other.y;
    }

    inline fn manhattanDistance(self: Vec2, other: Vec2) u32 {
        return @abs(self.x - other.x) + @abs(self.y - other.y); 
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
    g: u32,
    cost: u32,
    prev: ?*const Node,
    length: usize,

    const Self = @This();

    fn init(pos: Vec2, dir: Vec2, g: u32, cost: u32, prev: ?*const Node, length: usize,) Self {
        return .{
            .pos = pos,
            .dir = dir,
            .g = g,
            .cost = cost,
            .prev = prev,
            .length = length,
        };
    }

    fn initEmpty(cost: u32) Self {
        return .{
            .pos = Vec2.Zero,
            .dir = Vec2.Zero,
            .cost = cost,
        };
    }

    fn compareCost(n1: *const Node, n2: *const Node) bool {
        return n1.cost < n2.cost;
    }

    fn followPath(self: Self, allocator: Allocator) !std.ArrayList(Vec2) {
        var path = std.ArrayList(Vec2).init(allocator); 
        var current: ?*const Node = &self;
        
        while (current != null) {
            try path.append(current.?.pos);    
            current = current.?.prev;
        }

        return path;
    }
};

const Grid = struct {
    allocator: Allocator,
    lines: std.ArrayList([]u8),

    const Self = @This();

    fn parseInput(allocator: Allocator, file_path: []const u8) !Self {
        var file_handle = try std.fs.cwd().openFile(file_path, .{}); 
        defer file_handle.close(); 
        
        var buf_reader = std.io.bufferedReader(file_handle.reader());
        var in_stream = buf_reader.reader();

        var lines = std.ArrayList([]u8).init(allocator);
        while (try in_stream.readUntilDelimiterOrEofAlloc(allocator, '\n', 1024,)) |line| {
            try lines.append(line); 
        }

        return .{
            .allocator = allocator,
            .lines = lines
        };
    }

    fn deinit(self: Self) void {
        for (self.lines.items) |line| {
            self.allocator.free(line);
        }
        self.lines.deinit();
    }

    inline fn size(self: Self) usize {
        return self.lines.items.len;
    }

    inline fn isOffGrid(self: Self, pos: Vec2) bool {
        return pos.x < 0 or pos.y < 0 or pos.x >= self.size() or pos.y >= self.size();
    }

    inline fn get(self: Self, pos: Vec2) u8 {
        const x = @as(usize, @intCast(pos.x));
        const y = @as(usize, @intCast(pos.y));

        return self.lines.items[y][x];
    }

    fn getShortestPaths(self: Self, arena: *std.heap.ArenaAllocator) !std.ArrayList(*Node) {
        const allocator = arena.allocator();
        const normalize_factor = 100;
        var paths = std.ArrayList(*Node).init(allocator);
        
        const one_off = @as(i32, @intCast(self.size() - 2));
        const start_pos = Vec2.init(1, one_off);
        const end_pos = Vec2.init(one_off, 1);

        var current = PriorityQueue(*const Node, Node.compareCost).init(allocator);
        defer current.deinit();
        
        const start_node = try allocator.create(Node);
        start_node.* = Node.init(
            start_pos, 
            Vec2.init(1, 0), 
            0, 
            start_pos.manhattanDistance(end_pos) * normalize_factor, 
            null, 
            0
        );
        try current.insert(start_node);

        var lowest_cost: u32 = std.math.maxInt(u32);
        var lowest_length: usize = std.math.maxInt(usize);
        while (current.size() > 0) {
            const node = current.pop();

            const neighbor_directions: [3]Vec2 = .{
                node.dir,
                node.dir.right(),
                node.dir.left(),
            };
            for (neighbor_directions) |dir| {
                const neighbor = node.pos.add(dir);

                if (self.isOffGrid(neighbor) or self.get(neighbor) == '#') {
                    continue;
                }
                
                var g_cost: u32 = if (dir.eq(node.dir)) 1 else 1001;
                g_cost += node.g; 

                const h_cost = neighbor.manhattanDistance(end_pos) * normalize_factor;

                const cost = g_cost + h_cost;
                std.debug.print("{}\n", .{cost});

                const length = node.length + 1;

                if (cost > lowest_cost or length > lowest_length) {
                    continue;
                }

                const neighbor_node = try allocator.create(Node);
                neighbor_node.* = Node.init(neighbor, dir, g_cost, cost, node, length); 

                if (neighbor.eq(end_pos)) {
                    lowest_cost = @min(lowest_cost, cost);
                    lowest_length = @min(lowest_length, length);
                    std.debug.print("found shortest\n", .{});
                    try paths.append(neighbor_node);
                } else {
                    try current.insert(neighbor_node);
                }
            }
        }

        // filter paths
        var filtered_paths = std.ArrayList(*Node).init(allocator);
        for (paths.items) |end| {
            if (end.cost <= lowest_cost and end.length <= lowest_length) {
                try filtered_paths.append(end);
            }
        }

        return filtered_paths;
    }
};

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    const allocator = gpa.allocator();

    const grid = try Grid.parseInput(allocator, "input.txt");
    defer grid.deinit();

    var path_arena = std.heap.ArenaAllocator.init(allocator);
    defer path_arena.deinit();

    const shortest_paths = try grid.getShortestPaths(&path_arena);

    var bestPathSpots = std.AutoHashMap(Vec2, void).init(allocator);
    defer bestPathSpots.deinit();

    for (shortest_paths.items) |end| {
        std.debug.print("{}\n", .{end.length});
        const path = try end.followPath(allocator);
        defer path.deinit();

        for (path.items) |pos| {
            try bestPathSpots.put(pos, {});
        }
    }

    std.debug.print("{}\n", .{bestPathSpots.count()});
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

test "Parsing Input" {
    const allocator = std.testing.allocator;
    const grid = try Grid.parseInput(allocator, "input.txt");
    defer grid.deinit();
}
