#!/bin/sh

mkdir -p golang.org/x
cd golang.org/x

git clone https://github.com/golang/tools

go get github.com/jstemmer/gotags
go install golang.org/x/tools/cmd/guru
go install golang.org/x/tools/cmd/goimports
