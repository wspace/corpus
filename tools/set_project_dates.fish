#!/usr/bin/env fish

set root (git rev-parse --show-toplevel)
set projects (jq -c .[] projects.json)
echo '[' > projects.json
for p in $projects
  set path (echo $p | jq -r .path)
  # Only check submodules
  if test -n "$path" -a (git -C $path rev-parse --show-toplevel) != "$root"
    set dates (git -C $path log --reverse --format=%ai\n%ci | head -n2)
    set author_date $dates[1]
    set commit_date $dates[2]
    if test "$author_date" = "$commit_date"
      echo "$p" | jq -c --arg d $author_date '.date = $d' | tr \n , >> projects.json
      continue
    end
    # Manually review unmatching author and committer dates
    echo "Author date $author_date and committer date $commit_date differ: $path" >&2
  end
  echo "$p," >> projects.json
end
echo ']' >> projects.json

tools/format.sh
