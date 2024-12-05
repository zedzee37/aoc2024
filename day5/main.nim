import std/tables
import std/re
from std/sequtils import map
import std/strutils

let mappingExpr = re"(\d+)\|(\d+)"

proc genMappings(contents: string): Table[int, int] =
    result = initTable[int, int]()
    
    let mappings = findAll(contents, mappingExpr)

    for mapping in mappings:
        let split = mapping.split(re"\|").map(parseInt)
        result[split[0]] = split[1]


let contents = readFile("input.txt")
let mappings = genMappings(contents)
