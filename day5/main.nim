import std/tables
import std/re
from std/sequtils import map
import std/strutils
import std/sets

let mappingExpr = re"(\d+)\|(\d+)"

proc readMappings(contents: string): Table[int, seq[int]] =
    result = initTable[int, seq[int]]()
    
    let mappings = findAll(contents, mappingExpr)

    for mapping in mappings:
        let split = mapping.split(re"\|").map(parseInt)

        if not result.hasKey(split[1]):
            result[split[1]] = @[]
        
        result[split[1]].add(split[0])
        

let orderExpr = re"(\d+,*)+"

proc readOrder(contents: string): seq[seq[int]] =
    result = @[]
    let order = findAll(contents, orderExpr)

    for line in order:
        let split = line.split(re",").map(parseInt)
        result.add(split)


proc checkOrder(mappings: Table[int, seq[int]], order: seq[int]): bool =
    var failNums = toHashSet[int]([])
    var seen = toHashSet[int]([])

    for num in order:
        if num in failNums:
            return false
            
        if mappings.hasKey(num):
            let before = toHashSet(mappings[num])

            echo failNums
            failNums = failNums + before
            echo before
            echo failNums
    
    return true

let contents = readFile("input.txt")
let mappings = readMappings(contents)
let order = readOrder(contents)

var sum = 0
for line in order:
    if checkOrder(mappings, line):
        let middleIndex = int(len(line) / 2)
        sum += line[middleIndex]

echo sum
        