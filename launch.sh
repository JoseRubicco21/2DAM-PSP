#!/bin/bash

echo "Select which docker-compose to launch:"
echo "1) barebones"
echo "2) nvim"
read -p "Enter your choice [1-2]: " choice

case $choice in
  1)
    cd dockers/barebones
    docker-compose up
    ;;
  2)
    cd dockers/nvim
    docker-compose up
    ;;
  *)
    echo "Invalid choice."
    exit 1
    ;;
esac