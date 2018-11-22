#!/usr/bin/env bash

git clone --depth 1 https://github.com/sstephenson/bats.git
echo Installed $(./bats/bin/bats --version)
