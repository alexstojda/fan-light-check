#!/bin/sh

go build -o ./build/fan-light-check .
chmod +x ./build/fan-light-check
sudo cp ./build/fan-light-check /usr/bin/
