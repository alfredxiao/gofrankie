#!/bin/bash

go test -count=1 ./... -coverprofile=c.out && go tool cover -html=c.out
