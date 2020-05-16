#!/bin/bash

# Below command will run all tests and generate a coverage report (it should open browser automatically)
go test -count=1 ./... -coverprofile=c.out && go tool cover -html=c.out
