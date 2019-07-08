#!/usr/bin/env bash
sh build.sh
git add bin/octopus-collector
git commit -m "$@"
git push origin master
