#!/usr/bin/env bash

function git_init(){
	local repo_path=$1
	echo "Initialising repo in: $repo_path"
	mkdir $repo_path/.git
	mkdir $repo_path/.git/objects
	mkdir $repo_path/.git/objects/info
	mkdir $repo_path/.git/objects/pack
	mkdir $repo_path/.git/branches
	echo -n "ref: refs/heads/master" > $repo_path/.git/HEAD
	mkdir $repo_path/.git/refs
	mkdir $repo_path/.git/refs/tags
	mkdir $repo_path/.git/refs/heads
	mkdir $repo_path/.git/config
	mkdir $repo_path/.git/hooks
	mkdir $repo_path/.git/description
	mkdir $repo_path/.git/info
}

git_init $1