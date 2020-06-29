#!/bin/bash

repo_dir=$(dirname "$0")

function create_leetcode_dir() {
  if [[ $# != 1 ]]; then
    echo "usage create_leetcode_dir problem_name"
  fi

  number=$(cut -d'.' -f1 <<<"$1")
  number_padding=$(printf '%04d' "$number")
  name=$(cut -d'.' -f2 <<<"$1")
  name=$(echo "$name" | sed -E -e 's/^ *//g')
  name=$(echo "$name" | sed -E -e 's/ /-/g')

  full_dir_name="${repo_dir}/problems/p${number_padding}.${name}"
  if [[ -d "$full_dir_name" ]]; then
    echo "${full_dir_name} already exist"
    exit 1
  fi

  mkdir -p "${full_dir_name}"

  cd "$full_dir_name" || exit
  echo "package p${number}" >"p${number}.go"
}

create_leetcode_dir "$@"
