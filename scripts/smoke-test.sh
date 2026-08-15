#!/usr/bin/env bash
set -euo pipefail

BACKEND_URL="${BACKEND_URL:-http://localhost:9000/fortunes}"
FRONTEND_URL="${FRONTEND_URL:-http://localhost:8080}"
RETRIES="${RETRIES:-15}"
SLEEP_SECONDS="${SLEEP_SECONDS:-2}"

wait_for() {
  local name="$1" url="$2"
  echo "Checking $name at $url"
  for ((i=1; i<=RETRIES; i++)); do
    if curl --fail --silent --output /dev/null "$url"; then
      echo "$name is up"
      return 0
    fi
    sleep "$SLEEP_SECONDS"
  done
  echo "$name did not become healthy in time" >&2
  return 1
}

wait_for "Backend" "$BACKEND_URL"
wait_for "Frontend" "$FRONTEND_URL"