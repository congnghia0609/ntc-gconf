#!/bin/bash
# Author:       nghiatc
# Email:        congnghia0609@gmail.com

#source /etc/profile

echo "Install library dependencies..."
go get -u github.com/tools/godep
go get -u github.com/spf13/viper

echo "Install dependencies complete..."
