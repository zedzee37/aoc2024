const fs = require("fs");

function validateRow(row) {
    let ascending = undefined;
    let prev = undefined;

    for (let i = 0; i < row.length; i++) {
        let num = row[i];

        if (prev === undefined) {
            prev = num;
            continue;
        }

        let diff = num - prev;

        if (Math.abs(diff) > 3 || diff == 0) {
            return false;
        }
        
        if (ascending === undefined) {
            ascending = diff > 0;
        } else if ((ascending && diff < 0) || (!ascending && diff > 0)) {
            return false; 
        }

        prev = num;
    }

    return true;
}

function getAllPermutations(row) {
    let permutations = [];
    for (let i = 0; i < row.length; i++) {
        permutations.push([...row.slice(0, i), ...row.slice(i + 1)]);
    }
    return permutations;
}


fs.readFile("input.txt", "utf-8", (err, text) => {
    if (err) {
        console.error("Error reading file:", err);
        return
    }

    let lines = text.split("\n");

    let rows = lines
        .map(line => line.split(/[\s\r]+/)
        .filter(str => str !== '')
        .map(x => parseInt(x))
    );

    let failed = rows.filter(row => !validateRow(row));
    let passCount = rows.length - failed.length;
    let dampened = failed
        .map(row => getAllPermutations(row))
        .map(permutations => 
            permutations.map(row => validateRow(row))
                        .some(res => res))
        .filter(row => row).length;

    console.log(passCount + dampened);
});
