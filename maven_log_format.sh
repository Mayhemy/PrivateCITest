#!/bin/bash

# Input and output file names
input_file="cleaned_maven_log.txt"
output_file="formatted_maven_log.md"

# Check if input file exists
if [ ! -f "$input_file" ]; then
    echo "Input file does not exist."
    exit 1
fi

# Process the input file and save to output file
while IFS= read -r line; do
    # Replace "INFO", "WARNING", and "ERROR" with bold versions within the line
    modified_line="${line//INFO/**INFO**}"
    modified_line="${modified_line//WARNING/**WARNING**}"
    modified_line="${modified_line//ERROR/**ERROR**}"
    
    echo "$modified_line" >> "$output_file"
done < "$input_file"

echo "Formatted log saved to $output_file"

