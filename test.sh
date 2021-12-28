#!/bin/bash

module=$1

cd "$(pwd)/$(dirname "$0")/${module}" 
go test