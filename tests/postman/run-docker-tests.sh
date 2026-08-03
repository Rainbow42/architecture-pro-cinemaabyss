#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")"

docker build -t cinemaabyss-api-tests .
docker run --rm \
  --network=cinemaabyss-network \
  -v "$(pwd)/reports:/app/reports" \
  cinemaabyss-api-tests \
  --environment docker \
  --reporters cli,junit
