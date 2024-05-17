#!/bin/bash
# name: lint.sh
# description: Run all the linting tools

gum spin --spinner dot --title "Running Static Check" --show-output -- \
	staticcheck ./...

gum spin --spinner dot --title "Running GolangCI Lint" --show-output -- \
	golangci-lint run

gum spin --spinner dot --title "Running SQLC Vet" --show-output -- \
	sqlc vet

gum spin --spinner dot --title "Running Gocritic" --show-output -- \
	gocritic check ./... -enableAll
