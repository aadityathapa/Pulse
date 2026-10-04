#!/usr/bin/env bash
URL="${1:-http://localhost:8080/health}"
GREEN=$'\033[32m'
RED=$'\033[31m'
RESET=$'\033[0m'

while true; do
  if out=$(curl -fsS --max-time 2 "$URL" 2>/dev/null); then
    echo "${GREEN}● UP${RESET}   $(date +%T)  $out"
  else
    echo "${RED}● DOWN${RESET} $(date +%T)"
  fi
  sleep 2
done
