#!/bin/bash

# colours
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
# ----

declare -p commit_message
echo -e "${YELLOW}Enter commit message :${SET}"
read commit_message

declare -i day
echo -e "${YELLOW}Enter commit day [dd] :${SET}"
read day

declare -i month
echo -e "${YELLOW}Enter commit month [mm] :${SET}"
read month

declare -i year
echo -e "${YELLOW}Enter commit year [yyyy] :${SET}"
read year

# exec from here
# -----
echo -e "${GREEN}changed files ::${SET}"
git diff --stat HEAD~5 HEAD
echo -e "${GREEN}---"

echo -e "${GREEN}adding all git files to commit [git add .]${SET}"
git add .
echo -e "${GREEN}---${SET}"

echo -e "${GREEN}commiting ...${SET}"
git commit -m "$commit_message"
echo -e "${GREEN}---${SET}"

echo -e "${GREEN}adjusting time ...${SET}"
git commit  -m "$commit_message" --amend --date="$month/$day 22:00 $year +0530"
echo -e "${GREEN}DONE[OK]${SET}"

