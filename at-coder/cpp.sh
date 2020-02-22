#!/bin/sh
file=$1
objfile=`echo $file | sed 's/\.[^\.]*$//'`

# gcc
g++ -g -o $objfile $file
$objfile