#!/usr/bin/env bash
git config --global user.email "you@example.com"
git config --global user.name "Your Name"

rm -rf /opt/git/test
mkdir /opt/git/test
/opt/git/git-init.sh /opt/git/test
cd /opt/git/test
git status
echo 'version 1' > test.txt
git hash-object -w test.txt
echo 'version 2' > test.txt
git hash-object -w test.txt
find .git/objects -type f



#https://github.com/git/git/blob/master/Documentation/technical/index-format.txt
git update-index --add --cacheinfo 100644 83baae61804e65cc73a7201a7252750c76066a30 test.txt
echo "After update-index"
wc -c .git/index
git write-tree
wc -c .git/index
git cat-file -t d8329fc1cc938780ffdd9f94e0d364e0ea74f579
commit=$(echo 'first commit' | git commit-tree d8329f)
git cat-file -p $commit

echo "Before update HEAD"
echo "git status"
git status
echo "git log"
git log

echo $commit |tee ./.git/refs/heads/master
echo "After update HEAD"
echo "git status"
git status
echo "git log"
git log

#======================ruby=====================
#rm -rf /opt/git/test/.git
