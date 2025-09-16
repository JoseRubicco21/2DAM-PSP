#!/bin/bash

# This script adds an alias to your ~/.bashrc for the launch.sh script

ALIAS_NAME="psp"
SCRIPT_PATH="$(cd "$(dirname "$0")" && pwd)/launch.sh"
BASHRC="$HOME/.bashrc"

# Check if alias already exists
if grep -Fxq "alias $ALIAS_NAME=\"$SCRIPT_PATH\"" "$BASHRC"; then
  echo "Alias '$ALIAS_NAME' already exists in $BASHRC."
else
  echo "alias $ALIAS_NAME=\"$SCRIPT_PATH\"" >> "$BASHRC"
  echo "Alias '$ALIAS_NAME' added to $BASHRC."
  echo "Run 'source $BASHRC' or restart your terminal to use it."
fi
