#!/bin/bash

if [ "$1" = "readers" ]; then
	for func in $(cat readers.txt)
	do
		echo "************************************************************************"
		echo "Looking for reader: $func"
		grep $func . --exclude="readers.txt" --exclude="readers.out" -r
		echo "************************************************************************"
		echo
	done
elif [ "$1" = "writers" ]; then
	for func in $(cat writers.txt)
	do
		echo "************************************************************************"
		echo "Looking for writer: $func"
		grep $func . --exclude="writers.txt" --exclude="writers.out" -r
		echo "************************************************************************"
		echo
	done
else
	echo "error"
	exit 1
fi
