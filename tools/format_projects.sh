#!/bin/sh

# Format JSON and sort by decreasing date, then name
underscore print -i projects.json |
  jq -c '[group_by(.date)[] | sort_by(.name)] | sort_by(.[0].date) | reverse | map(.[]) | sort_by(.name == "wspace 0.2")' |
  underscore print |
  tools/sponge projects.json
