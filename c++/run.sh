#!/bin/sh
mkdir -p build
g++ -g -O2 -std=gnu++17 "$1".cpp -o build/"$1"
./build/"$1"
