#!/bin/bash

# Capture all passed arguments
ARGS="$@"

# Run the Makefile commands
make build && make run ARGS="$ARGS"