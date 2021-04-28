#!/usr/bin/env fish

set projects (jq -c .[] projects.json)
echo '[' > projects.json
for p in $projects
  set path (echo $p | jq -r .path)
  if test -n "$path"
    set first (git -C $path rev-list --max-parents=0 HEAD)
    set dates (git -C $path show -s --format=%ai\n%ci $first)
    set author_date $dates[1]
    set commit_date $dates[2]
    if test "$author_date" = "$commit_date"
      echo "$p" | jq -c --arg d $author_date '.date = $d' >> projects.json
      echo , >> projects.json
      continue
    end
    echo "Author date $author_date and committer date $commit_date differ: $path" >&2
  end
  echo "$p," >> projects.json
end
echo ']' >> projects.json

tools/format.sh
