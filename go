#!/bin/bash

targetdir=`bm -g $1`
if [ -d "$targetdir" ]; then
    cd $targetdir
else
    echo "$targetdir"
fi
