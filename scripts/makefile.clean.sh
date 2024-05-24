#!/bin/bash
# name: makefile.clean.sh
# description: Cleans the project
#              by removing all the generated files
#              and the node_modules folder
#              and all the *.go files in the data folder
#              and the coverage.out file
# url: github.com/conneroisu/dblogger/scripts/makefile.clean.sh

# if there is a tmp folder, delete it
if [ -d "tmp" ]; then
    rm -rf tmp
fi

# if there is a bin folder, delete it
if [ -d "bin" ]; then
    rm -rf bin
fi

# if there is a node_modules folder, delete it
if [ -d "node_modules" ]; then
    rm -rf node_modules
fi

# if there are *.go files in the data folder, delete them
for file in ./data/*.go; do
    if [ -f "$file" ]; then
        rm "$file"
    fi
done

# if there is a coverage.out file, delete it
if [ -f "coverage.out" ]; then
    rm coverage.out
fi
