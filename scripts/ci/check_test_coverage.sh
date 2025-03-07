#!/bin/bash

coverage=$(go tool cover -func=coverage.out | grep total | awk '{print substr($3, 1, length($3)-1)}')
echo "Total coverage: $coverage%"
coverage_threshold=70.0
if (( $(echo "$coverage < $coverage_threshold" | /usr/bin/bc -l) )); then
  echo "Test coverage ($coverage%) is below the threshold ($coverage_threshold%)."
  exit 1
fi