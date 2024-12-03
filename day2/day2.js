const fs = require("fs");

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