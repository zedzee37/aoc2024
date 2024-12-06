import std/tables
import std/re
import std/strutils
import std/sets
import std/strformat
import std/sequtils
import std/enumerate

func readMappings(contents: string): Table[int, seq[int]] =
    result = initTable[int, seq[int]]()
    
    let mappings = findAll(contents, re"(\d+)\|(\d+)")

    for mapping in mappings:
        let split = mapping.split(re"\|").map(parseInt)

        var n1 = split[0]
        var n2 = split[1]

        if not result.hasKey(n1):
            result[n1] = @[]
        
        result[n1].add(n2)

func readOrder(contents: string): seq[seq[int]] =
    result = @[]
    let order = findAll(contents, re"(\d+,)+\d+")

    for line in order:
        let split = line.split(re",").map(parseInt)
        result.add(split)

func checkOrder(mappings: Table[int, seq[int]], order: seq[int]): bool =
    var seen = toHashSet[int]([])

    for num in order:
        seen.incl(num)

        if mappings.hasKey(num):
            let after = mappings[num]

            for val in after:
                if val in seen:
                    return false
    
    return true

proc fixOrder(mappings: Table[int, seq[int]], order: seq[int]): seq[int] =
    result = order

    var hasConflict = true
    while hasConflict:
        hasConflict = false
        var seen = toHashSet[int]([])
        var newRes = result

        for i, num in enumerate(result):
            seen.incl(num)

            if not mappings.hasKey(num):
                continue

            let after = mappings[num]

            for val in after:
                if val in seen:
                    let idx = newRes.find(val)  # Find the current index of val

                    if idx >= 0 and idx < len(result):
                        newRes[i] = val
                        newRes[idx] = num

                        hasConflict = true
                        break
        
        result = newRes

    
let contents = readFile("input.txt")

let mappings = readMappings(contents)
let order = readOrder(contents)

var partOneSum = 0
var partTwoSum = 0
for line in order:
    let middleIndex = int(len(line) / 2)

    if checkOrder(mappings, line):
        partOneSum += line[middleIndex]
    else:
        let fixed = fixOrder(mappings, line)
        partTwoSum += fixed[middleIndex]


echo fmt"The solution to part 1 is: {partOneSum}"
echo fmt"The solution to part 2 is: {partTwoSum}"