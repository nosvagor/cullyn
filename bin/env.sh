#!/usr/bin/env bash

# This script parses specified TOML files into a single .env file.
# It supports basic key-value pairs and sections.

set -euo pipefail

output_file="config.env"
toml_files=("./config/config.toml" "./config/secrets.toml")

process_toml_file() {
    local file="$1"
    local section=""

    while IFS= read -r line || [[ -n "$line" ]]; do
        # Trim leading and trailing whitespace
        line="${line#"${line%%[![:space:]]*}"}"
        line="${line%"${line##*[![:space:]]}"}"

        # Ignore comments and empty lines
        [[ $line =~ ^# ]] && continue
        [[ -z $line ]] && continue

        if [[ $line =~ ^\[(.+)\]$ ]]; then
            section=${BASH_REMATCH[1]}
        elif [[ $line =~ ^([^=]+)=(.*)$ ]]; then
            # Key-value pair
            key=${BASH_REMATCH[1]}
            value=${BASH_REMATCH[2]}
            # Remove leading/trailing whitespace and quotes
            key="${key#"${key%%[![:space:]]*}"}"
            key="${key%"${key##*[![:space:]]}"}"
            key="${key#[\"\']}"
            key="${key%[\"\']}"
            value="${value#"${value%%[![:space:]]*}"}"
            value="${value%"${value##*[![:space:]]}"}"
            value="${value#[\"\']}"
            value="${value%[\"\']}"
            if [[ -n $section ]]; then
                echo "${section^^}_${key^^}=${value}"
            else
                echo "${key^^}=${value}"
            fi
        fi
    done < "$file"
}

main() {
    # Clear existing .env file
    : > "$output_file"

    # Process each hardcoded TOML file
    for file in "${toml_files[@]}"; do
        if [[ -f "$file" ]]; then
            process_toml_file "$file" >> "$output_file"
        else
            echo "Warning: $file not found" >&2
        fi
    done

    echo "env file created at $output_file"
}

main "$@"
