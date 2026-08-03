#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")/.."

echo "Building and starting services..."
docker compose up -d --build --wait

echo "Waiting for Kafka to stabilize..."
sleep 20

echo "Service status:"
docker compose ps
