#!/bin/bash
# name: format.sh
# description: Run all the formatting tools
go fmt *.go
golines -w --max-len=80 *.go
