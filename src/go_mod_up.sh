#!/usr/bin/env bash

for dir in $(find . -type d -depth 1); do
	echo $dir
	cd $dir
	GO111MODULE=on go get -u
	cd ..
done
