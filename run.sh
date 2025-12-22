#!/usr/bin/env bash
set -euo pipefail

if [[ $# -lt 1 || $# -gt 2 ]]; then
  echo "Usage: $0 <dir> [input_file]" >&2
  exit 1
fi

dir=$1
input=${2:-input.txt}

if [[ ! -d "$dir" ]]; then
  echo "Directory '$dir' not found" >&2
  exit 1
fi

if [[ ! -f "$dir/$input" ]]; then
  echo "Input file '$dir/$input' not found" >&2
  exit 1
fi

(
  cd "$dir"
  go run . < "$input"
)
