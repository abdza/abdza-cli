#!/bin/bash

targetdir=`bm -g $1`
if [ -d "$targetdir" ]; then
    echo $targetdir
    cd $targetdir
else
    echo "$targetdir"
fi
