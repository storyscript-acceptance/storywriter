#!/usr/bin/env bash

set -eu

echo "Running omg story..."

output=$(omg run -r write -a template='On a fine &1 he saw &2.' -a words='["morning", "a sparrow"]')
echo "${output}"

echo "Asserting the story has been templated correctly..."

if [ "${output}" == "On a fine morning he saw a sparrow." ]; then
  echo "SUCCESS"
  exit 0
else
  echo "FAILED"
  exit 1
fi
