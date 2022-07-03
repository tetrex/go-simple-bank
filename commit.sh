declare -p commit_message
echo "Enter commit message :"
read commit_message

declare -i day
echo "Enter commit day [dd] :"
read day

declare -i month
echo "Enter commit month [mm] :"
read month

declare -i year
echo "Enter commit year [yyyy] :"
read year

# exec from here
# -----
echo "changed files ::"
git diff --stat HEAD~5 HEAD
echo "---"

echo "adding all git files to commit [git add -A]"
git add -A
echo "---"

echo "commiting ..."
git commit -m "$commit_message"
echo "---"

echo "adjusting time ..."
git commit  -m "$commit_message" --amend --date="$month/$day 22:00 $year +0530"
echo "DONE[OK]"

