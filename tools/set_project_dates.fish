#!/usr/bin/env fish

set root (git rev-parse --show-toplevel)
echo '[' > tmp
for p in (jq -c .[] projects.json)
  set path (echo $p | jq -r .path)
  # Only check submodules
  if test -n "$path" -a -d "$path"
    set project_root (git -C $path rev-parse --show-toplevel)
    if test "$project_root" != "$root"
      set dates (git -C $path log --reverse --format=%ai\n%ci | head -n2)
      set author_date $dates[1]
      set commit_date $dates[2]
      if test "$author_date" = "$commit_date"
        echo "$p" | jq -c --arg d $author_date '.date = $d' | tr \n , >> tmp
        continue
      end
      # Manually review unmatching author and committer dates
      echo "Author date $author_date and committer date $commit_date differ: $path" >&2
    end
  end
  echo "$p," >> tmp
end
echo ']' >> tmp

mv tmp projects.json
tools/format_projects.bash
