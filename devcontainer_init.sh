#!/usr/bin/sh

git config --local core.hooksPath .githooks/

go mod download

localstack auth set-token $LOCALSTACK_AUTH_TOKEN
localstack start -d