#!/usr/bin/env bash

set -e

##Create query with keywords java, rmote, location remote.
echo
echo "Creating java-remote, remote query."
echo
curl -X POST http://localhost:9192/queries \
  -H "Content-type: application/json" \
  -d "{\"keywords\": [\"java\", \"remote\"], \"location\": \"remote\", \"active\": true}" \
  -w "\n"

##Create query with keywords go, remote, location remote.
echo
echo "Creating go-remote, remote query."
echo
curl -X POST http://localhost:9192/queries \
  -H "Content-type: application/json" \
  -d "{\"keywords\": [\"go\", \"remote\"], \"location\": \"remote\", \"active\": true}" \
  -w "\n"

##Create query with keywords go, golang, remote, location remote.
echo
echo "Creating go-golang-remote, remote query."
echo
curl -X POST http://localhost:9192/queries \
  -H "Content-type: application/json" \
  -d "{\"keywords\": [\"go\", \"golang\", \"remote\"], \"location\": \"remote\", \"active\": true}" \
  -w "\n"

##Create query with keywords groovy, remote, location remote.
echo
echo "Creating groovy-remote, remote query."
echo
curl -X POST http://localhost:9192/queries \
  -H "Content-type: application/json" \
  -d "{\"keywords\": [\"groovy\", \"remote\"], \"location\": \"remote\", \"active\": true}" \
  -w "\n"


##Create query with keywords kotlin, remote, location remote.
echo
echo "Creating kotlin-remote, remote query."
echo
curl -X POST http://localhost:9192/queries \
  -H "Content-type: application/json" \
  -d "{\"keywords\": [\"kotlin\", \"remote\"], \"location\": \"remote\", \"active\": true}" \
  -w "\n"