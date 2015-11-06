#!/bin/bash
timeout=$1
cmd=$(echo "$2" | base64 -d)

echo "CMD: $cmd"

timeout $1 nohup $cmd  > /dev/null 2>&1 &
