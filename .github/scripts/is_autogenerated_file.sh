#!/usr/bin/env bash

FILEPATH="$1"

grep -q "Code generated by gtrace. DO NOT EDIT." "$FILEPATH" && echo YES && exit 0
grep -q "Code generated by MockGen. DO NOT EDIT." "$FILEPATH" && echo YES && exit 0
grep -q "Code generated by controller-gen. DO NOT EDIT." "$FILEPATH" && echo YES && exit 0

echo NO
exit 0