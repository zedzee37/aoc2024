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

        var n1 = split[0]
        var n2 = split[1]

        if not result.hasKey(n1):
            result[n1] = @[]
        
        result[n1].add(n2)
        

let orderExpr = re"(\d+,)+\d+"

proc readOrder(contents: string): seq[seq[int]] =
    result = @[]
    let order = findAll(contents, orderExpr)

    for line in order:
        let split = line.split(re",").map(parseInt)
        result.add(split)


proc checkOrder(mappings: Table[int, seq[int]], order: seq[int]): bool =
    var seen = toHashSet[int]([])

    for num in order:
        seen.incl(num)

        if mappings.hasKey(num):
            let after = mappings[num]

            for val in after:
                if val in seen:
                    return false
    
    return true


proc part1(contents: string) =
    let mappings = readMappings(contents)
    let order = readOrder(contents)

    var sum = 0
    for line in order:
        if checkOrder(mappings, line):
            let middleIndex = int(len(line) / 2)
            sum += line[middleIndex]

    echo sum


let contents = readFile("input.txt")
part1(contents)
        