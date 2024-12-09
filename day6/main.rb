require "set"

def find_guard_pos(grid)
  for y in 0..grid.length
    for x in 0..grid.length
      if grid[y][x] == '^'
        return x, y
      end
    end
  end
end

def right(dx, dy)
  return dy, -dx
end

def is_off_grid(size, x, y)
  return x < 0 || y < 0 || x >= size || y >= size
end

def trace_path(grid, start_pos, start_dir, obstacle)
  x, y = start_pos 
  dx, dy = start_dir
  size = grid.length
  path = Set.new

  turn_path = Set.new

  loop do
    next_x, next_y = x + dx, y - dy
    
    if is_off_grid(size, next_x, next_y)
      break
    end

    while [next_x, next_y] == obstacle || grid[next_y][next_x] == '#'
      dx, dy = right(dx, dy)
      next_x, next_y = x + dx, y - dy

      state = [next_x, next_y, dx, dy].hash
      
      if turn_path.include?(state)
        return true, path
      end
      turn_path.add([next_x, next_y, dx, dy].hash)
    end

    if is_off_grid(size, next_x, next_y)
      break
    end

    x, y = next_x, next_y
    path.add([x, y, dx, dy].freeze)
  end  

  return false, path
end

contents = File.read("input.txt")
grid = contents.split("\n")
guard_pos = find_guard_pos(grid)
_, path = trace_path(grid, guard_pos, [0, 1], [-1, -1])

count = 0
for state in path
  x, y, dx, dy = state[0], state[1], state[2], state[3]
  
  looped, _ = trace_path(grid, [x - dx, y + dy], [dx, dy], [x, y])

  if looped
    count += 1
  end
end

puts(count)
