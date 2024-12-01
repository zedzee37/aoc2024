#!/bin/bash

cd build || exit 1  # Change to the build directory, exit if it fails
ninja || exit 1     # Build the project, exit if it fails

# Loop through files starting with "day" followed by a number
for file in day*; do
    if [[ -x "$file" && "$file" =~ ^day[0-9]+$ ]]; then
        echo "Running $file..."
        ./"$file"
    fi
done
