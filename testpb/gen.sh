#!/bin/bash

protoc -I$(pwd) --go_out=paths=source_relative:. test.proto
