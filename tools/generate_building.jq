# Generate building.md

def ok: .bin!=null and .build_errors==null and .usage!=null and .run_errors==null;
def status: if . then "" else "⚠️ " end;
def escape: gsub("_"; "\\_")? | gsub("\\*"; "\\*")?;
def fmt:
  ((.bin | escape) // "*unspecified*") +
  if .build!=null or .build_errors!=null or .run_errors!=null then ":" else "" end +
  if .build!=null then " `\(.build)`" else "" end +
  if .build_errors!=null then " \(.build_errors | escape)" else "" end +
  if .run_errors!=null then " \(.run_errors | escape)" else "" end;

([(.[] | select(.build_errors!=null or .run_errors!=null or .source_unavailable).id),
    ($dockerfiles | split(" ")[] | sub("/Dockerfile$"; ""))] |
  map({key:.}) | from_entries) as $dockerfiles |

"# Building projects

<!-- Generated by tools/generate_building.jq; DO NOT EDIT. -->

Projects that can be built with Docker:
",
(
  map(select(.id | in($dockerfiles))) | sort_by(.id)[] |
  "- \(.id | escape)" +
  ([.build_errors, .run_errors,
      (if .source_unavailable then "Source unavailable" else null end) |
      select(.!=null)] |
    join("; ") |
    if . != "" then ": " + . else "" end)
),
"
Building status of individual executables:
",
(
  map(select(.id | in($dockerfiles) | not)) | sort_by(.id)[] |
  (.id | escape) as $id |
  if (.commands|length) == 0 and
      .languages == ["Whitespace"] and .tags == ["programs"] then empty
  else
    .commands |
    if length == 0 then "- ❌ \($id)"
    elif length == 1 then
      .[0] | "- \(ok | status)\($id)/" + fmt
    else
      "- \(all(ok) | status)\($id):",
      (sort_by(.bin)[] | "  - \(ok | status)" + fmt)
    end
  end
)
