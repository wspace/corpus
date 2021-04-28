#!/usr/bin/env fish

set root (git rev-parse --show-toplevel)
rm -f tmp
for p in (jq -c .[] projects.json)
  set path (echo $p | jq -r .path)
  # Only check submodules
  if test -n "$path" -a (git -C $path rev-parse --show-toplevel) != "$root"
    set dates (git -C $path log --reverse --format=%ai\n%ci | head -n2)
    set author_date $dates[1]
    set commit_date $dates[2]
    if test "$author_date" = "$commit_date"
      echo "$p" | jq -c --arg d $author_date '.date = $d' >> tmp
      continue
    end
    # Manually review unmatching author and committer dates
    echo "Author date $author_date and committer date $commit_date differ: $path" >&2
  end
  echo $p >> tmp
end

jq -sc 'sort_by(.date) | reverse' tmp > projects.json
rm tmp
tools/format.sh
