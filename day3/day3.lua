require("io")

local file = io.open("input.txt", "r")

if not file then
    return
end

local contents = file:read("*all")
file:close()

local pattern = "mul%((%d+),(%d+)%)"
local doPattern = "do%(%)"
local dontPattern = "don%'t%(%)"
print(("don't()"):find(dontPattern))
local current = 1
local sum = 0
local shouldMatch = false

while current < #contents do
    if shouldMatch then
        local start, finish, n1, n2 = contents:find(pattern, current)

        if not finish then
            break
        end

        local dontStart, dontFinish = contents:sub(0, start):find(dontPattern, current)

        if dontFinish then
            shouldMatch = false
            current = dontFinish + 1
        else
            sum = sum + n1 * n2
            current = finish + 1
        end
    else
        local start, finish = contents:find(doPattern, current)

        if not finish then
            break
        end

        current = finish + 1
        shouldMatch = true
    end 
end

print(sum)
