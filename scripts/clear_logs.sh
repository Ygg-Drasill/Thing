#!/bin/bash

# Get the directory of the script
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"

LOGS="$DIR/../src/logs/penalties.log"
if [ ! -f "$LOGS" ]; then
    echo "Logs file not found!"
    exit 1
else
    echo "Following logs found: $(basename "$LOGS")"
fi

echo "Do you want to clear logs? (y/n)"
read -r response
if [ "$response" == "y" ]; then
    echo "Clearing logs..."
    rm "$DIR/../src/logs/penalties.log"
    echo "Logs cleared!"
    exit 0
elif [ "$response" == "n" ]; then
    echo "Logs not cleared!"
    exit 0
else
    echo "Invalid response!"
    exit 1
fi