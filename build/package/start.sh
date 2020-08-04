#!/bin/bash

# launch appboot server
cd /server
./server &

# launch web
cd /app
sh start.sh
