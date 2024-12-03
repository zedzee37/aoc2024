const fs = require("fs");

function isDigit(char) {
    return !isNaN(parseInt(char, 10)) && char !== ' ';
}

fs.readFile("input.txt", "utf-8", (err, text) => {
    if (err) {
        console.error("Error reading file:", err);
        return;
    }

    let lines = text.split("\n");
    let passed = lines.length;

    for (let line of lines) { // Iterate over the content of each line
        let current = 0;
        let prev = undefined;
        let ascending = undefined;

        while (current < line.length - 1) {
            current = consumeNonDigits(line, current); 
            let start = current;

            while (isDigit(line[current])) {
                current++; 
            }

            let numStr = line.slice(start, current); 
            let num = parseInt(numStr);

            if (prev === undefined) {
                prev = num;
                continue;
            }

            if (num === prev) {
                passed--;
                break;
            }

            let diff = num - prev;
            if (Math.abs(diff) > 3) {
                passed--; 
                break;
            }

            if (ascending === undefined) {
                ascending = diff > 0;
            } else if ((ascending && diff < 0) || (!ascending && diff > 0)) {
                passed--;
                break;
            }

            prev = num;
        }
    }

    console.log(passed);
});

function consumeNonDigits(line, start) {
    let current = start;
    while (current < line.length && !isDigit(line[current])) {
        current++;
    }
    return current;
}
