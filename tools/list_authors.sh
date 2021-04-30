#!/bin/sh

git -C "$1" log --format='%an <%ae> | %cn <%ce>' | sort | uniq -c | sort -rn
