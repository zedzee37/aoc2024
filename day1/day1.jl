function parseLine(line::String)
    current = 1

    res = parseNum(line, current)
    firstNum = res[1]
    current = res[2]

    current = consumeNonDigits(line, current)

    res = parseNum(line, current)
    secondNum = res[1]

    return (firstNum, secondNum)
end

function parseNum(line::String, start::Int)
    current = start

    while current <= lastindex(line) && isdigit(line[current])
        current += 1
    end

    numStr = line[start:current - 1]
    return (parse(Int, numStr), current)
end

function consumeNonDigits(line::String, start::Int)
    current = start

    while !isdigit(line[current])
        current += 1
    end

    return current
end

list1 = []
list2 = []

for line in readlines("input.txt")
    res = parseLine(line)
    num1 = res[1]
    num2 = res[2]
    push!(list1, num1)
    push!(list2, num2)
end

count = Dict()

for num in list2
    if !haskey(count, num)
        count[num] = 0
    end
    count[num] += 1
end

sum = 0
for num in list1
    if !haskey(count, num)
        continue
    end
    
    global sum += num * count[num]
end

println(string(sum))