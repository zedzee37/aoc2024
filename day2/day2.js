const fs = require("fs");

function isDigit(char) {
    return !isNaN(parseInt(char, 10)) && char !== ' ';
}

let rows = fs.readFile("input.txt", "utf-8", (err, text) => {
    if (err) {
        console.error("Error reading file:", err);
        return;
    }

    let lines = text.split("\n");
    let rows = [];

    lines.forEach(line => rows.push(line.split(" ").forEach(parseInt)));
    return rows;
});

function consumeNonDigits(line, start) {
    let current = start;
    while (current < line.length && !isDigit(line[current])) {
        current++;
    }
    return current;
}
