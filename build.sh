#!/bin/bash
# Check if an environment variable is passed

if [ -z "$1" ]; then
    echo "Please specify the environment (production)"
    exit 1
fi

# Copy the appropriate ignore file based on the environment variable
case "$1" in
  "production")
    cp -f production.containerignore .containerignore
    ;;
  *)
    echo "Invalid environment variable. Must be one of: production"
;;
esac

# Build the image using Podman Compose
podman compose up -d
