#!/bin/bash
# Convert a Dockerfile to an Earthfile \
if (( "$#" != 1 )); then echo "Usage: dockerfile_to_earthfile.jq [Dockerfile]" >&2; fi; exec jq -Rrsf "$0" -- "${@:1}"

(input_filename // "File") as $filename |
[scan("(?:[^\\\\\n]|\\\\\n)*\n?") | sub("\n$"; "")] |
. as $lines |
# Desugar heredocs
(reduce range(0; length) as $i ([];
  def prev: .[-1];
  $lines[$i] as $line |
  if prev | type == "object" then
    if $line == prev.heredoc then
      # End of heredoc
      prev |= (.data | "RUN " + sub("(?:\\s*\\\\\n)+$"; ""; ""))
    else
      # Middle of heredoc
      prev.data +=
        if prev.data == "" then "" else "    " end
        + $line
        + if $line == "" then "" else " " end
        + "\\\n"
    end
  elif ($line | test("^\\s*#")) and (prev | . != null and test("^\\s*#")) then
    # Join adjacent comments
    prev += "\n" + $line
  else
    # Start of heredoc
    ($line | first(capture("^RUN <<(?<term>.+)$").term, null)) as $heredoc |
    . + [if $heredoc then {$heredoc, data: ""} else $line end]
  end
)) as $lines |
[
  # Group lines into FROM stages
  range(0; $lines | length) as $i |
  if $lines[$i] | test("^FROM ") then
    $i | if $lines[.-1] | test("^\\s*#") then .-1 end
  else empty end
] as $ranges |
$ranges |
if $ranges[0] != 0 then
  "Error: \($filename) does not start with FROM\n" | halt_error(1)
end |
[
  range(0; $ranges | length) as $i |
  $lines[$ranges[$i] : $ranges[$i+1]] as $stage | $stage |
  # Strip empty lines after FROM and at the end
  first(range(0; length) | select($stage[.] | test("^FROM "))) as $from |
  if .[$from+1] == "" then .[:$from+1] + .[$from+2:] end |
  if .[-1] == "" then .[:-1] end
] |
if length > 2 then
  "Error: \($filename) does not have 1 or 2 stages\n" | halt_error(1)
end |
map(
  ["stage:", (.[] | split("\n")[] | if . != "" then "    " + . end)] |
  join("\n")
) |
join("\n\n")
