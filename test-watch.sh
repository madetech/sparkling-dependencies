#!/bin/bash

RED='\033[0;31m'
GREEN='\033[0;32m'
NOCOLOR='\033[0m'

if make test; then
  echo -e "${GREEN}All tests passed.${NOCOLOR}"
else
  echo -e "${RED}Test Failures.${NOCOLOR}"
fi

while true
do
    if inotifywait -qq -r -e create,close_write,modify,move,delete ./ && make test; then
      echo -e "${GREEN}All tests passed.${NOCOLOR}"
    else
      echo -e "${RED}Test Failures.${NOCOLOR}"
    fi
done